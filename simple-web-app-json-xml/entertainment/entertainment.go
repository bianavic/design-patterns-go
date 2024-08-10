package entertainment

type Entertainment struct {
	// music or tv
	Type        string `json: "type"`
	MediaTitle  string `json: "media_title"`
	Genre       string `json: "genre"`
	Description string `json: "description", omitempty`
	Country     string `json: "country", omitempty`
	Rating      int    `json: "rating", omitempty`
	ReleaseYear int    `json: "release_year"`
	Available   bool   `json: "available"`
}
