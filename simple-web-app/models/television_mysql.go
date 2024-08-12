package models

import (
	"context"
	"log"
	"net/http"
	"time"
)

func (m *mysqlRepository) AllTelevision() ([]*Television, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select ID, Name, MediaType, Genre, coalesce(Details, ''), Country, ReleaseDate 
	   from media_stream.Television order by name asc`

	var televisions []*Television

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("Internal Server Error", http.StatusInternalServerError)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var television Television
		err := rows.Scan(
			&television.ID,
			&television.Name,
			&television.MediaType,
			&television.Genre,
			&television.Details,
			&television.Country,
			&television.ReleaseDate)
		if err != nil {
			log.Println("Internal Server Error", http.StatusInternalServerError)
			return nil, err
		}
		televisions = append(televisions, &television)

		// Check for errors from iterating over rows
		if err := rows.Err(); err != nil {
			log.Println("Internal Server Error", http.StatusInternalServerError)
		}
	}

	return televisions, nil
}

func (m *mysqlRepository) GetTelevisionByName(name string) (*Television, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select ID, Name, MediaType, Genre, coalesce(Details, ''), Country, ReleaseDate 
	   from media_stream.Television where Name = ?`

	var television Television

	err := m.DB.QueryRowContext(ctx, query, name).Scan(
		&television.ID,
		&television.Name,
		&television.MediaType,
		&television.Genre,
		&television.Details,
		&television.Country,
		&television.ReleaseDate)
	if err != nil {
		return nil, err
	}
	return &television, nil
}
