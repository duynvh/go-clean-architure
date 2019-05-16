package repoimpl

import (
	"github.com/jinzhu/gorm"
	"go-clean-architure/models"
	repo "go-clean-architure/repositories"
)

type UserRepoImpl struct {
	Db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repo.UserRepo {
	return &UserRepoImpl {
		Db: db,
	}
}

func (u *UserRepoImpl) Create(user *models.User) (error) {
	if err := u.Db.New().Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepoImpl)GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := u.Db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}