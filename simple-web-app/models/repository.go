package models

import "database/sql"

type Repository interface {
	AllMusicAlbums() ([]*MusicAlbum, error)
}

// simple wrapper for sql.DB type, used to return MySQL repository
type mysqlRepository struct {
	DB *sql.DB
}

// factory method to return a new mysqlRepository
func newMysqlRepository(conn *sql.DB) Repository {
	return &mysqlRepository{
		DB: conn,
	}
}
