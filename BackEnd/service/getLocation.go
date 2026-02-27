package service

import repo "emergency-backend/Repo"

// import "emergency-backend/repo"

// func FindNearestHospital(lat, lng float64) (interface{}, error) {
// 	return repo.GetNearestHospital(lat, lng)
// }

func FindNearestHospital(lat, lng float64) (interface{}, error) {
	return repo.GetNearestHospital(lat, lng)
}
