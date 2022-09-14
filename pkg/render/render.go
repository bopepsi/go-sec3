package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// serve html files
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsed, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsed.Execute(w, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
