package store

import "context"

type Store interface {
	InsertDatabase(ctx context.Context) error
}
