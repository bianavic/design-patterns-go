package models

import (
	"context"
	"log"
	"net/http"
	"time"
)

func (m *mysqlRepository) AllMusicAlbums() ([]*MusicAlbum, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select ID, Name, MediaType, Genre, coalesce(Details, ''), Country, ReleaseDate 
       from media_stream.MusicAlbum order by name asc`

	var musicAlbums []*MusicAlbum

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("Internal Server Error", http.StatusInternalServerError)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var musicAlbum MusicAlbum
		err := rows.Scan(
			&musicAlbum.ID,
			&musicAlbum.Name,
			&musicAlbum.MediaType,
			&musicAlbum.Genre,
			&musicAlbum.Details,
			&musicAlbum.Country,
			&musicAlbum.ReleaseDate)
		if err != nil {
			log.Println("Internal Server Error", http.StatusInternalServerError)
			return nil, err
		}
		musicAlbums = append(musicAlbums, &musicAlbum)

		// Check for errors from iterating over rows
		if err := rows.Err(); err != nil {
			log.Println("Internal Server Error", http.StatusInternalServerError)
		}
	}

	return musicAlbums, nil
}

func (m *mysqlRepository) GetMusicAlbumByName(name string) (*MusicAlbum, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select ID, Name, MediaType, Genre, coalesce(Details, ''), Country, ReleaseDate 
	   from media_stream.MusicAlbum where name = ?`

	var musicAlbum MusicAlbum

	err := m.DB.QueryRowContext(ctx, query, name).Scan(
		&musicAlbum.ID,
		&musicAlbum.Name,
		&musicAlbum.MediaType,
		&musicAlbum.Genre,
		&musicAlbum.Details,
		&musicAlbum.Country,
		&musicAlbum.ReleaseDate)
	if err != nil {
		log.Println("error getting album music by name", err)
		return nil, err
	}

	return &musicAlbum, nil
}
