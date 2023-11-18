package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
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

	// Create a new gorilla/mux router
	router := mux.NewRouter()

	// Register the handler function for the /mobiles endpoint
	router.HandleFunc("/mobiles", api.Handler).Methods("GET")
	router.HandleFunc("/create", api.Handler).Methods("POST")

	// Enable CORS
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	// Apply CORS middleware to the router
	router.Use(cors)

	// Handle OPTIONS requests
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Start the server with CORS middleware
	http.ListenAndServe(":8080", router)
}
