package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/irisida/gobnb/pkg/config"
	"github.com/irisida/gobnb/pkg/models"
)

// map of functions that can be used in a template
var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData accepts and returns a templatedata type. This routine
// should be used to add any required default data.
func AddDefaultData(templateData *models.TemplateData) *models.TemplateData {
	return templateData
}

// renderTemplate renders temples using html.template
func RenderTemplate(w http.ResponseWriter, tmpl string, templateData *models.TemplateData) {
	var templateCache map[string]*template.Template

	// get the templateCache from the app.config when
	// useCache is set to true in the app config
	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	t, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("could not get template for the templateCache")
	}

	// create a buffer, prepare the data that will be inserted by
	// passing it to the default data handler and execute the template.
	buf := new(bytes.Buffer)
	templateData = AddDefaultData(templateData)
	_ = t.Execute(buf, templateData)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
}

// CreateTemplateCache creates a template cache as a map. To create the map
// the function collects the names of files matching the page.tmpl extension
// from the templates directory and creates a slice of string.
// For each page template we are checking for presence of functions required
// processing and parsing.
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = templateSet
	}

	return myCache, nil
}
