package listing

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MakeRoutes(r *gin.Engine, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)

	r.GET("/listings", service.GetListings)
	r.GET("/listings/:id", service.GetListingById)

	r.POST("/listings", service.CreateListing)

	r.PUT("/listings/:id", service.UpdateListing)
}
