package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/scale", scaleHandler)

	port := "8082"
	log.Printf("Controller Service is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
