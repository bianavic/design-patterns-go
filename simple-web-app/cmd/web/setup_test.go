package main

import (
	"log"
	"os"
	"simple-web-app/models"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {

	dsn := "root:@tcp(localhost:3306)/media_stream?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s"
	db, err := initMySQLDB(dsn)
	if err != nil {
		log.Panic(err)
	}

	testApp = application{
		DB:     db,
		Models: *models.New(db),
	}

	os.Exit(m.Run())
}
