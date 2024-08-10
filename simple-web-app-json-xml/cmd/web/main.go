package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"simple-web-app/configuration"
	"time"
)

const port = ":8081"

// receiver for handlers
type application struct {
	config appConfig
	App    *configuration.Application
}

type appConfig struct {
	useCache bool
	dsn      string
}

func main() {

	var config appConfig

	flag.StringVar(&config.dsn, "dsn", "root:@tcp(localhost:3306)/media_stream?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s", "Use DSN")
	flag.Parse()

	// get db
	db, err := initMySQLDB(config.dsn)
	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		// return and populates its fields
		App:    configuration.New(db),
		config: config,
	}

	// create a new server
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
