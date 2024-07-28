package main

import (
	"os"
	"simple-web-app/configuration"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {
	testApp = application{
		App: configuration.New(nil),
	}

	os.Exit(m.Run())
}
