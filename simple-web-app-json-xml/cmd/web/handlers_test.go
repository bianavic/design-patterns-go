package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplication_GetAllMusicAlbumsJSON(t *testing.T) {

	// create request
	req, _ := http.NewRequest("GET", "/api/music-albums", nil)
	// create a response recorder
	response := httptest.NewRecorder()
	// create handler
	handler := http.HandlerFunc(testApp.GetAllMusicAlbumsJSON)
	// serve the handler
	handler.ServeHTTP(response, req)
	// check response
	if response.Code != http.StatusOK {
		t.Errorf("expected status 200 but got %d", response.Code)
	}

}
