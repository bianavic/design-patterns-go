package entertainment

import (
	"errors"
	"fmt"
	"simple-web-app/cmd/web/adapters"
	"simple-web-app/configuration"
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
	NewEntertainmentWithMusicAlbums(music string) MediaInterface
	NewEntertainmentWithTVShows(tvShow string) MediaInterface
}

type MusicAbstractFactory struct{}

func (maf *MusicAbstractFactory) newMedia() MediaInterface {
	return &MusicFromFactory{
		Entertainment: &models.Song{},
	}
}

func (maf *MusicAbstractFactory) NewEntertainmentWithMusicAlbums(music string) MediaInterface {
	app := configuration.GetInstance()
	album, _ := app.Models.MusicAlbum.GetMusicAlbumByName(music)
	return &MusicFromFactory{
		Entertainment: &models.Song{
			// because Cannot use 'album' (type *MusicAlbum) as the type MusicAlbum
			// DEFERENCE the pointer to get the value - MediaType: *album
			MediaType: *album,
		},
	}
}

type TelevisionAbstractFactory struct{}

func (taf *TelevisionAbstractFactory) newMedia() MediaInterface {
	return &TelevisionFromFactory{
		Entertainment: &models.TVShow{},
	}
}

func (taf *TelevisionAbstractFactory) NewEntertainmentWithTVShows(tvShow string) MediaInterface {
	app := configuration.GetInstance()
	tv, _ := app.Models.Television.GetTelevisionByName(tvShow)
	return &TelevisionFromFactory{
		Entertainment: &models.TVShow{
			MediaType: *tv,
		},
	}
}

//func NewEntertainmentFromAbstractFactory(media string) (MediaInterface, error) {
//	switch media {
//	case "music":
//		var musicFactory MusicAbstractFactory
//		music := musicFactory.newMedia()
//		return music, nil
//	case "television":
//		var televisionFactory TelevisionAbstractFactory
//		television := televisionFactory.newMedia()
//		return television, nil
//	default:
//		return nil, errors.New("invalid media supplied")
//	}
//}

// refactor
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

func NewEntertainmentTypesFromAbstractFactory(media, music string, tvShow string, adapter *adapters.RemoteService) (MediaInterface, error) {
	switch media {
	case "music":
		var musicFactory MusicAbstractFactory
		album := musicFactory.NewEntertainmentWithMusicAlbums(music)
		// return music with music album embedded
		return album, nil
	case "television":
		var televisionFactory TelevisionAbstractFactory
		tv := televisionFactory.NewEntertainmentWithTVShows(tvShow)
		// return television with tv show embedded
		return tv, nil

	default:
		return nil, errors.New("invalid medias supplied")
	}
}
