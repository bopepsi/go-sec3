package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request)  {
	RenderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request)  {
	parsed,_ := template.ParseFiles("./templates/about.page.html")
	err := parsed.Execute(w, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// serve html files
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsed, _ := template.ParseFiles("./templates/"+tmpl)
	err := parsed.Execute(w, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}