package entertainment

import (
	"fmt"
	"log"
	"simple-web-app/configuration"
	"simple-web-app/models"
)

type MediaInterface interface {
	Show() string
}

type EntertainmentFactoryInterface interface {
	newMedia() MediaInterface
	newMediaWithMediaType(media string) MediaInterface
}

//type MusicFromFactory struct {
//	Entertainment *models.Song
//}

type TelevisionFromFactory struct {
	Entertainment *models.Television
}

type MusicAbstractFactory struct{}

type TelevisionAbstractFactory struct{}

//func (mff *MusicFromFactory) Show() string {
//	return fmt.Sprintf("Media type is: %s", mff.Entertainment.MediaType.Name)
//}

func (tff *TelevisionFromFactory) Show() string {
	return fmt.Sprintf("Media Type is: %s", tff.Entertainment.MediaType.Title)
}

//func (maf *MusicAbstractFactory) newMedia() MediaInterface {
//	return &MusicFromFactory{
//		Entertainment: &models.Song{},
//	}
//}

func (mff *TelevisionAbstractFactory) newMedia() MediaInterface {
	return &TelevisionFromFactory{
		Entertainment: &models.Television{},
	}
}

//func (mff *MusicAbstractFactory) newMediaWithMediaType(mt string) MediaInterface {
//	app := configuration.GetInstance()
//	mediaType, err := app.Models.MusicAlbum.GetMusicAlbumByName(mt)
//	if err != nil {
//		log.Println(err)
//		return nil
//	}
//	return &MusicFromFactory{
//		Entertainment: &models.Song{
//			MediaType: *mediaType,
//		},
//	}
//}

func (taf *TelevisionAbstractFactory) newMediaWithMediaType(mt string) MediaInterface {
	app := configuration.GetInstance()
	mediaType, err := app.Models.TVShow.GetTVShowByName(mt)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &TelevisionFromFactory{
		Entertainment: &models.Television{
			MediaType: *mediaType,
		},
	}
}

// refactor
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

//func NewEntertainmentMediaTypeFromAbstractFactory(entertainment, music string, tvShow string, adapter *adapters.RemoteService) (MediaInterface, error) {
//	switch entertainment {
//	case "music":
//		var musicFactory MusicAbstractFactory
//		album := musicFactory.newMediaWithMediaType(music)
//		// return music with music album embedded
//		return album, nil
//	case "television":
//		var televisionFactory TelevisionAbstractFactory
//		tv := televisionFactory.newMediaWithMediaType(tvShow)
//		// return television with tv show embedded
//		return tv, nil
//
//	default:
//		return nil, errors.New("invalid medias supplied")
//	}
//}
