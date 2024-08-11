package models

import "database/sql"

type Repository interface {
	AllMusicAlbums() ([]*MusicAlbum, error)
	AllTVShows() ([]*TVShow, error)
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

func (m *testRepository) AllTVShows() ([]*TVShow, error) {
	return []*TVShow{}, nil
}
