package api

import (
	"encoding/json"
	"net/http"
	"myapi/db"
	"myapi/models"
	//"github.com/bxcodec/faker/v3"
)

// CreateProduct creates a new product and saves it to the database
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request body into a Mobile struct
	var newProduct models.Mobile
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	// Create a database connection
	db := db.GetDB()
	if db == nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}

	// Save the new product to the database
	db.Create(&newProduct)

	// Return the created product as JSON response
	response, err := json.Marshal(newProduct)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Set content type header
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(response)
}

// GetMobiles fetches and returns the list of mobiles from the database
func GetMobiles(w http.ResponseWriter, r *http.Request) {
	// Create a database connection
	db := db.GetDB()
	if db == nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}

	// Fetch all mobiles from the database
	var mobiles []models.Mobile
	db.Find(&mobiles)

	// Return the list of mobiles as JSON response
	response, err := json.Marshal(mobiles)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Set content type header
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(response)
}

// Handler function to handle incoming HTTP requests
func Handler(w http.ResponseWriter, r *http.Request) {
	// Create a database connection
	db := db.GetDB()
	if db == nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost && r.URL.Path == "/create" {
		CreateProduct(w, r)
		return
	}

	// For example, you can call GetMobiles when a GET request to /mobiles is received
	if r.Method == http.MethodGet && r.URL.Path == "/mobiles" {
		GetMobiles(w, r)
		return
	}

	// Migrate the models (not necessary if done in db.Init())
	// db.AutoMigrate(&models.Mobile{})

	// Generate a slice of 5 random Mobiles
	var mobiles []models.Mobile

	// Convert slice to JSON
	response, err := json.Marshal(mobiles)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Set content type header
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(response)
}
