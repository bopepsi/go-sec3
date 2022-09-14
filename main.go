package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World")
		if err != nil {
			fmt.Println("Err when handle req", err)
		}
		fmt.Println("Bytes sent", n)
	})

	log.Fatal(http.ListenAndServe(":8080",nil))
}