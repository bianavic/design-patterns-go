package models

import (
	"context"
	"log"
	"net/http"
	"time"
)

func (m *mysqlRepository) AllTVShows() ([]*TVShow, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select tv.ID, tv.TelevisionID, tv.Year, tv.PersonID, tv.Title, tv.Director,
                     t.Name as MediaType, p.Name as Person
              from media_stream.TVShow tv
              join media_stream.Television t on tv.TelevisionID = t.ID
              join media_stream.Person p on tv.PersonID = p.ID
              join media_stream.Television t on tv.TelevisionID = t.ID
              join media_stream.Person p on tv.PersonID = p.ID
              order by tv.Title asc`

	var tvShows []*TVShow

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("Internal Server Error", http.StatusInternalServerError)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tvShow TVShow
		err := rows.Scan(
			&tvShow.ID,
			&tvShow.TelevisionID,
			&tvShow.Year,
			&tvShow.PersonID,
			&tvShow.Title,
			&tvShow.Director,
			&tvShow.MediaType,
			&tvShow.Person)
		if err != nil {
			log.Println("Internal Server Error", http.StatusInternalServerError)
			return nil, err
		}
		tvShows = append(tvShows, &tvShow)

		// Check for errors from iterating over rows
		if err := rows.Err(); err != nil {
			log.Println("Internal Server Error", http.StatusInternalServerError)
		}
	}

	return tvShows, nil
}

func (m *mysqlRepository) GetTVShowByName(name string) (*TVShow, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select ID, TelevisionID, Year, PersonID, Title, Director from media_stream.TVShow where Title = ?`

	var tvShow TVShow

	err := m.DB.QueryRowContext(ctx, query, name).Scan(
		&tvShow.ID,
		&tvShow.TelevisionID,
		&tvShow.Year,
		&tvShow.PersonID,
		&tvShow.Title,
		&tvShow.Director,
		&tvShow.MediaType,
		&tvShow.Person)
	if err != nil {
		log.Println("Internal Server Error", http.StatusInternalServerError)
		return nil, err
	}

	return &tvShow, nil
}
