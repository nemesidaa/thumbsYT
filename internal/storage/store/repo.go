package store

import "context"

type ThumbRepo interface {
	Save(ctx context.Context, id, path, resolution string) (*ThumbRepo, context.Context, error)
	GetByID(ctx context.Context, id string) (*ThumbRepo, context.Context, error)
}
