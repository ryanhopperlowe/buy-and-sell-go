package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ryanhopperlowe/buy-and-sell-go/model"
	"github.com/ryanhopperlowe/buy-and-sell-go/repo"
)

type ListingService struct {
	r *repo.ListingRepository
}

func NewListingService(r *repo.ListingRepository) *ListingService {
	return &ListingService{r}
}

func (s *ListingService) CreateListing(ctx *gin.Context) {
	var newListing model.CreateListingRequest

	if err := ctx.BindJSON(&newListing); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	listing, err := s.r.CreateListing(newListing)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, listing)
}

func (c *ListingService) GetListings(ctx *gin.Context) {
	listings, err := c.r.GetListings()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, listings)
}

func (c *ListingService) GetListingById(ctx *gin.Context) {

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	listing, err := c.r.GetListingById(model.Identifier(id))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, listing)
}

func (c *ListingService) UpdateListing(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var listing model.Listing

	if err := ctx.BindJSON(&listing); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	listing.ID = model.Identifier(id)
	updated, err := c.r.UpdateListing(listing)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updated)
}
