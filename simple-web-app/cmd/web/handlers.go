package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (app *application) ShowHome(writer http.ResponseWriter, request *http.Request) {
	// because it is a pointer to templateData, can pass nil
	app.render(writer, "home.page.gohtml", nil)
}

func (app *application) ShowPage(writer http.ResponseWriter, request *http.Request) {
	page := chi.URLParam(request, "page")
	app.render(writer, fmt.Sprintf("%s.page.gohtml", page), nil)
}

func (app *application) ShowMusicSongs(writer http.ResponseWriter, request *http.Request) {

	songs := []struct {
		Title string
		URL   string
	}{
		{"Taylor Swift - Exile", "https://www.youtube.com/embed/o5SQIECedTY?si=HkY7GBXmsDw9g71q"},
		{"Dr Dre - Still D.R.E ft. Snoop Dog", "https://www.youtube.com/embed/_CL6n0FJZpk"},
		{"50 Cent - In Da Club", "https://www.youtube.com/embed/5qm8PH4xAss"},
		{"Pink - TrustFall", "https://www.youtube.com/embed/D2KE2a5qo0g"},
	}

	// Render the page with song data
	data := &templateData{
		Data: map[string]any{
			"Songs": songs,
		},
	}

	app.render(writer, "music-songs.page.gohtml", data)
}

func (app *application) ShowTVShows(writer http.ResponseWriter, request *http.Request) {

	tvShows := []struct {
		Title string
		URL   string
	}{
		{"Voice - Kdrama", "https://www.youtube.com/embed/dHLeCdpnGpY?si=rebMKCXJP14WDsQK"},
		{"The Bequeathed - Kdrama", "https://www.youtube.com/embed/aNugXXbWtAw?si=xfW_J3Y834ElvNEM"},
		{"Beyond Evil - Kdrama", "https://www.youtube.com/embed/zQCOdcBAWtU?si=3LPEG1p7WIsO1VbG"},
		{"The Good Detective - Kdrama", "https://www.youtube.com/embed/nY4JDIyzA3Q?si=kG6GKFd05vEElFFp"},
	}

	// Render the page with song data
	data := &templateData{
		Data: map[string]any{
			"TVShow": tvShows,
		},
	}

	app.render(writer, "tv-shows.page.gohtml", data)
}
