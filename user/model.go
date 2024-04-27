package user

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ryanhopperlowe/buy-and-sell-go/model"
)

type User struct {
	model.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
	Role     string `json:"-"`
}

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(email string, passwordHash string) *User {
	return &User{
		Email:    email,
		Password: passwordHash,
		Role:     "user",
	}
}

type Claims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}
