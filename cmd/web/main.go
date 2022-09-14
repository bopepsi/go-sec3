package main

import (
	"log"
	"net/http"

	"github.com/bopepsi/go-app/pkg/handlers"
)

const portNumber = ":8080"

func main()  {
	
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Fatal(http.ListenAndServe(portNumber,nil))

}