package handlers

import (
	allstruct "emergency-backend/AllStruct"
	"emergency-backend/service"
	"encoding/json"
	"net/http"
)

func FindNearestHospital(w http.ResponseWriter, r *http.Request) {

	var req allstruct.LocationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	response, err := service.FindNearestHospital(req.Latitude, req.Longitude)
	if err != nil {
		http.Error(w, "Error finding hospital", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
