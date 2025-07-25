package services

import (
	"jwt_user/models"
	"jwt_user/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repositories.UserRepository{},
	}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.userRepo.Create(user)
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.userRepo.Update(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.FindAll()
}
