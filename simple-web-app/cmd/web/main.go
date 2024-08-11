package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"simple-web-app/cmd/web/adapters"
	"simple-web-app/configuration"
	"time"
)

const port = ":4000"

// receiver for handlers
type application struct {
	templateMap   map[string]*template.Template
	config        appConfig
	App           *configuration.Application
	remoteService *adapters.RemoteService
}

type appConfig struct {
	useCache bool
	dsn      string
}

func main() {

	app := &application{
		// return and populates its fields
		templateMap: make(map[string]*template.Template),
	}

	flag.BoolVar(&app.config.useCache, "cache", false, "Use template cache")
	flag.StringVar(&app.config.dsn, "dsn", "root:@tcp(localhost:3306)/media_stream?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s", "Use DSN")
	flag.Parse()

	// get db
	db, err := initMySQLDB(app.config.dsn)
	if err != nil {
		log.Panic(err)
	}

	//jsonBackend := adapters.JSONBackend{}
	//jsonAdapter := adapters.RemoteService{Remote: &jsonBackend}

	xmlBackend := adapters.XMLBackend{}
	xmlAdapter := adapters.RemoteService{Remote: &xmlBackend}

	app.App = configuration.New(db)
	//app.remoteService = &jsonAdapter
	app.remoteService = &xmlAdapter

	server := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	fmt.Println("Server running on port", port)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
