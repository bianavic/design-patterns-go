package models

import (
	"context"
	"time"
)

func (m *MusicAlbum) AllMusicAlbums() ([]*MusicAlbum, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select ID, Name, MediaType, Genre, coalesce(Details, ''), Country, ReleaseDate 
       from media_stream.MusicAlbum order by name asc`

	var musicAlbums []*MusicAlbum

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
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
			return nil, err
		}
		musicAlbums = append(musicAlbums, &musicAlbum)
	}

	return musicAlbums, nil
}
