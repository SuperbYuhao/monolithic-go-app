package repositories

import (
	"monolithic-go-app/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *UserRepository) FindByID(id uint) (models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	return user, result.Error
}

func (r *UserRepository) Create(user *models.User) error {
	result := r.db.Create(user)
	return result.Error
}

func (r *UserRepository) Update(user *models.User) error {
	result := r.db.Save(user)
	return result.Error
}

func (r *UserRepository) Delete(id uint) error {
	result := r.db.Delete(&models.User{}, id)
	return result.Error
}
