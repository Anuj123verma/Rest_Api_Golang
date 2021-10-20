// @title Swagger Covid Information Storing In MongoDb API
// @version 1.0
// @description This is a server that can give store the covid related information in mongodb database.

// @contact.name Anuj Verma
// @contact.email anujssooni360@gmail.con

// @host localhost:8001
// @BasePath /
// @Schemes http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
package main

import (
	"github.com/Anuj123Verma/Rest_Api_Golang/controller"
	_ "github.com/Anuj123Verma/Rest_Api_Golang/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const (
	port = "8001"
)

func main() {
	// Echo Instance
	e := echo.New()
	// get request for consuming the api
	e.GET("/store", controller.Storedata)
	// get request for swagger
	e.GET("/swagger/*any", echoSwagger.WrapHandler)
	// Start Server
	e.Logger.Fatal(e.Start(":" + port))
}
