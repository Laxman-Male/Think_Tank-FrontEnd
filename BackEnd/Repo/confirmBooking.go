package repo

import (
	"database/sql"
	allqueries "emergency-backend/AllQueries"
	allstruct "emergency-backend/AllStruct"
	"encoding/json"
	"log"
	"strconv"
)

// Assume DB is a global variable initialized in main.go
var DB *sql.DB

// SaveBooking stores patient data in the database
func ConfirmBooking(patient allstruct.PatientData, userID int) error {
	age, err := strconv.Atoi(patient.Age)
	if err != nil {
		return err
	}

	// Convert hospital struct to JSON
	hospitalJSON, err := json.Marshal(patient.Hospital)
	if err != nil {
		return err
	}

	_, err = DB.Exec(allqueries.ConfirmBooking,
		userID,
		patient.Name,
		age,
		patient.Gender,
		patient.Blood,
		patient.Type,
		hospitalJSON, // single JSON column
	)
	if err != nil {
		log.Println("DB insert error:", err)
		return err
	}

	return nil
}
