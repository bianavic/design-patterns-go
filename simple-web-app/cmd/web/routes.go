package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

func (app *application) routes() http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second))

	//fileServer := http.FileServer(http.Dir("./static"))
	//mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	// FileServer handler with MIME type detection
	mux.Handle("/static/*", http.StripPrefix("/static", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		ext := strings.ToLower(filepath.Ext(path))

		// Set the Content-Type based on the file extension
		switch ext {
		case ".webp":
			w.Header().Set("Content-Type", "image/webp")
		case ".jpeg", ".jpg":
			w.Header().Set("Content-Type", "image/jpeg")
		case ".png":
			w.Header().Set("Content-Type", "image/png")
		case ".gif":
			w.Header().Set("Content-Type", "image/gif")
		default:
			// Optionally handle other MIME types or use a default type
			w.Header().Set("Content-Type", "application/octet-stream")
		}

		// Serve the file
		http.ServeFile(w, r, "./static"+path)
	})))

	mux.Get("/music-of-month", app.MusicOfMonth)

	// display test page
	mux.Get("/test-patterns", app.TestPatterns)

	// factory routes
	mux.Get("/api/music-factory", app.CreateMusicFromFactory)
	mux.Get("/api/television-factory", app.CreateTelevisionFromFactory)
	//mux.Get("/api/music-abstract-factory", app.CreateMusicFromAbstractFactory)
	//mux.Get("/api/television-abstract-factory", app.CreateTelevisionFromAbstractFactory)

	// builder routes
	mux.Get("/api/music-album-builder", app.CreateMusicAlbumWithBuilder)
	mux.Get("/api/television-builder", app.CreateTelevisionWithBuilder)

	mux.Get("/", app.ShowHome)
	mux.Get("/{page}", app.ShowPage)

	mux.Get("/music-video-clips", app.ShowMusicSongs)
	mux.Get("/tv-shows", app.ShowTVShows)

	mux.Get("/api/music-albums", app.GetAllMusicAlbums)
	mux.Get("/api/tv-shows-json", app.GetAllTVShows)

	//mux.Get("/api/media-from-abstract-factory/{entertainment}/{media_type}", app.EntertainmentFromAbstractFactory)

	//mux.Get("/api/entertainment-from-abstract-factory/{entertainment}/{mediaType}", app.EntertainmentFromAbstractFactory)

	//mux.Get("/api/entertainment", app.GetAllEntertainmentJSON)

	//mux.Get("/api/music-albums", app.ShowMusicAlbums)

	return mux
}
