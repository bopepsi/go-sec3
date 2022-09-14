package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

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

func main()  {
	
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Fatal(http.ListenAndServe(portNumber,nil))

}