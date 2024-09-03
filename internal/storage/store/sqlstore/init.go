package store

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	DB    *sql.DB
	thumb *ThumbRepo
}

func NewStorage() *Storage {
	db, err := sql.Open("sqlite3", "./capsule/storage.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	return &Storage{
		DB: db,
	}
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
