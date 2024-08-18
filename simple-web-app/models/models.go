package models

import "database/sql"

// place store db connections
var repo Repository

type Models struct {
	MusicAlbum    MusicAlbum
	TVShow        TVShow
	Entertainment Entertainment
	Television    Television
	Music         Music
	MusicOfMonth  MusicOfMonth
}

// factory function
func New(conn *sql.DB) *Models {
	if conn != nil {
		repo = newMysqlRepository(conn)
	} else {
		repo = newTestMysqlRepository(nil)
	}

	return &Models{
		MusicAlbum:    MusicAlbum{},
		TVShow:        TVShow{},
		Entertainment: Entertainment{},
		Television:    Television{},
		Music:         Music{},
		MusicOfMonth:  MusicOfMonth{},
	}
}

type Entertainment struct {
	Media       string `json:"media" xml:"Media"`
	MediaType   string `json:"media_type" xml:"MediaType"`
	Type        string `json:"type" xml:"Type"`
	Description string `json:"description" xml:"Description"`
}

type MusicAlbum struct {
	ID          int    `json:"id" xml:"ID"`
	Name        string `json:"name" xml:"Name"`
	MediaType   string `json:"media_type" xml:"MediaType"`
	Genre       string `json:"genre" xml:"Genre"`
	Details     string `json:"details" xml:"Details"`
	Country     string `json:"country" xml:"Country"`
	ReleaseDate string `json:"release_date" xml:"ReleaseDate"`
}

type MusicOfMonth struct {
	ID    int    `json:"id"`
	Music *Music `json:"music"`
	Video string `json:"video"`
	Image string `json:"image"`
}

type Television struct {
	ID          int    `json:"id" xml:"ID"`
	Name        string `json:"name" xml:"Name"`
	MediaType   TVShow `json:"media_type" xml:"MediaType"`
	Genre       string `json:"genre" xml:"Genre"`
	Details     string `json:"details" xml:"Details"`
	Country     string `json:"country" xml:"Country"`
	ReleaseDate string `json:"release_date" xml:"ReleaseDate"`
}

type Music struct {
	ID           int        `json:"id" xml:"ID"`
	MusicAlbumID int        `json:"music_album_id" xml:"MusicAlbumID"`
	Year         int        `json:"year" xml:"Year"`
	PersonID     int        `json:"person_id" xml:"PersonID"`
	Title        string     `json:"song_title" xml:"Title"`
	Artist       string     `json:"artist" xml:"Artist"`
	MediaType    MusicAlbum `json:"album" xml:"MediaType"`
	Person       Person     `json:"person" xml:"Person"`
}

type TVShow struct {
	ID           int    `json:"id" xml:"ID"`
	TelevisionID int    `json:"television_id" xml:"TelevisionID"`
	Year         int    `json:"year" xml:"Year"`
	PersonID     int    `json:"person_id" xml:"PersonID"`
	Title        string `json:"title" xml:"Title"`
	Director     string `json:"director" xml:"Director"`
	MediaType    string `json:"media_type" xml:"MediaType"`
	Person       Person `json:"person" xml:"Person"`
}

type Person struct {
	ID          int           `json:"id" xml:"ID"`
	Name        string        `json:"name" xml:"Name"`
	Address     string        `json:"address" xml:"Address"`
	Country     string        `json:"country" xml:"Country"`
	ZipCode     string        `json:"zip_code" xml:"ZipCode"`
	Email       string        `json:"email" xml:"Email"`
	Active      int           `json:"active" xml:"Active"`
	MusicAlbums []*MusicAlbum `json:"music_albums" xml:"MusicAlbums"`
	Televisions []*Television `json:"televisions" xml:"Televisions"`
}

func (m *MusicAlbum) All() ([]*MusicAlbum, error) {
	return repo.AllMusicAlbums()
}

func (tv *TVShow) All() ([]*TVShow, error) {
	return repo.AllTVShows()
}

func (m *MusicAlbum) GetMusicAlbumByName(name string) (*MusicAlbum, error) {
	return repo.GetMusicAlbumByName(name)
}

func (tv *Entertainment) AllEntertainment() ([]*Entertainment, error) {
	return repo.AllEntertainment()
}

func (tv *TVShow) GetTVShowByName(name string) (*TVShow, error) {
	return repo.GetTVShowByName(name)
}

func (e *Entertainment) GetEntertainmentByMediaType(mediaType string) (*Entertainment, error) {
	return repo.GetEntertainmentByMediaType(mediaType)
}

func (t *Television) GetTelevisionByName(name string) (*Television, error) {
	return repo.GetTelevisionByName(name)
}

func (t *Music) GetMusicOfMonthByID(id int) (*MusicOfMonth, error) {
	return repo.GetMusicOfMonthByID(id)
}

func (m *Music) GetMusicByTitle(title string) (*Music, error) {
	return repo.GetMusicByTitle(title)
}
