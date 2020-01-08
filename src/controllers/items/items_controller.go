package items

import (
	"github.com/Teslenk0/jr-management_inventory/src/models/items"
	"github.com/Teslenk0/jr-management_inventory/src/services"
	"github.com/Teslenk0/utils-go/rest_errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Retrieves a given item from database
func Get(c *gin.Context) {
	//Get the id from the GET request
	itemCode := c.Param("code")

	result, getErr := services.ItemsServices.GetItem(itemCode)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

//Create - Function to create an item
func Create(c *gin.Context) {
	var item items.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.ItemsServices.CreateItem(item)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

//Delete - Deletes a given item from database
func Delete(c *gin.Context) {
	itemCode := c.Param("code")

	if err := services.ItemsServices.DeleteItem(itemCode); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

//Update - function to update data
func Update(c *gin.Context) {
	itemCode := c.Param("code")

	var item items.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	item.Code = itemCode

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.ItemsServices.UpdateItem(isPartial, item)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
