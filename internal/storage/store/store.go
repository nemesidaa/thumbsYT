package store

import "context"

type Store interface {
	Thumb() ThumbRepo
	Clear(ctx context.Context) error
}
