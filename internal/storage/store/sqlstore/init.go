package store

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Storage struct {
	DB    *sql.DB
	thumb *ThumbRepo
}

func InitStorage(name string) error {
	db, err := sql.Open("sqlite", fmt.Sprintf("./%s", name))
	if err != nil {
		return err
	}

	store := &Storage{
		DB: db,
	}

	err = store.Init()
	if err != nil {
		return err
	}

	return nil
}
func (s *Storage) Init() error {
	// migrations for laziest
	createThumbsSQL := `
    CREATE TABLE IF NOT EXISTS thumb (
        id varchar(100) PRIMARY KEY,
        data TEXT NOT NULL,
        resolution string NOT NULL
    );
    `
	_, err := s.DB.Exec(createThumbsSQL)
	if err != nil {
		return err
	}
	if err = s.DB.Ping(); err != nil {
		return err
	}
	return nil
}

func (s *Storage) Thumb() *ThumbRepo {
	if s.thumb != nil {
		return s.thumb
	}
	s.thumb = &ThumbRepo{Store: s}
	return s.thumb
}

func NewConn(name string) (*Storage, func() error, error) {
	db, err := sql.Open("sqlite", name)
	if err != nil {
		return nil, nil, err
	}

	return &Storage{
		DB: db,
	}, db.Close, err

}
