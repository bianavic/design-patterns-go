package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/tsawler/toolbox"
	"net/http"
	"net/url"
	"simple-web-app/entertainment"
	"simple-web-app/models"
)

func (app *application) ShowHome(writer http.ResponseWriter, request *http.Request) {
	// because it is a pointer to templateData, can pass nil
	app.render(writer, "home.page.gohtml", nil)
}

func (app *application) ShowPage(writer http.ResponseWriter, request *http.Request) {
	page := chi.URLParam(request, "page")
	app.render(writer, fmt.Sprintf("%s.page.gohtml", page), nil)
}

func (app *application) MusicOfMonth(writer http.ResponseWriter, request *http.Request) {
	// get music
	musicTitle, err := app.App.Models.Music.GetMusicByTitle("Trustfall")
	if err != nil {
		http.Error(writer, "Error retrieving music", http.StatusInternalServerError)
		return
	}
	if musicTitle == nil {
		http.Error(writer, "Music not found", http.StatusNotFound)
		return
	}

	// get album of the month from db
	mom, err := app.App.Models.Music.GetMusicOfMonthByID(2)
	if err != nil {
		http.Error(writer, "Error retrieving music of the month", http.StatusInternalServerError)
		return
	}
	if mom == nil {
		http.Error(writer, "Music of the month not found", http.StatusNotFound)
		return
	}

	// create music and decorate it
	// decorator variable
	music := models.MusicOfMonth{
		ID: 2,
		// embedded
		Music: &models.Music{
			ID:           2,
			MusicAlbumID: musicTitle.ID,
			Year:         2023,
			Title:        "Trustfall",
			Artist:       "Pink",
		},
		// decorates with video and image
		Video: mom.Video,
		Image: mom.Image,
	}

	// Serve the web page
	data := make(map[string]any)
	data["music"] = music

	app.render(writer, "music-of-month.page.gohtml", &templateData{
		Data: data,
	})
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

	app.render(writer, "music-video-clips.page.gohtml", data)
}

func (app *application) ShowMusicAlbums(writer http.ResponseWriter, request *http.Request) {

	var t toolbox.Tools
	musicAlbums, err := app.App.Models.MusicAlbum.All()
	if err != nil {
		_ = t.ErrorJSON(writer, err, http.StatusBadRequest)
		return
	}

	data := &templateData{
		Data: map[string]any{
			"MusicAlbums": musicAlbums,
		},
	}

	app.render(writer, "music-albums.page.gohtml", data)
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

func (app *application) CreateMusicFromFactory(writer http.ResponseWriter, request *http.Request) {
	var t toolbox.Tools
	_ = t.WriteJSON(writer, http.StatusOK, entertainment.NewEntertainment("Music created from factory"))
}

func (app *application) CreateTelevisionFromFactory(writer http.ResponseWriter, request *http.Request) {
	var t toolbox.Tools
	_ = t.WriteJSON(writer, http.StatusOK, entertainment.NewEntertainment("Television created from factory"))
}

func (app *application) TestPatterns(writer http.ResponseWriter, request *http.Request) {
	app.render(writer, "test.page.gohtml", nil)
}

func (app *application) CreateMusicFromAbstractFactory(writer http.ResponseWriter, request *http.Request) {
	var t toolbox.Tools
	music, err := entertainment.NewEntertainmentFromAbstractFactory("music")
	if err != nil {
		_ = t.WriteJSON(writer, http.StatusBadRequest, err.Error())
		return
	}
	// send to user as a response
	_ = t.WriteJSON(writer, http.StatusOK, music)
}

func (app *application) CreateTelevisionFromAbstractFactory(writer http.ResponseWriter, request *http.Request) {
	var t toolbox.Tools
	television, err := entertainment.NewEntertainmentFromAbstractFactory("television")
	if err != nil {
		_ = t.WriteJSON(writer, http.StatusBadRequest, err.Error())
		return
	}
	// send to user as a response
	_ = t.WriteJSON(writer, http.StatusOK, television)
}

func (app *application) GetAllMusicAlbumsJSON(writer http.ResponseWriter, request *http.Request) {
	var t toolbox.Tools
	musicAlbums, err := app.App.Models.MusicAlbum.All()
	if err != nil {
		_ = t.ErrorJSON(writer, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(writer, http.StatusOK, musicAlbums)
}

func (app *application) CreateMusicAlbumWithBuilder(writer http.ResponseWriter, request *http.Request) {
	var t toolbox.Tools

	// create a new music album using builder pattern
	m, err := entertainment.NewEntertainmentBuilder().
		SetType("Music Album").
		SetMediaType("CD").
		SetGenre("Pop").
		SetDescription("Description 1").
		SetCountry("USA").
		SetRating(5).
		SetReleaseYear(2022).
		SetAvailable(true).
		Build()
	if err != nil {
		_ = t.ErrorJSON(writer, err, http.StatusBadRequest)
		return
	}

	// send to user as a response
	_ = t.WriteJSON(writer, http.StatusOK, m)
}

func (app *application) CreateTelevisionWithBuilder(writer http.ResponseWriter, request *http.Request) {
	var t toolbox.Tools

	// create a new music album using builder pattern
	m, err := entertainment.NewEntertainmentBuilder().
		SetType("TV Show").
		SetMediaType("Streaming").
		SetGenre("Crime").
		SetDescription("Description 1").
		SetCountry("Korea").
		SetRating(9).
		SetReleaseYear(2023).
		SetAvailable(true).
		Build()
	if err != nil {
		_ = t.ErrorJSON(writer, err, http.StatusBadRequest)
		return
	}

	// send to user as a response
	_ = t.WriteJSON(writer, http.StatusOK, m)
}

func (app *application) GetAllTVShows(writer http.ResponseWriter, request *http.Request) {
	var t toolbox.Tools

	tvShows, err := app.App.TVShowService.GetAllTVShows()
	if err != nil {
		_ = t.ErrorJSON(writer, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(writer, http.StatusOK, tvShows)
}

func (app *application) GetAllMusicAlbums(writer http.ResponseWriter, request *http.Request) {
	var t toolbox.Tools

	tvShows, err := app.App.Models.MusicAlbum.All()
	if err != nil {
		_ = t.ErrorJSON(writer, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(writer, http.StatusOK, tvShows)
}

/*
STUB HANDLER
new factory function for Entertainment with embbed music albums and tv shows
*/
func (app *application) EntertainmentFromAbstractFactory(writer http.ResponseWriter, request *http.Request) {
	var t toolbox.Tools

	en := chi.URLParam(request, "entertainment")
	mt := chi.URLParam(request, "media_type")

	mediaType, _ := url.QueryUnescape(mt)

	fmt.Println("entertainment: ", en, "mediaType: ", mediaType)

	// create entertainment from abstract factory
	result, err := entertainment.NewEntertainmentMediaTypeFromAbstractFactory(en, mediaType, "", nil)
	if err != nil {
		_ = t.ErrorJSON(writer, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(writer, http.StatusOK, result)
}
