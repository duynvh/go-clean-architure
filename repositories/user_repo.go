package repositories

import "go-clean-architure/models"

type UserRepo interface {
	Create(user *models.User) (error)
	GetByEmail(email string) (*models.User, error)
}