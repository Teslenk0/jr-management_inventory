package services

import (
	"github.com/Teslenk0/jr-management_inventory/src/models/items"
	"github.com/Teslenk0/utils-go/rest_errors"
)

//Interface with methods
type itemsServiceInterface interface {
	GetItem(itemCode string) (*items.Item, *rest_errors.RestError)
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
