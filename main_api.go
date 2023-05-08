//go:build api

package main

import (
	"log"
	"net/http"

	"vita/transport/api"
)

func main() {
	http.HandleFunc("/", api.LambdaHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
