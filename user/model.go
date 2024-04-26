package user

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ryanhopperlowe/buy-and-sell-go/listing"
	"github.com/ryanhopperlowe/buy-and-sell-go/model"
)

type User struct {
	model.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"index:unique"`
	Password string `json:"-"`
	Role     string `json:"role"`

	Listings []listing.Listing `json:"listings" gorm:"foreignKey:UserId"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(data CreateUserRequest) *User {
	return &User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Role:     "user",
	}
}

type Claims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}
