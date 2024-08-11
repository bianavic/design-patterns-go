package adapters

import (
	"encoding/json"
	"io"
	"net/http"
	"simple-web-app/models"
)

type MusicAlbumsInterface interface {
	GetAllMusicAlbums() ([]*models.MusicAlbum, error)
}

type EntertainmentInterface interface {
	GetAllTEntertainments() ([]*models.Entertainment, error)
}

type TVShowInterface interface {
	GetAllTVShows() ([]*models.TVShow, error)
}

type RemoteService struct {
	Remote   TelevisionInterface
	ERemote  EntertainmentInterface
	MRemote  MusicAlbumsInterface
	TVRemote TVShowInterface
}

// setup adapter
type JSONBackend struct{}

func (rs *RemoteService) GetAllMusicAlbums() ([]*models.MusicAlbum, error) {
	return rs.MRemote.GetAllMusicAlbums()
}

func (jb *JSONBackend) GetAllMusicAlbums() ([]*models.MusicAlbum, error) {
	// logic to connect to json backend
	resp, err := http.Get("http://localhost:8081/api/music-albums/all/json")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var albums []*models.MusicAlbum
	err = json.Unmarshal(body, &albums)
	if err != nil {
		return nil, err
	}

	return albums, nil
}

type TelevisionInterface interface {
	GetAllTelevisions() ([]*models.Television, error)
}

func (rs *RemoteService) GetAllTelevisions() ([]*models.Television, error) {
	return rs.Remote.GetAllTelevisions()
}

func (jb *JSONBackend) GetAllTelevisions() ([]*models.Television, error) {
	// logic to connect to json backend
	resp, err := http.Get("http://localhost:8081/api/televisions/all/json")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var televisions []*models.Television
	err = json.Unmarshal(body, &televisions)
	if err != nil {
		return nil, err
	}

	return televisions, nil
}

func (rs *RemoteService) GetAllEntertainments() ([]*models.Entertainment, error) {
	return rs.ERemote.GetAllTEntertainments()
}

func (jb *JSONBackend) GetAllEntertainments() ([]*models.Entertainment, error) {
	// logic to connect to json backend
	resp, err := http.Get("http://localhost:8081/api/televisions/all/json")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var entertainments []*models.Entertainment
	err = json.Unmarshal(body, &entertainments)
	if err != nil {
		return nil, err
	}

	return entertainments, nil
}

func (rs *RemoteService) GetAllTVShows() ([]*models.TVShow, error) {
	return rs.TVRemote.GetAllTVShows()
}

func (jb *JSONBackend) GetAllTVShows() ([]*models.TVShow, error) {
	// logic to connect to json backend
	resp, err := http.Get("http://localhost:8081/api/tv-shows/all/json")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tvShows []*models.TVShow
	err = json.Unmarshal(body, &tvShows)
	if err != nil {
		return nil, err
	}

	return tvShows, nil
}
