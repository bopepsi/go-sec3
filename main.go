package main

import (
	"log"
	"net/http"
)

const portNumber = ":8080"

func main()  {
	
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Fatal(http.ListenAndServe(portNumber,nil))

}