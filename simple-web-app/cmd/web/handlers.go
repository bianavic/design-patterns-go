package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (app *application) ShowHome(writer http.ResponseWriter, request *http.Request) {
	// because it is a pointer to templateData, can pass nil
	app.render(writer, "home.page.gohtml", nil)
}

func (app *application) ShowPage(writer http.ResponseWriter, request *http.Request) {
	page := chi.URLParam(request, "page")
	app.render(writer, fmt.Sprintf("%s.page.gohtml", page), nil)
	if page == "" {
		page = "home"
	}

	//if page != "home" && page != "about" && page != "contact" {
	//	app.notFound(writer)
	//	return
	//}

}
