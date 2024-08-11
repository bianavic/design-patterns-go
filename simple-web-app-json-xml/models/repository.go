package models

import "database/sql"

type Repository interface {
	AllMusicAlbums() ([]*MusicAlbum, error)
	GetMusicAlbumByID(id int) (*MusicAlbum, error)
	RandomMusicAlbum() (*MusicAlbum, error)
	GetMusicAlbumByName(name string) (*MusicAlbum, error)
	//RandomMusicAlbumByGenre() (*MusicAlbum, error)
	AllTelevision() ([]*Television, error)
	GetTelevisionByID(id int) (*Television, error)
	GetTelevisionByName(id string) (*Television, error)
	RandomTelevision() (*Television, error)
	//RandomTelevisionByGenre() ([]*Television, error)
}

// wrapper for sql.DB type, to return MySQL repository
type mysqlRepository struct {
	DB *sql.DB
}

// factory method to return a new mysqlRepository
func newMysqlRepository(conn *sql.DB) Repository {
	return &mysqlRepository{
		DB: conn,
	}
}

// wrapper for the *sql.DB type, to return a test repository
type testRepository struct {
	DB *sql.DB
}

// factory method to return a new mysqlRepository
func newTestMysqlRepository(conn *sql.DB) Repository {
	return &testRepository{
		DB: nil,
	}
}
