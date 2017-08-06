package main

import (
	"log"
	"net/http"
	"os"
)

// Environment variables

type ENV struct {
	DB_NAMESPACE string
	API_ROUTE    string
	REQUIRE_AUTH string
	AUTH_SECRET  string
}

var Env ENV

func main() {
	defaultEnv := ENV{
		"gurlin",
		"gurlin_api",
		"false",
		"",
	}

	Env = ENV{
		DB_NAMESPACE: getEnv("GURLIN_DB_NAMESPACE", defaultEnv.DB_NAMESPACE),
		API_ROUTE:    getEnv("GURLIN_API_ROUTE", defaultEnv.API_ROUTE),
		REQUIRE_AUTH: getEnv("GURLIN_REQUIRE_AUTH", defaultEnv.REQUIRE_AUTH),
		AUTH_SECRET:  getEnv("GURLIN_SECRET", defaultEnv.AUTH_SECRET),
	}

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if (len(value)) == 0 {
		return defaultValue
	}
	return value
}
