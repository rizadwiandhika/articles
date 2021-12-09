package database

import (
	"rizadwi.com/config"
	"rizadwi.com/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
