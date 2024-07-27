package models

// DogBreed
type MusicAlbum struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// Breed
	MediaType   string `json:"media_type"`
	Genre       string `json:"genre"`
	Details     string `json:"details"`
	Country     string `json:"country"`
	ReleaseDate string `json:"release_date"`
}

// CatBreed
type Television struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// Breed
	MediaType   string `json:"media_type"`
	Genre       string `json:"genre"`
	Details     string `json:"details"`
	Country     string `json:"country"`
	ReleaseDate string `json:"release_date"`
}

// Dog
type Song struct {
	ID           int    `json:"id"`
	MusicAlbumID int    `json:"music_album_id"`
	Year         int    `json:"year"`
	PersonID     int    `json:"person_id"`
	SongTitle    string `json:"song_title"`
	Artist       string `json:"artist"`
	// Breed
	MediaType MusicAlbum `json:"album"`
	// Breeder
	Person Person `json:"person"`
}

// Cat
type TVShow struct {
	ID           int    `json:"id"`
	TelevisionID int    `json:"television_id"`
	Year         int    `json:"year"`
	PersonID     int    `json:"person_id"`
	Title        string `json:"title"`
	Director     string `json:"director"`
	// 	Breed
	MediaType Television `json:"television"`
	// Breeder
	Person Person `json:"person"`
}

// Breeder
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

// Pet
type Entertainment struct {
	// Species
	Media string `json:"media"`
	// Breed
	MediaType   string `json:"media_type"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
