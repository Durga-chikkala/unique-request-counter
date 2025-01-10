package store

import "context"

type Store interface {
	RequestCount(ctx context.Context) int64
	Flush(ctx context.Context) bool
	LockId(ctx context.Context, id int64) bool
	IncrementCount(ctx context.Context)
}
