package listing

import "github.com/ryanhopperlowe/buy-and-sell-go/model"

type Listing struct {
	model.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Views       uint    `json:"views"`
	UserId      string  `json:"userId"`
}

type CreateRequest struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	UserId      string  `json:"userId"`
}

func NewListing(data CreateRequest) *Listing {
	return &Listing{
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		UserId:      data.UserId,
		Views:       0,
	}
}
