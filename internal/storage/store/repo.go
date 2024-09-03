package store

import "context"

type ThumbRepo interface {
	Save(ctx context.Context, id, data, resolution string) (*ThumbRepo, context.Context, error)
	GetByID(ctx context.Context, id string) (*ThumbRepo, context.Context, error)
}
