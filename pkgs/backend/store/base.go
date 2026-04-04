package store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type base struct {
	db *sqlx.DB
}

func (b *base) InsertDatabase(ctx context.Context) error {
	return nil
}
