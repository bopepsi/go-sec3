package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/bopepsi/go-app/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	parsed, _ := template.ParseFiles("./templates/about.page.html")
	err := parsed.Execute(w, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
