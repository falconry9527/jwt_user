package services

import (
	"jwt_user/db"
	"jwt_user/models"
)

func CreateUser(user *models.User) error {
	return db.Mysql.Create(user).Error
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := db.Mysql.First(&user, id).Error
	return &user, err
}

func UpdateUser(user *models.User) error {
	return db.Mysql.Save(user).Error
}

func DeleteUser(id uint) error {
	return db.Mysql.Delete(&models.User{}, id).Error
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := db.Mysql.Find(&users).Error
	return users, err
}
