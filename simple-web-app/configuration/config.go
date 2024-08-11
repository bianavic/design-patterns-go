package configuration

import (
	"database/sql"
	"simple-web-app/cmd/web/adapters"
	"simple-web-app/models"
	"sync"
)

type Application struct {
	Models        *models.Models
	remoteService *adapters.RemoteService
	AlbumService  *adapters.RemoteService
	TVService     *adapters.RemoteService
	TVShowService *adapters.RemoteService
}

var instance *Application
var once sync.Once
var db *sql.DB
var musicAlbumService *adapters.RemoteService
var tvService *adapters.RemoteService
var tvShowService *adapters.RemoteService

func New(pool *sql.DB, rs *adapters.RemoteService) *Application {
	db = pool
	musicAlbumService = rs
	tvService = rs
	tvShowService = rs
	return GetInstance()
}

func GetInstance() *Application {
	once.Do(func() {
		instance = &Application{
			Models:        models.New(db),
			AlbumService:  musicAlbumService,
			TVService:     tvService,
			TVShowService: tvShowService,
		}
	})
	return instance
}
