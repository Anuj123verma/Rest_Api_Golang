package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Anuj123Verma/Rest_Api_Golang/entity"
	"github.com/Anuj123Verma/Rest_Api_Golang/mongodb"
)

// CreateResource godoc
// @Summary Store the covid related information about the state in mongodb
// @Description Database name and Collection name as input and create it in mongodb local.
// @Tags Information Storage
// @Param db query string true "Database Name"
// @Param collection query string true "Collection Name"
// @Consume  application/json
// @Produce json
// @Success 200  {object} entity.Success
// @Failure 500  {string} string "error"
// @Router /store [get]
func Storedata(c echo.Context) error {

	// storing the values of query params in variables
	dbname := c.QueryParam("db")
	collectionname := c.QueryParam("collection")

	// CovidCases struct variable declared
	var states entity.CovidCases
	// Consuming the another rest api for getting the statewise covid casesS
	var nstates entity.CovidCases = ConsumeApi(states)
	// connecting with the mongodb
	mongodb.Createdb(nstates, dbname, collectionname)

	// returning the response.
	return c.JSON(http.StatusOK, &entity.Success{
		Success: "data added from the rest API",
	})
}
