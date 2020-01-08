package services

import (
	"github.com/Teslenk0/jr-management_inventory/src/models/items"
	"github.com/Teslenk0/utils-go/date"
	"github.com/Teslenk0/utils-go/rest_errors"
	"strings"
)

//Interface with methods
type itemsServiceInterface interface {
	GetItem(itemCode string) (*items.Item, *rest_errors.RestError)
	CreateItem(item items.Item) (*items.Item, *rest_errors.RestError)
	DeleteItem(itemCode string) *rest_errors.RestError
	UpdateItem(isPartial bool, item items.Item) (*items.Item, *rest_errors.RestError)
}

//Struct
type itemsService struct {
}

//Implementing the interface
var (
	ItemsServices itemsServiceInterface = &itemsService{}
)

func (s *itemsService) GetItem(itemCode string) (*items.Item, *rest_errors.RestError) {
	if itemCode == "" {
		return nil, rest_errors.NewBadRequestError("item code must not be blank")
	}

	var result = &items.Item{Code: itemCode}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *itemsService) CreateItem(item items.Item) (*items.Item, *rest_errors.RestError) {

	if err := item.Validate(); err != nil {
		return nil, err
	}

	item.DateCreated = date.GetNowDBString()
	item.SoldQuantity = 0
	item.Provider = strings.ToUpper(strings.TrimSpace(item.Provider))
	item.Category = strings.ToUpper(strings.TrimSpace(item.Category))

	if err := item.Save(); err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *itemsService) DeleteItem(itemCode string) *rest_errors.RestError {
	dao := &items.Item{Code: itemCode}
	return dao.Delete()
}

//UpdateItem - this function updates the user with the data given
func (s *itemsService) UpdateItem(isPartial bool, item items.Item) (*items.Item, *rest_errors.RestError) {
	current := &items.Item{Code: item.Code}
	if err := current.Get(); err != nil {
		return nil, err
	}

	current.Category = strings.ToUpper(strings.TrimSpace(current.Category))

	if isPartial {
		if item.Title != "" {
			current.Title = item.Title
		}

		if item.Description != "" {
			current.Description = item.Description
		}

		if item.Picture != "" {
			current.Picture = item.Picture
		}

		if item.Price != 0 {
			current.Price = item.Price
		}

		if item.InternalPrice != 0 {
			current.InternalPrice = item.InternalPrice
		}

		if item.AvailableQuantity != 0 {
			current.AvailableQuantity = item.AvailableQuantity
		}

		if item.SoldQuantity != 0 {
			current.SoldQuantity = item.SoldQuantity
		}

		if item.Category != "" {
			current.Category = item.Category
		}

	} else {
		current.Title = item.Title
		current.Description = item.Description
		current.Picture = item.Picture
		current.Price = item.Price
		current.InternalPrice = item.InternalPrice
		current.AvailableQuantity = item.AvailableQuantity
		current.SoldQuantity = item.SoldQuantity
		current.Category = item.Category
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}
