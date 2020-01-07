package items

import (
	"github.com/Teslenk0/jr-management_inventory/src/services"
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
