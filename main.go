package main

import (
	"github.com/Anuj123Verma/Rest_Api_Golang/api"
	"github.com/labstack/echo/v4"
)

func main() {
	// Echo Instance
	e := echo.New()

	// Route => handler
	e.GET("/store", api.Storedata)

	// Start Server
	e.Logger.Fatal(e.Start(":8001"))
}
