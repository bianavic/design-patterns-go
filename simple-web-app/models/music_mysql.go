package models

import (
	"context"
	"database/sql"
	"log"
	"time"
)

func (m *mysqlRepository) GetMusicOfMonthByID(id int) (*MusicOfMonth, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, image, video FROM MusicOfMonth WHERE id = ?`
	row := m.DB.QueryRowContext(ctx, query, id)

	var music MusicOfMonth
	err := row.Scan(
		&music.ID,
		&music.Video,
		&music.Image,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No result found, return nil
		}
		log.Println("error getting music of the month by ID:", err)
		return nil, err
	}

	return &music, nil
}

func (m *mysqlRepository) GetMusicByTitle(title string) (*Music, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// TODO update query to get music
	query := `SELECT ID, MusicAlbumID, Year, PersonID, Title, Artist FROM media_stream.Music WHERE Title = ?`
	row := m.DB.QueryRowContext(ctx, query, title)

	var music Music
	err := row.Scan(
		&music.ID,
		&music.MusicAlbumID,
		&music.Year,
		&music.PersonID,
		&music.Title,
		&music.Artist,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No result found, return nil
		}
		log.Println("error getting music by title:", err)
		return nil, err
	}

	return &music, nil
}
