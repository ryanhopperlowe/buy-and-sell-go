package user

import (
	"github.com/ryanhopperlowe/buy-and-sell-go/model"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *User) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserById(id model.Identifier) (*User, error)
	GetUsers() ([]User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) CreateUser(user *User) (*User, error) {
	result := r.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *repository) GetUserByEmail(email string) (*User, error) {
	var user User

	result := r.db.First(&user, &User{Email: email})

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *repository) GetUserById(id model.Identifier) (*User, error) {
	var user User

	result := r.db.First(&user, &model.Model{ID: id})

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *repository) GetUsers() ([]User, error) {
	var users []User

	result := r.db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
