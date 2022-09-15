package main

import (
	"log"
	"net/http"

	"github.com/bopepsi/go-app/pkg/config"
	"github.com/bopepsi/go-app/pkg/handlers"
	"github.com/bopepsi/go-app/pkg/render"
)

const portNumber = ":8080"

func main() {

	// Setup app wide config
	var app config.AppConfig
	cache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = cache
	app.UseCache = true

	// Setup handlers repo and render template cache
	handlers.SetupRepo(&app)
	render.SetupTmplCacheMap(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	log.Fatal(http.ListenAndServe(portNumber, nil))

}
