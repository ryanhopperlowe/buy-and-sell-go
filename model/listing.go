package model

type Listing struct {
	Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Views       uint    `json:"views"`
	UserId      string  `json:"userId"`
}

type CreateListingRequest struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	UserId      string  `json:"userId"`
}

func NewListing(data CreateListingRequest) *Listing {
	return &Listing{
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		UserId:      data.UserId,
		Views:       0,
	}
}
