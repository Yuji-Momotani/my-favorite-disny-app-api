package repository

import (
	"my-favorite-disney-app-api/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user *model.User) error
	GetUserByUserName(user *model.User, name string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUserByUserName(user *model.User, name string) error {
	if err := ur.db.Where("name = ?", name).First(user).Error; err != nil {
		return err
	}
	return nil
}
