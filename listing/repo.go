package listing

import (
	"github.com/ryanhopperlowe/buy-and-sell-go/model"
	"gorm.io/gorm"
)

type Repository interface {
	CreateListing(listing CreateRequest) (*Listing, error)
	GetListings() ([]Listing, error)
	GetListingById(id model.Identifier) (*Listing, error)
	UpdateListing(updates Listing) (*Listing, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateListing(listing CreateRequest) (*Listing, error) {
	newListing := NewListing(listing)
	result := r.db.Create(&newListing)

	if result.Error != nil {
		return nil, result.Error
	}

	return newListing, nil
}

func (r *repository) GetListings() ([]Listing, error) {
	var listings []Listing

	result := r.db.Find(&listings)

	if result.Error != nil {
		return nil, result.Error
	}

	return listings, nil
}

func (r *repository) GetListingById(id model.Identifier) (*Listing, error) {
	var listing Listing

	result := r.db.First(&listing, model.Model{ID: id})

	if result.Error != nil {
		return nil, result.Error
	}

	return &listing, nil
}

func (r *repository) UpdateListing(updates Listing) (*Listing, error) {
	existingResult := r.db.First(&Listing{}, model.Model{ID: updates.ID})

	if existingResult.Error != nil {
		return nil, existingResult.Error
	}

	result := r.db.Save(&updates)

	if result.Error != nil {
		return nil, result.Error
	}

	return &updates, nil
}
