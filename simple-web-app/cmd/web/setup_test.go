package main

import (
	"os"
	"simple-web-app/cmd/web/adapters"
	"simple-web-app/configuration"
	"simple-web-app/models"
	"testing"
)

// set testing environment
var testApp application

func TestMain(m *testing.M) {

	testBackend := &TestBackend{}
	testAdapter := adapters.RemoteService{MRemote: testBackend}

	testApp = application{
		App:          configuration.New(nil),
		albumService: &testAdapter,
	}

	os.Exit(m.Run())
}

type TestBackend struct{}

func (tb *TestBackend) GetAllMusicAlbums() ([]*models.MusicAlbum, error) {
	albums := []*models.MusicAlbum{
		&models.MusicAlbum{
			ID:          1,
			Name:        "Abbey Road",
			MediaType:   "CD",
			Genre:       "Rock",
			Details:     "The Beatles",
			Country:     "UK",
			ReleaseDate: "1969-09-26",
		},
	}
	return albums, nil
}

//func (tb *TestBackend) GetAllTelevisions() ([]*models.Television, error) {
//	tvs := []*models.Television{
//		&models.Television{
//			ID:          1,
//			Name:        "The Crown",
//			MediaType:   "TV",
//			Genre:       "Drama",
//			Details:     "The Crown is a historical drama streaming television series about the reign of Queen Elizabeth II, created and principally written by Peter Morgan, and produced by Left Bank Pictures and Sony Pictures Television for Netflix.",
//			Country:     "UK",
//			ReleaseDate: "2016-11-04",
//		},
//	}
//	return tvs, nil
//}
