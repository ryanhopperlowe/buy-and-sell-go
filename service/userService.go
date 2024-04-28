package service

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ryanhopperlowe/buy-and-sell-go/model"
	"github.com/ryanhopperlowe/buy-and-sell-go/repo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	r *repo.UserRepository
}

func NewService(r *repo.UserRepository) *UserService {
	return &UserService{r}
}

func (s *UserService) Signup(ctx *gin.Context) {
	var body model.SignupRequest

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user, err := s.r.CreateUser(model.NewUser(body.Email, string(hash)))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (s *UserService) Login(ctx *gin.Context) {
	var body model.LoginRequest

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	user, err := s.r.GetUserByEmail(body.Email)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Password"})
		return
	}

	claims := model.DefaultClaims(strconv.FormatUint(uint64(user.ID), 10))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "/", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{})
}

func (s *UserService) ValidateToken(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Token is valid"})
}

func (s *UserService) GetUserById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id"})
		return
	}

	user, err := s.r.GetUserById(model.Identifier(id))

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (s *UserService) GetUsers(ctx *gin.Context) {
	users, err := s.r.GetUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
