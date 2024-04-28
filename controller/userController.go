package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ryanhopperlowe/buy-and-sell-go/middleware"
	"github.com/ryanhopperlowe/buy-and-sell-go/repo"
	"github.com/ryanhopperlowe/buy-and-sell-go/service"
	"gorm.io/gorm"
)

func MakeUserRoutes(r *gin.Engine, db *gorm.DB) {
	repo := repo.NewUserRepository(db)
	service := service.NewService(repo)

	r.GET("/users", service.GetUsers)
	r.GET("/users/:id", service.GetUserById)
	r.GET("/validate", middleware.RequireAuth, service.ValidateToken)

	r.POST("/signup", service.Signup)
	r.POST("/login", service.Login)
}
