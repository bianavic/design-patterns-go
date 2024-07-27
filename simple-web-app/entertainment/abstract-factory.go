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
	Entertainment *models.MusicTrack
}

func (mff *MusicFromFactory) Show() string {
	return fmt.Sprintf("Music album name is: %s", mff.Entertainment.MusicAlbum.Name)
}

type TelevisionFromFactory struct {
	Entertainment *models.TVShow
}

func (tff *TelevisionFromFactory) Show() string {
	return fmt.Sprintf("TV Show name is: %s", tff.Entertainment.Television.Name)
}

type EntertainmentFactory interface {
	newMedia() MediaInterface
}

type MusicAbstractFactory struct{}

func (maf *MusicAbstractFactory) newMedia() MediaInterface {
	return &MusicFromFactory{
		Entertainment: &models.MusicTrack{},
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
