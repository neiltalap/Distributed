package main

import (
	"context"
	"sync/atomic"
	"time"
)

type HealthManager struct {
	repo      *Repository
	isReady   atomic.Bool
	isStarted atomic.Bool
	trigger   chan struct{}
}

func NewHealthManager(repo *Repository) *HealthManager {
	return &HealthManager{
		repo:    repo,
		trigger: make(chan struct{}, 1),
	}
}

func (h *HealthManager) Run(ctx context.Context) {
	// Check every 15s, but also check immediately on startup
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	// Initial check
	h.check()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			h.check()
		case <-h.trigger:
			h.check()
		}
	}
}

func (h *HealthManager) check() {
	// Dependency Check: Ping the SQLite DB
	if err := h.repo.DB.Ping(); err != nil {
		h.isReady.Store(false)
		return
	}

	// If we successfully pinged, we are ready to take traffic
	h.isReady.Store(true)

	// If we successfully pinged for the first time, we are officially "started"
	// Startup probes stop once this returns a success code.
	h.isStarted.Store(true)
}

func (h *HealthManager) TriggerCheck() {
	select {
	case h.trigger <- struct{}{}:
	default:
		// Check already pending
	}
}
