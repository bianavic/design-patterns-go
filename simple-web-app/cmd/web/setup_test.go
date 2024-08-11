package main

import (
	"os"
	"simple-web-app/cmd/web/adapters"
	"simple-web-app/configuration"
	"testing"
)

// set testing environment
var testApp application

func TestMain(m *testing.M) {

	testBackend := &adapters.TestBackend{}
	testAdapter := adapters.RemoteService{MRemote: testBackend}

	testApp = application{
		App: configuration.New(nil, &testAdapter),
	}

	os.Exit(m.Run())
}
