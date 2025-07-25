package repositories

import (
	"jwt_user/config"
	"jwt_user/models"
)

type UserRepository struct{}

func (r *UserRepository) Create(user *models.User) error {
	return config.DB.Create(user).Error
}

func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := config.DB.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) Update(user *models.User) error {
	return config.DB.Save(user).Error
}

func (r *UserRepository) Delete(id uint) error {
	return config.DB.Delete(&models.User{}, id).Error
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := config.DB.Find(&users).Error
	return users, err
}
