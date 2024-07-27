package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// save in memory so do not need to render every request
type templateData struct {
	Data map[string]any
}

// add pointer *templateData because if it does not have any data, will pass nil
func (app *application) render(writer http.ResponseWriter, t string, td *templateData) {
	var tmpl *template.Template
	// get the template cache from the app config
	// a map of template sets (stores information), with the template name as the key and the compiled template set as the value
	// the template set is a map of file names to *template.Template

	// if you use template cache, try to get template from map, stored in the receiver
	if app.config.useCache {
		if templateFromMap, ok := app.templateMap[t]; ok {
			tmpl = templateFromMap
		}
	}

	if tmpl == nil {
		// if not use template cache, parse the template from the file system every time
		newTemplate, err := app.buildTemplateFromDisk(t)
		if err != nil {
			log.Println("error building template", err)
			return
		}
		log.Println("template built from disk")
		tmpl = newTemplate
	}

	if td == nil {
		td = &templateData{}
	}

	if err := tmpl.ExecuteTemplate(writer, t, td); err != nil {
		log.Println("error executing template", err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	// initialize a new buffer

	// execute the template set, passing the current request to the template

	// write the contents of the buffer to the http.ResponseWriter
}

func (app *application) buildTemplateFromDisk(t string) (*template.Template, error) {
	// create a new template set
	templateSlice := []string{
		"./templates/base.layout.gohtml",
		"./templates/partials/header.partial.gohtml",
		"./templates/partials/footer.partial.gohtml",
		fmt.Sprintf("./templates/%s", t),
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		return nil, err
	}

	// if flag cache == true, use template cached instead from disk
	app.templateMap[t] = tmpl

	// return the template set
	return tmpl, nil
}
