package services

import (
	"monolithic-go-app/internal/models"
	"monolithic-go-app/internal/repositories"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}

type userService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.FindAll()
}

func (s *userService) GetUserByID(id uint) (models.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *userService) CreateUser(user *models.User) error {
	return s.userRepo.Create(user)
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.userRepo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}
