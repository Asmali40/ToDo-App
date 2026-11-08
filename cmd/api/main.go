package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("[INFO] Starting server...")

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatalf("[ERROR] Server stopped %v", err)
	}
}
