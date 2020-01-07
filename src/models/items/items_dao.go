package items

import (
	"fmt"
	"github.com/Teslenk0/jr-management_inventory/src/database/mysql"
	"github.com/Teslenk0/utils-go/logger"
	"github.com/Teslenk0/utils-go/rest_errors"
	"strings"
)

const (
	//Basic CRUD queries
	queryInsertItem = "INSERT INTO items (code, title, description, picture, price, " +
		"internal_price, available_quantity, sold_quantity) VALUES (?,?,?,?,?,?,?,?);"
	queryGetItemByCode = "SELECT code, title, COALESCE(description, '') as description, COALESCE(picture, '') as picture, price, internal_price, available_quantity, sold_quantity, provider FROM items WHERE code=?;"
	queryUpdateItem    = "UPDATE items SET title=?, description=?, picture=?, price=?, internal_price=?, available_quantity=?, sold_quantity=? WHERE code=?;"
	queryDeleteItem    = "DELETE FROM items WHERE code=?;"
)

//Get item
func (item *Item) Get() *rest_errors.RestError {
	//Prepares the query
	stmt, err := mysql.Client.Prepare(queryGetItemByCode)

	if err != nil {
		logger.Error("error when trying to prepare the get user statement", err)
		return rest_errors.NewInternalServerError("error when trying to prepare the get item statement", err)
	}

	//Close the stametent when the function returns
	defer stmt.Close()

	//Make a select and looks for only one result
	result := stmt.QueryRow(item.Code)

	//Populates the user given with the data from DB
	if getErr := result.Scan(&item.Code, &item.Title, &item.Description, &item.Picture, &item.Price, &item.InternalPrice, &item.AvailableQuantity, &item.SoldQuantity, &item.Provider); getErr != nil {
		if strings.Contains(getErr.Error(), "no rows in result set") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("item %s not found", item.Code))
		}

		logger.Error("error when trying to get item by id", getErr)
		return rest_errors.NewInternalServerError("error when trying to get item by id", getErr)

	}

	return nil
}
