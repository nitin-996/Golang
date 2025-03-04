package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// UserHandler extracts and prints the "id" from the URL
func UserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract path variables from the request
	vars := mux.Vars(r)
	id := vars["id"] // Get the "id" parameter from the URL

	// Send a response with the extracted ID
	fmt.Fprintf(w, "User ID: %s", id)
}

func main() {
	r := mux.NewRouter()

	// Route with a path variable {id}
	r.HandleFunc("/user/{id}", UserHandler).Methods("GET")

	// Start server
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
