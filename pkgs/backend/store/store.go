package store

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/deepakdinesh1123/mezcala/pkgs/backend/config"
	"github.com/jmoiron/sqlx"
)

func NewDBStore(ctx context.Context, envConfig *config.EnvConfig) (Store, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to resolve home dir: %w", err)
	}

	dbDir := filepath.Join(home, ".mez")
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create db dir: %w", err)
	}

	db, err := sqlx.Open("sqlite", filepath.Join(dbDir, "mez.db"))
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	return &sqliteStore{base: base{
		db: db,
	}}, nil
}
