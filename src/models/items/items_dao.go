package items

import (
	"fmt"
	"github.com/Teslenk0/jr-management_inventory/src/database/mysql"
	"github.com/Teslenk0/utils-go/logger"
	"github.com/Teslenk0/utils-go/rest_errors"
	"strings"
)

const (
	//Basic CRUD operations
	queryInsertItem    = "INSERT INTO items (code, title, description, picture, price, internal_price, available_quantity, sold_quantity, provider, category, date_created) VALUES (?,?,?,?,?,?,?,?,?,?,?);"
	queryGetItemByCode = "SELECT code, title, COALESCE(description, '') as description, COALESCE(picture, '') as picture, price, internal_price, available_quantity, sold_quantity, provider, category, date_created FROM items WHERE code=?;"
	queryUpdateItem    = "UPDATE items SET title=?, description=?, picture=?, price=?, internal_price=?, available_quantity=?, sold_quantity=?, category=? WHERE code=?;"
	queryDeleteItem    = "DELETE FROM items WHERE code=?;"

	//Complex operations
	querySearchItem = `SELECT code, title, COALESCE(description, '') as description, COALESCE(picture, '') as picture, price, internal_price, available_quantity, sold_quantity, provider, category, date_created FROM items WHERE title LIKE '%'||?||'%' ;`
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
	if getErr := result.Scan(&item.Code, &item.Title, &item.Description, &item.Picture, &item.Price, &item.InternalPrice, &item.AvailableQuantity, &item.SoldQuantity, &item.Provider, &item.Category, &item.DateCreated); getErr != nil {
		if strings.Contains(getErr.Error(), "no rows in result set") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("item %s not found", item.Code))
		}

		logger.Error("error when trying to get item by id", getErr)
		return rest_errors.NewInternalServerError("error when trying to get item by id", getErr)

	}

	return nil
}

//Save item
func (item *Item) Save() *rest_errors.RestError {

	//Prepares the statement
	stmt, err := mysql.Client.Prepare(queryInsertItem)
	//Ask if there was errors when attempting for preparing the stmt
	if err != nil {
		logger.Error("error when trying to prepare the save item statement", err)
		return rest_errors.NewInternalServerError("error when trying to prepare insert item statement", err)
	}
	//Close the connection when the functions returns
	defer stmt.Close()

	//Exec the statement
	_, saveErr := stmt.Exec(item.Code, item.Title, item.Description, item.Picture, item.Price, item.InternalPrice, item.AvailableQuantity, item.SoldQuantity, item.Provider, item.Category, item.DateCreated)
	if saveErr != nil {
		logger.Error("error when saving the item", saveErr)
		return rest_errors.NewInternalServerError("error when trying to save the item", saveErr)
	}

	return nil
}

//Delete - deletes a given item
func (item *Item) Delete() *rest_errors.RestError {

	stmt, err := mysql.Client.Prepare(queryDeleteItem)
	if err != nil {
		logger.Error("error when trying to prepare the delete item statement", err)
		return rest_errors.NewInternalServerError("error when trying to prepare the delete item statement", err)
	}

	defer stmt.Close()

	_, delErr := stmt.Exec(item.Code)
	if delErr != nil {
		logger.Error("error when trying to delete the item", delErr)
		return rest_errors.NewInternalServerError("error when trying to delete the item", delErr)
	}
	return nil
}

//Updates an item
func (item *Item) Update() *rest_errors.RestError {
	stmt, err := mysql.Client.Prepare(queryUpdateItem)
	if err != nil {
		logger.Error("error when trying to prepare the update item statement", err)
		return rest_errors.NewInternalServerError("error when trying to prepare update item statement", err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(item.Title, item.Description, item.Picture, item.Price, item.InternalPrice, item.AvailableQuantity, item.SoldQuantity, item.Category, item.Code)
	if err != nil {
		logger.Error("error when trying to update the item", err)
		return rest_errors.NewInternalServerError("error when trying to update the item", err)
	}
	return nil

}

func (item *Item) SearchItem(wanted string) ([]Item, *rest_errors.RestError) {

	stmt, err := mysql.Client.Prepare(querySearchItem)
	if err != nil {
		logger.Error("error when trying to prepare the search item statement", err)
		return nil, rest_errors.NewInternalServerError("error when trying to prepare search item statement", err)
	}

	defer stmt.Close()
	rows, err := stmt.Query(wanted)
	if err != nil {
		logger.Error("error when trying to search the users", err)
		return nil, rest_errors.NewInternalServerError("error when trying to find the item", err)
	}
	defer rows.Close()

	results := make([]Item, 0)
	for rows.Next() {
		var obj Item
		if err := rows.Scan(&obj.Code, &obj.Title, &obj.Description, &obj.Picture, &obj.Price, &obj.InternalPrice, &obj.AvailableQuantity, &obj.SoldQuantity, &obj.Provider, &obj.Category, &obj.DateCreated); err != nil {
			logger.Error("error when trying fill the struct with database data in search item method", err)
			return nil, rest_errors.NewInternalServerError("database error", err)
		}
		results = append(results, obj)
	}

	if len(results) == 0 {
		return nil, rest_errors.NewNotFoundError("no item matching given parameter")
	}

	return results, nil

}
