package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"net/http"
)

func (app *application) routes() http.Handler {

	mux := chi.NewRouter()
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	mux.Get("/api/music-albums/all/json", app.GetAllMusicAlbumsJSON)
	mux.Get("/api/music-albums/all/xml", app.GetAllMusicAlbumsXML)
	//mux.Get("/api/music-albums", app.ShowMusicAlbums)

	return mux
}
