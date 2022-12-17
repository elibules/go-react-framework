package main

import "net/http"

func routes() {
	http.HandleFunc("/", spa)

	// EXAMPLE CALL TO SERVER
	// http.HandleFunc("/api/items", items)
}
