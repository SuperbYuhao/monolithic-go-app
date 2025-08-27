package repository

import (
	"database/sql"
	"monolithic-go-app/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	rows, err := r.db.Query("SELECT id, name, email, created_at, updated_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) FindByID(id int) (models.User, error) {
	var user models.User
	err := r.db.QueryRow("SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) Create(user *models.User) error {
	err := r.db.QueryRow(
		"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, created_at, updated_at",
		user.Name, user.Email).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	return err
}

func (r *UserRepository) Update(user *models.User) error {
	_, err := r.db.Exec(
		"UPDATE users SET name = $1, email = $2, updated_at = NOW() WHERE id = $3",
		user.Name, user.Email, user.ID)
	return err
}

func (r *UserRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
