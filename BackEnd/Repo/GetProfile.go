package repo

import (
	allqueries "emergency-backend/AllQueries"
	"emergency-backend/db"
)

func GetProfile(mobile string) (map[string]interface{}, error) {

	var email, role, dbMobile string

	err := db.Conn.QueryRow(allqueries.GetProfileDetails, mobile).Scan(&email, &role, &dbMobile)
	if err != nil {
		return nil, err
	}

	response := map[string]interface{}{
		"email":  email,
		"role":   role,
		"mobile": dbMobile,
	}

	return response, nil
}
