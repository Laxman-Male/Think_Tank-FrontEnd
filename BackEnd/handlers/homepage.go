package handlers

import (
	"fmt"
	"log"
	"net/http"

	"emergency-backend/db"
)

// Homepage prints hello and inserts dummy data into the database
func Homepage(w http.ResponseWriter, r *http.Request) {
	// Print hello to terminal
	fmt.Println("Hello! Request received at homepage")

	// Insert dummy data into the app table
	query := `INSERT INTO app (user_name, service_name, status) VALUES (?, ?, ?)`

	result, err := db.Conn.Exec(query, "John Doe", "Emergency Response System", "Pending")
	if err != nil {
		log.Printf("Error inserting data: %v", err)
		http.Error(w, "Failed to insert data", http.StatusInternalServerError)
		return
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID: %v", err)
	}

	log.Printf("Data inserted successfully with ID: %d", lastID)

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Hello! Data inserted successfully","id":` + fmt.Sprintf("%d", lastID) + `}`))
}
