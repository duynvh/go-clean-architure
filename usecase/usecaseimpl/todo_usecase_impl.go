package usecaseimpl

import (
	"go-clean-architure/models"
	usecase "go-clean-architure/usecase"
	repo "go-clean-architure/repositories"
)

type TodoUseCaseImpl struct {
	repo repo.TodoRepo
}

func NewTodoUsecase(repo repo.TodoRepo) usecase.TodoUseCase {
	return &TodoUseCaseImpl {
		repo: repo,
	}
}

func (u *TodoUseCaseImpl) Create(todo *models.Todo) (error) {
	return u.repo.Create(todo)
}

func (u *TodoUseCaseImpl)Update(todo *models.Todo) (error) {
	return u.repo.Update(todo)
}

func (u *TodoUseCaseImpl)Delete(id int) (error) {
	return u.repo.Delete(id)
}

func (u *TodoUseCaseImpl)FetchById(id int) (*models.Todo, error) {
	return u.repo.FetchById(id)
}

func (u *TodoUseCaseImpl)Fetch(page int) ([]*models.Todo, error) {
	return u.repo.Fetch(page)
}
