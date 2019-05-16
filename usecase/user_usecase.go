package usecase

import "go-clean-architure/models"

type UserUseCase interface {
	Create(user *models.User) (error)
	GetByEmail(email string) (*models.User, error)
}