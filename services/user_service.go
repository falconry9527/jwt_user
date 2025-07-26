package services

import (
	"jwt_user/db"
	"jwt_user/models"
)

func CreateUser(user *models.User) error {
	return db.DB.Create(user).Error
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := db.DB.First(&user, id).Error
	return &user, err
}

func UpdateUser(user *models.User) error {
	return db.DB.Save(user).Error
}

func DeleteUser(id uint) error {
	return db.DB.Delete(&models.User{}, id).Error
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := db.DB.Find(&users).Error
	return users, err
}
