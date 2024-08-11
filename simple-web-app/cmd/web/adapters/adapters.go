package adapters

import (
	"encoding/json"
	"encoding/xml"
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

type TelevisionInterface interface {
	GetAllTelevisions() ([]*models.Television, error)
}

type RemoteService struct {
	Remote   TelevisionInterface
	ERemote  EntertainmentInterface
	MRemote  MusicAlbumsInterface
	TVRemote TVShowInterface
}

// setup json adapter
type JSONBackend struct{}

// setup xml adapter
type XMLBackend struct{}

func (xb *XMLBackend) GetAllTelevisions() ([]*models.Television, error) {
	resp, err := http.Get("http://localhost:8081/api/televisions/all/xml")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type televisions struct {
		XMLName     xml.Name             `xml:"televisions"`
		Televisions []*models.Television `xml:"television"`
	}

	var tvs televisions
	err = xml.Unmarshal(body, &tvs)
	if err != nil {
		return nil, err
	}

	return tvs.Televisions, nil
}

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

// implement interface to xml backend
func (xb *XMLBackend) GetAllMusicAlbums() ([]*models.MusicAlbum, error) {
	resp, err := http.Get("http://localhost:8081/api/music-albums/all/xml")
	if err != nil {
		return nil, err
	}

	// prevents resource leak
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// consume xml
	type musicAlbums struct {
		//XMLName xml.Name `xml:"musicAlbums"`
		XMLName struct{}             `xml:"music-albums"`
		Albums  []*models.MusicAlbum `xml:"music-albums"`
	}

	var albums musicAlbums
	err = xml.Unmarshal(body, &albums)
	if err != nil {
		return nil, err
	}

	return albums.Albums, nil
}

func (rs *RemoteService) GetAllMusicAlbumsXML() ([]*models.MusicAlbum, error) {
	return rs.MRemote.GetAllMusicAlbums()
}
