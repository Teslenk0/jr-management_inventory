package items

import (
	"github.com/Teslenk0/utils-go/rest_errors"
)

type Item struct {
	Code              string  `json:"code"`
	Title             string  `json:"title"`
	Description       string  `json:"description"`
	Picture           string  `json:"picture"`
	Price             float32 `json:"price"`
	InternalPrice     float32 `json:"internal_price"`
	AvailableQuantity int     `json:"available_quantity"`
	SoldQuantity      int     `json:"sold_quantity"`
}

type Items []Item

func (i *Item) Validate() *rest_errors.RestError {
	if i.Code == "" {
		return rest_errors.NewBadRequestError("code must not be blank")
	}

	if i.Title == "" {
		return rest_errors.NewBadRequestError("title must not be blank")
	}

	if i.Price <= 0 {
		return rest_errors.NewBadRequestError("price must be greater than zero")
	}

	if i.InternalPrice <= 0 {
		return rest_errors.NewBadRequestError("internal price must be grater than zero")
	}

	if i.AvailableQuantity < 0 {
		return rest_errors.NewBadRequestError("available quantity must be zero or greater")
	}
	return nil
}
