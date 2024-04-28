package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ryanhopperlowe/buy-and-sell-go/repo"
	"github.com/ryanhopperlowe/buy-and-sell-go/service"
	"gorm.io/gorm"
)

func MakeListingRoutes(r *gin.Engine, db *gorm.DB) {
	repo := repo.NewListingRepository(db)
	service := service.NewListingService(repo)

	r.GET("/listings", service.GetListings)
	r.GET("/listings/:id", service.GetListingById)

	r.POST("/listings", service.CreateListing)

	r.PUT("/listings/:id", service.UpdateListing)
}
