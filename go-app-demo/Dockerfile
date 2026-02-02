# --- Stage 1: The Builder ---
# Use the official Golang image to compile the binary
FROM golang:1.25-alpine AS builder

# Set the working directory
WORKDIR /app

# Install build dependencies (git is often needed for go modules)
RUN apk add --no-cache git

# Copy go.mod and go.sum first to leverage Docker cache for dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
# CGO_ENABLED=0: Disables CGO to ensure a static binary that works in 'scratch' or 'distroless'
# GOOS=linux: Ensures the binary is built for the Linux kernel
# -ldflags="-w -s": Strips debug information to reduce binary size
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main .

# --- Stage 2: The Final SOTA Image ---
# 'distroless/static-debian12' contains only the CA certificates and timezone data
# It has NO shell, NO package manager, and NO root user—maximum security.
FROM gcr.io/distroless/static-debian12:latest

# Set the working directory in the final container
WORKDIR /

# Copy the binary from the builder stage
COPY --from=builder /app/main /main

# Fiber default port
EXPOSE 8080

# Use the pre-defined non-root 'nonroot' user (UID 65532) for security
USER 65532:65532

# Run the binary
ENTRYPOINT ["/main"]