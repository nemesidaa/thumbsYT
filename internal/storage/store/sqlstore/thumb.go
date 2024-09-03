package store

import (
	"context"

	"github.com/nemesidaa/thumbsYT/internal/storage/model"
)

type ThumbRepo struct {
	Store *Storage
}

func (s *ThumbRepo) Save(ctx context.Context, id, path, resolution string) (*model.Thumb, context.Context, error) {
	_, err := s.Store.DB.Exec("INSERT INTO thumb VALUES (?, ?, ?)", id, path, resolution)
	if err != nil {
		return nil, context.WithValue(ctx, model.FailedThumbTX{}, id), err
	}
	return &model.Thumb{
		ID:         id,
		Path:       path,
		Resolution: resolution,
	}, context.WithValue(ctx, model.SuccessThumbTX{}, id), nil
}

func (s *ThumbRepo) GetByID(ctx context.Context, id string) (*model.Thumb, context.Context, error) {
	o := &model.Thumb{}
	err := s.Store.DB.QueryRow("SELECT id, path, resolution FROM thumb WHERE id = ?", id).Scan(
		&o.ID, &o.Path, &o.Resolution,
	)
	if err != nil {
		return nil, context.WithValue(ctx, model.FailedThumbTX{}, id), err
	}
	return o, context.WithValue(ctx, model.SuccessThumbTX{}, id), nil
}
