package usecaseimpl

import (
	"go-clean-architure/models"
	usecase "go-clean-architure/usecase"
	repo "go-clean-architure/repositories"
)

type UserUseCaseImpl struct {
	repo repo.UserRepo
}

func NewUserUsecase(repo repo.UserRepo) usecase.UserUseCase {
	return &UserUseCaseImpl {
		repo: repo,
	}
}

func (u *UserUseCaseImpl) Create(user *models.User) (error) {
	return u.repo.Create(user)
}

func (u *UserUseCaseImpl)GetByEmail(email string) (*models.User, error) {
	return u.repo.GetByEmail(email)
}