package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// renderTemplate is responsible for parsing and executing
// a passed template which writes out the template to the
// responseWriter that was passed in.
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template", err)
	}
}
