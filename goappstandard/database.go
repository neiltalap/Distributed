package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite" 
	// CGO-free driver: perfect for "scratch" or "distroless" images
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(dbPath string) (*Repository, error) {
	// 1. Open connection
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	// 2. Production-ready SQLite settings
	// WAL mode allows concurrent reads while a write is happening,
	// which is vital for maintaining readiness during heavy operations.
	_, err = db.Exec("PRAGMA journal_mode=WAL;")
	if err != nil {
		return nil, err
	}

	// Simple schema for stateful persistence
	query := `
	CREATE TABLE IF NOT EXISTS app_state (
		id INTEGER PRIMARY KEY,
		key TEXT UNIQUE,
		value TEXT
	);
	INSERT OR IGNORE INTO app_state (key, value) VALUES ('initialized_at', datetime('now'));
	`
	if _, err := db.Exec(query); err != nil {
		return nil, fmt.Errorf("failed to init schema: %w", err)
	}

	return &Repository{DB: db}, nil
}
