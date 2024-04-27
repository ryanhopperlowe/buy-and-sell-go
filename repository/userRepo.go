package repository

import (
	"github.com/ryanhopperlowe/buy-and-sell-go/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(user *model.User) (*model.User, error) {
	result := r.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	result := r.db.First(&user, &model.User{Email: email})

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepository) GetUserById(id model.Identifier) (*model.User, error) {
	var user model.User

	result := r.db.First(&user, &model.Model{ID: id})

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepository) GetUsers() ([]model.User, error) {
	var users []model.User

	result := r.db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
