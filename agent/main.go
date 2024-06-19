package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/manage", manageHandler)

	port := "8081"
	log.Printf("Agent Service is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
