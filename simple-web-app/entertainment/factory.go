package entertainment

import "simple-web-app/models"

func NewEntertainment(media string) *models.Entertainment {
	entertainment := models.Entertainment{
		Media:       media,
		Type:        "Entertainment",
		Description: "no description entered yet",
	}

	return &entertainment
}
