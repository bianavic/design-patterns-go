package models

import (
	"context"
	"log"
	"net/http"
	"time"
)

func (m *mysqlRepository) AllEntertainment() ([]*Entertainment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select  Media, MediaType, Type, Description from media_stream.Entertainment order by Type asc`

	var entertainments []*Entertainment

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var entertainment Entertainment
		err := rows.Scan(
			&entertainment.Media,
			&entertainment.MediaType,
			&entertainment.Type,
			&entertainment.Description)
		if err != nil {
			return nil, err
		}
		entertainments = append(entertainments, &entertainment)

		if err := rows.Err(); err != nil {
			log.Println("Internal Server Error", http.StatusInternalServerError)
		}
	}
	return entertainments, nil
}

func (m *mysqlRepository) GetEntertainmentByMediaType(mediaType string) (*Entertainment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select Media, MediaType, Type, Description from media_stream.Entertainment where Type = ?`

	var entertainment Entertainment

	err := m.DB.QueryRowContext(ctx, query, mediaType).Scan(
		&entertainment.Media,
		&entertainment.MediaType,
		&entertainment.Type,
		&entertainment.Description)
	if err != nil {
		return nil, err
	}
	return &entertainment, nil
}
