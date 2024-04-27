package listing

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ryanhopperlowe/buy-and-sell-go/model"
)

type Service interface {
	CreateListing(ctx *gin.Context)
	GetListings(ctx *gin.Context)
	GetListingById(ctx *gin.Context)
	UpdateListing(ctx *gin.Context)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) CreateListing(ctx *gin.Context) {
	var newListing CreateRequest

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

func (c *service) GetListings(ctx *gin.Context) {
	listings, err := c.r.GetListings()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, listings)
}

func (c *service) GetListingById(ctx *gin.Context) {

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

func (c *service) UpdateListing(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var listing Listing

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
