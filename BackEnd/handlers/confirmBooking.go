package handlers

import (
	allstruct "emergency-backend/AllStruct"
	"emergency-backend/middleware"
	"strconv"

	"emergency-backend/service"
	"encoding/json"
	"net/http"
)

// POST /confirm-booking
func ConfirmBookingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	userid, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Decode the incoming JSON payload
	var req struct {
		PatientData allstruct.PatientData `json:"patientData"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userid)

	// Save to database using repo directly
	// err = service.ConfirmBooking(req.PatientData)
	err = service.ConfirmBooking(req.PatientData, userID)
	if err != nil {
		http.Error(w, "Failed to store booking", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Booking stored",
	})
}
