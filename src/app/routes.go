package app

import (
	"github.com/Teslenk0/jr-management_inventory/src/controllers/items"
	"github.com/Teslenk0/jr-management_inventory/src/controllers/ping"
)

func routes(){

	//Endpoint to verify if API is listening
	router.GET("/ping", ping.Ping)
	//-------------------------------------------

	router.GET("/items/:code", items.Get)


}