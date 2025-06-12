package auth

import (
	"spliteasy/internal/database"
	"spliteasy/internal/models"
)

func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("email = ?", email).Find(&user)
	return &user, result.Error
}
