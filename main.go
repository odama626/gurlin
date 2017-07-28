package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	router := NewRouter()

	// for i := 0; i < 65*3; i++ {
	// 	fmt.Println(ItoS(i))
	// }
	// fmt.Println()

	AddRedirect(Redirect{
		"abc",
		"https://google.com",
		true,
		time.Now(),
	})

	// datastore.Write("test", Redirect{})

	log.Fatal(http.ListenAndServe(":8080", router))
}
