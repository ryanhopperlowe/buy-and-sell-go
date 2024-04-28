package repository

import (
	"github.com/ryanhopperlowe/buy-and-sell-go/model"
	"gorm.io/gorm"
)

type ListingRepository struct {
	db *gorm.DB
}

func NewListingRepository(db *gorm.DB) *ListingRepository {
	return &ListingRepository{db}
}

func (r *ListingRepository) CreateListing(listing model.CreateListingRequest) (*model.Listing, error) {
	newListing := model.NewListing(listing)
	result := r.db.Create(&newListing)

	if result.Error != nil {
		return nil, result.Error
	}

	return newListing, nil
}

func (r *ListingRepository) GetListings() ([]model.Listing, error) {
	var listings []model.Listing

	result := r.db.Find(&listings)

	if result.Error != nil {
		return nil, result.Error
	}

	return listings, nil
}

func (r *ListingRepository) GetListingById(id model.Identifier) (*model.Listing, error) {
	var listing model.Listing

	result := r.db.First(&listing, model.Model{ID: id})

	if result.Error != nil {
		return nil, result.Error
	}

	return &listing, nil
}

func (r *ListingRepository) UpdateListing(updates model.Listing) (*model.Listing, error) {
	existingResult := r.db.First(&model.Listing{}, model.Model{ID: updates.ID})

	if existingResult.Error != nil {
		return nil, existingResult.Error
	}

	result := r.db.Save(&updates)

	if result.Error != nil {
		return nil, result.Error
	}

	return &updates, nil
}
