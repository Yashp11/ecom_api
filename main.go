package main

import (
	"net/http"
	"myapi/api"
	"myapi/db"
)

func main() {
	// Initialize the database
	_, err := db.Init()
	if err != nil {
		panic(err)
	}
	defer db.GetDB().Close()

	// Register the handler function for the /mobiles endpoint
	http.HandleFunc("/mobiles", api.Handler)
	http.HandleFunc("/create", api.Handler)

	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", nil)
}
