package service

import (
	repo "emergency-backend/Repo"
)

func GetProfile(userID string) (map[string]interface{}, error) {

	response, err := repo.GetProfile(userID)
	if err != nil {
		return nil, err
	}

	return response, nil
}
