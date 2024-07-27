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

	// display test page
	mux.Get("/test-patterns", app.TestPatterns)
	// factory routes
	mux.Get("/api/music-factory", app.CreateMusicFromFactory)
	mux.Get("/api/television-factory", app.CreateTelevisionFromFactory)

	mux.Get("/", app.ShowHome)
	mux.Get("/{page}", app.ShowPage)
	mux.Get("/music-songs", app.ShowMusicSongs)
	mux.Get("/tv-shows", app.ShowTVShows)

	return mux
}
