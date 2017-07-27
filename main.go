package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	router := NewRouter()

	AddRedirect(Redirect{
		"abc",
		"https://google.com",
		true,
		time.Now(),
	})

	// datastore.Write("test", Redirect{})

	log.Fatal(http.ListenAndServe(":8080", router))
}
