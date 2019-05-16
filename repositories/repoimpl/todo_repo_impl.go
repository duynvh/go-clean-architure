package repoimpl

import (
	"github.com/jinzhu/gorm"
	"go-clean-architure/models"
	repo "go-clean-architure/repositories"
	"os"
	"strconv"
)

type TodoRepoImpl struct {
	Db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) repo.TodoRepo {
	return &TodoRepoImpl {
		Db: db,
	}
}

func (u *TodoRepoImpl) Create(todo *models.Todo) (error) {
	if err := u.Db.New().Create(&todo).Error; err != nil {
		return err
	}
	return nil
}

func (u *TodoRepoImpl)Fetch(page int) ([]*models.Todo, error) {
	limit, _ := strconv.Atoi(os.Getenv("DB_LIMIT"))
	offset := (page - 1) * limit
	var todos []*models.Todo
	if err := u.Db.Offset(offset).Limit(limit).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (u *TodoRepoImpl) Update(todo *models.Todo) (error) {
	if err := u.Db.Save(&todo).Error; err != nil {
		return err
	}
	return nil
}

func (u *TodoRepoImpl) Delete(id int) (error) {
	todo, _ := u.FetchById(id)

	if err := u.Db.Delete(&todo).Error; err != nil {
		return err
	}
	return nil
}

func (u *TodoRepoImpl) FetchById(id int) (*models.Todo, error) {
	var todo models.Todo

	if err := u.Db.New().First(&todo, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &todo, nil
}