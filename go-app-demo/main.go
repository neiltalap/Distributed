package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"
)

func main() {
	dbPath := os.Getenv("DATABASE_URL")
	if dbPath == "" {
		dbPath = "./app.db"
	}

	repo, err := NewRepository(dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 1. Context for the background Health Manager
	healthCtx, stopHealth := context.WithCancel(context.Background())
	defer stopHealth()

	h := NewHealthManager(repo)
	go h.Run(healthCtx)

	app := fiber.New(fiber.Config{
		AppName: "Go App Standard",
	})

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("State of the art Go on k8s!")
	})

	// --- Health Probes ---

	// Liveness: Process is alive.
	// Avoid complex checks here so K8s doesn't restart the pod during transient DB blips.
	app.Get("/health/liveness", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "up"})
	})

	// Readiness: Can we handle traffic?
	// Reports the state managed by our background ticker.
	app.Get("/readiness", func(c fiber.Ctx) error {
		if h.isReady.Load() {
			return c.SendStatus(fiber.StatusOK)
		}
		return c.SendStatus(fiber.StatusServiceUnavailable)
	})

	// Startup: Has the app finished initializing?
	// K8s will wait for this to succeed before starting liveness/readiness probes.
	app.Get("/health/startupz", func(c fiber.Ctx) error {
		if h.isStarted.Load() {
			return c.SendStatus(fiber.StatusOK)
		}
		return c.Status(fiber.StatusServiceUnavailable).SendString("Starting up...")
	})

	// --- Graceful Shutdown Logic ---

	// Listen for SIGTERM (K8s) or SIGINT (Ctrl+C)
	sigCtx, stopSignals := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stopSignals()

	// Start the HealthManager in the background
	// We use the same sigCtx so it stops when the app shuts down
	go h.Run(sigCtx)

	// Wait for the termination signal
	<-sigCtx.Done()
	log.Println("Shutting down... (K8s SIGTERM received)")

	// Wait 5-10 seconds before actually stopping the server
	// this allows k8s to propagate the "this pod is terminating" status
	// to the K8s Service Mesh/Ingress so we don't get 502/504 errors 
	// on the last few requests
	time.Sleep(5 * time.Second)

	// 1. Define a grace period context
	shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancelShutdown()

	// 2. Shut down the web server first (stops accepting new requests)
	if err := app.ShutdownWithContext(shutdownCtx); err != nil {
		log.Printf("Server shutdown failed: %v", err)
	}

	// 3. Stop background health checks
	stopHealth()

	// 4. Finally, close the database connection
	if err := repo.DB.Close(); err != nil {
		log.Printf("DB close error: %v", err)
	}

	log.Println("Server gracefully stopped.")
}
