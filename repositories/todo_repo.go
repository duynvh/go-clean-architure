package repositories

import "go-clean-architure/models"

type TodoRepo interface {
	Create(todo *models.Todo) (error)
	Fetch(page int) ([]*models.Todo, error)
	Update(todo *models.Todo) (error)
	Delete(id int) (error)
	FetchById(id int) (*models.Todo, error)
}