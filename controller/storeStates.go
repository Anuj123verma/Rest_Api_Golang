package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Anuj123Verma/Rest_Api_Golang/entity"
	"github.com/Anuj123Verma/Rest_Api_Golang/mongodb"
)

func Storedata(c echo.Context) error {
	dbname := c.QueryParam("db")
	collectionname := c.QueryParam("col")
	var states entity.CovidCases
	// Consuming the another rest api for getting the statewise covid casesS
	var nstates entity.CovidCases = ConsumeApi(states)
	// connecting with the mongodb
	mongodb.Createdb(nstates, dbname, collectionname)
	return c.JSON(http.StatusOK, map[string]string{
		"sucess": "data added from the rest API",
	})
}
