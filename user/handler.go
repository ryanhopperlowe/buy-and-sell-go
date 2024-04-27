package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MakeRoutes(r *gin.Engine, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)

	r.GET("/users", service.GetUsers)
	r.GET("/users/:id", service.GetUserById)

	r.POST("/signup", service.Signup)
	r.POST("/login", service.Login)
}
