package main

import (
	"encoding/xml"
	"github.com/go-chi/chi/v5"
	"github.com/tsawler/toolbox"
	"net/http"
	"simple-web-app/models"
)

func (app *application) GetAllMusicAlbumsJSON(writer http.ResponseWriter, request *http.Request) {
	var t toolbox.Tools
	musicAlbums, err := app.App.Models.MusicAlbum.All()
	if err != nil {
		_ = t.ErrorJSON(writer, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(writer, http.StatusOK, musicAlbums)
}

func (app *application) GetAllMusicAlbumsXML(writer http.ResponseWriter, request *http.Request) {
	var t toolbox.Tools

	// Get all breeds from the database.
	allMediaTypes, err := app.App.Models.MusicAlbum.All()
	if err != nil {
		_ = t.ErrorJSON(writer, err, http.StatusBadRequest)
		return
	}

	// Since we are sending a slice, we need a wrapper, or we will not have a root element.
	type mediaTypes struct {
		XMLName    xml.Name             `xml:"musicAlbums"`
		MediaTypes []*models.MusicAlbum `xml:"musicAlbum"`
	}

	// Structure the data we want to convert to XML.
	mediasType := mediaTypes{
		MediaTypes: allMediaTypes,
	}

	_ = t.WriteJSON(writer, http.StatusOK, mediasType)
}

func (app *application) GetAllTelevisionJSON(writer http.ResponseWriter, request *http.Request) {
	name := chi.URLParam(request, "television")
	var t toolbox.Tools

	// Get all  from the database.
	tv, err := app.App.Models.Television.GetTelevisionByName(name)
	if err != nil {
		_ = t.ErrorJSON(writer, err, http.StatusBadRequest)
		return
	}
	// Write the JSON out.
	_ = t.WriteJSON(writer, http.StatusOK, tv)
}

func (app *application) GetAllTelevisionNamesXML(writer http.ResponseWriter, request *http.Request) {
	name := chi.URLParam(request, "television")
	var t toolbox.Tools

	// Get all breeds from the database.
	mediaType, err := app.App.Models.Television.GetTelevisionByName(name)
	if err != nil {
		_ = t.ErrorJSON(writer, err, http.StatusBadRequest)
		return
	}
	// Write the XML out.
	_ = t.WriteXML(writer, http.StatusOK, mediaType)
}
