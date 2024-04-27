package model

type User struct {
	Model
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
