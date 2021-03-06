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
	Provider          string  `json:"provider"`
	Category          string  `json:"category"`
	DateCreated       string  `json:"date_created"`
}

type Items []Item

func (item *Item) Validate() *rest_errors.RestError {
	if item.Code == "" {
		return rest_errors.NewBadRequestError("code must not be blank")
	}

	if item.Title == "" {
		return rest_errors.NewBadRequestError("title must not be blank")
	}

	if item.Price <= 0 {
		return rest_errors.NewBadRequestError("price must be greater than zero")
	}

	if item.InternalPrice <= 0 {
		return rest_errors.NewBadRequestError("internal price must be grater than zero")
	}

	if item.AvailableQuantity < 0 {
		return rest_errors.NewBadRequestError("available quantity must be zero or greater")
	}

	if item.Category == "" {
		return rest_errors.NewBadRequestError("category must not be blank")
	}

	if item.Provider == "" {
		return rest_errors.NewBadRequestError("provider must not be blank")
	}

	return nil
}
