package main

import (
	"github.com/labstack/echo/v4"

	".\GoWeb\api"
)

func main() {
	// Echo Instance
	e := echo.New()

	// Route => handler
	e.GET("/store", api.storeData)

	// Start Server
	e.Logger.Fatal(e.Start(":8001"))
}
