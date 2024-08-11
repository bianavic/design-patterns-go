package models

import "database/sql"

type Repository interface {
	AllMusicAlbums() ([]*MusicAlbum, error)
	GetMusicAlbumByName(name string) (*MusicAlbum, error)
	AllTVShows() ([]*TVShow, error)
	GetTVShowByName(name string) (*TVShow, error)
	AllEntertainment() ([]*Entertainment, error)
	GetEntertainmentByMediaType(mediaType string) (*Entertainment, error)
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

type testRepository struct {
	DB *sql.DB
}

func newTestMysqlRepository(conn *sql.DB) Repository {
	return &testRepository{
		DB: nil,
	}
}
