package entertainment

import (
	"errors"
	"fmt"
	"simple-web-app/models"
)

type MediaInterface interface {
	Show() string
}

type MusicFromFactory struct {
	Entertainment *models.Song
}

func (mff *MusicFromFactory) Show() string {
	return fmt.Sprintf("Media type is: %s", mff.Entertainment.MediaType.Name)
}

type TelevisionFromFactory struct {
	Entertainment *models.TVShow
}

func (tff *TelevisionFromFactory) Show() string {
	return fmt.Sprintf("Media Type is: %s", tff.Entertainment.MediaType.Name)
}

type EntertainmentFactory interface {
	newMedia() MediaInterface
}

type MusicAbstractFactory struct{}

func (maf *MusicAbstractFactory) newMedia() MediaInterface {
	return &MusicFromFactory{
		Entertainment: &models.Song{},
	}
}

type TelevisionAbstractFactory struct{}

func (taf *TelevisionAbstractFactory) newMedia() MediaInterface {
	return &TelevisionFromFactory{
		Entertainment: &models.TVShow{},
	}
}

func NewEntertainmentFromAbstractFactory(media string) (MediaInterface, error) {
	switch media {
	case "music":
		var musicFactory MusicAbstractFactory
		music := musicFactory.newMedia()
		return music, nil
	case "television":
		var televisionFactory TelevisionAbstractFactory
		television := televisionFactory.newMedia()
		return television, nil
	default:
		return nil, errors.New("invalid media supplied")
	}
}
