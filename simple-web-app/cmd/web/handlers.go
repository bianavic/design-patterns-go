package main

import (
	"net/http"
)

func (app *application) ShowHome(writer http.ResponseWriter, request *http.Request) {
	//fmt.Fprint(writer, "Hello, Stranger!")
	writer.Write([]byte("Hello, Stranger!"))
}
