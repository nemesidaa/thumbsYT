package store

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type Storage struct {
	DB    *sql.DB
	thumb *ThumbRepo
}

func NewStorage(name string) *Storage {
	db, err := sql.Open("sqlite", fmt.Sprintf("./%s", name))
	if err != nil {
		log.Fatalf("Failed to open database: %s", err)
	}

	store := &Storage{
		DB: db,
	}

	err = store.Init()
	if err != nil {
		log.Fatalf("Failed to init database: %s", err)
	}

	return store
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
		log.Fatalf("Failed to create table: %s", err)
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
