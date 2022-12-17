package main

import (
	"net/http"
)

func spa(responseWriter http.ResponseWriter, request *http.Request) {
	http.ServeFile(responseWriter, request, "./spa/build/index.html")
}

// EXAMPLE FUNCTION

// func items(w http.ResponseWriter, r *http.Request) {
// 	data := getItems()
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(data)
// }
