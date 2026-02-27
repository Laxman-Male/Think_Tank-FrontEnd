package repo

import (
	allqueries "emergency-backend/AllQueries"
	"emergency-backend/db"
)

type Hospital struct {
	ID                  int     `json:"id"`
	Name                string  `json:"hospital_name"`
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	AvailableAmbulances int     `json:"available_ambulances"`
	Distance            float64 `json:"distance"`
}

func GetNearestHospital(userLat, userLng float64) ([]Hospital, error) {

	rows, err := db.Conn.Query(allqueries.GetNearestHospital, userLat, userLng, userLat)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hospitals []Hospital

	for rows.Next() {
		var h Hospital
		err := rows.Scan(
			&h.ID,
			&h.Name,
			&h.Latitude,
			&h.Longitude,
			&h.AvailableAmbulances,
			&h.Distance,
		)
		if err != nil {
			return nil, err
		}
		hospitals = append(hospitals, h)
	}

	return hospitals, nil
}
