package models

type MusicAlbum struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Genre       string `json:"genre"`
	Details     string `json:"details"`
	Country     string `json:"country"`
	ReleaseDate string `json:"release_date"`
}

type Television struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Genre       string `json:"genre"`
	Details     string `json:"details"`
	Country     string `json:"country"`
	ReleaseDate string `json:"release_date"`
}

type MusicTrack struct {
	ID           int        `json:"id"`
	MusicAlbumID int        `json:"music_album_id"`
	Year         int        `json:"year"`
	PersonID     int        `json:"person_id"`
	Title        string     `json:"title"`
	Artist       string     `json:"artist"`
	Music        MusicAlbum `json:"music_album"`
	Person       []*Person  `json:"person"`
}

type TVShow struct {
	ID           int        `json:"id"`
	TelevisionID int        `json:"television_id"`
	Year         int        `json:"year"`
	PersonID     int        `json:"person_id"`
	Title        string     `json:"title"`
	Director     string     `json:"director"`
	Television   Television `json:"television"`
	Person       []*Person  `json:"person"`
}

type Person struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Address     string        `json:"address"`
	Country     string        `json:"country"`
	ZipCode     string        `json:"zip_code"`
	Email       string        `json:"email"`
	Active      int           `json:"active"`
	MusicTracks []*MusicTrack `json:"music_tracks"`
	TVShows     []*TVShow     `json:"tv_shows"`
}

type Entertainment struct {
	ID          int    `json:"id"`
	Media       string `json:"media"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
