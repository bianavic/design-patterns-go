package models

import "database/sql"

// place store db connections
var repo Repository

// wrapper for all database models
type Models struct {
	MusicAlbum    MusicAlbum
	Television    Television
	Song          Song
	TVShow        TVShow
	Person        Person
	Entertainment Entertainment
}

// factory pattern
func New(conn *sql.DB) *Models {
	if conn != nil {
		repo = newMysqlRepository(conn)
	}
	//else {
	//	repo = newTestMysqlRepository(conn)
	//}

	return &Models{
		MusicAlbum: MusicAlbum{},
		Television: Television{},
	}
}

type Entertainment struct {
	Media       string `json:"media"`
	MediaType   string `json:"media_type"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type MusicAlbum struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	MediaType   string `json:"media_type"`
	Genre       string `json:"genre"`
	Details     string `json:"details"`
	Country     string `json:"country"`
	ReleaseDate string `json:"release_date"`
}

type Television struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	MediaType   string `json:"media_type"`
	Genre       string `json:"genre"`
	Details     string `json:"details"`
	Country     string `json:"country"`
	ReleaseDate string `json:"release_date"`
}

type Song struct {
	ID           int        `json:"id"`
	MusicAlbumID int        `json:"music_album_id"`
	Year         int        `json:"year"`
	PersonID     int        `json:"person_id"`
	SongTitle    string     `json:"song_title"`
	Artist       string     `json:"artist"`
	MediaType    MusicAlbum `json:"album"`
	Person       Person     `json:"person"`
}

type TVShow struct {
	ID           int        `json:"id"`
	TelevisionID int        `json:"television_id"`
	Year         int        `json:"year"`
	PersonID     int        `json:"person_id"`
	Title        string     `json:"title"`
	Director     string     `json:"director"`
	MediaType    Television `json:"television"`
	Person       Person     `json:"person"`
}

type Person struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Address     string        `json:"address"`
	Country     string        `json:"country"`
	ZipCode     string        `json:"zip_code"`
	Email       string        `json:"email"`
	Active      int           `json:"active"`
	MusicAlbums []*MusicAlbum `json:"music_albums"`
	Televisions []*Television `json:"televisions"`
}

func (m *MusicAlbum) All() ([]*MusicAlbum, error) {
	return repo.AllMusicAlbums()
}

func (m *MusicAlbum) Random() (*MusicAlbum, error) {
	return repo.RandomMusicAlbum()
}

func (m *MusicAlbum) Get(id int) (*MusicAlbum, error) {
	return repo.GetMusicAlbumByID(id)
}

func (t *Television) All() ([]*Television, error) {
	return repo.AllTelevision()
}

func (t *Television) Random() (*Television, error) {
	return repo.RandomTelevision()
}

func (t *Television) Get(id int) (*Television, error) {
	return repo.GetTelevisionByID(id)
}

func (m *Television) GetTelevisionByName(id string) (*Television, error) {
	return repo.GetTelevisionByName(id)
}

func (m *MusicAlbum) GetEntertainment() (*Entertainment, error) {
	return &Entertainment{}, nil
}

func (m *MusicAlbum) ReleaseMusicDate() string {
	return m.ReleaseDate
}

func (t *Television) GetEntertainment() (*Entertainment, error) {
	return &Entertainment{}, nil
}

func (t *Television) ReleaseTelevisionDate() string {
	return t.ReleaseDate
}
