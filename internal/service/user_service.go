package service

import (
	"monolithic-go-app/internal/models"
	"monolithic-go-app/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.FindAll()
}

func (s *UserService) GetUserByID(id int) (models.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.userRepo.Create(user)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.userRepo.Update(user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.userRepo.Delete(id)
}
