package main

import (
	"fmt"
	"log"
	"net/http"
)

// main server sends our basic html page to the browser, duh?
// also I guess we are handling auth through here, frick.
func main() {
	fs := http.FileServer(http.Dir("public"))

	http.Handle("/", fs)

	fmt.Println("server starting up...")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("error happened while starting server")
	}
	fmt.Println("server started on localhost:8000")
}
