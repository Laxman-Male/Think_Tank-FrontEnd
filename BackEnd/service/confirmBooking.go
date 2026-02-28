package service

import (
	allstruct "emergency-backend/AllStruct"
	repo "emergency-backend/Repo"
)

// ConfirmBooking calls the repo function directly
func ConfirmBooking(patient allstruct.PatientData, userID int) error {
	return repo.ConfirmBooking(patient, userID)
}
