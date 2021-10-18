package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Connection = "mongodb://0.0.0.0:27017"
)

type CovidCases struct {
	Success bool `json:"success"`
	Data    struct {
		Source        string    `json:"source"`
		LastRefreshed time.Time `json:"lastRefreshed"`
		Total         struct {
			Confirmed int `json:"confirmed"`
			Recovered int `json:"recovered"`
			Deaths    int `json:"deaths"`
			Active    int `json:"active"`
		} `json:"total"`
		Statewise []struct {
			State     string `json:"state" bson:"state"`
			Confirmed int    `json:"confirmed" bson:"confirmed cases"`
			Recovered int    `json:"recovered" bson:"recovered cases"`
			Deaths    int    `json:"deaths" bson:"deaths"`
			Active    int    `json:"active" bson:"active cases"`
		} `json:"statewise"`
	} `json:"data"`
	LastRefreshed    time.Time `json:"lastRefreshed"`
	LastOriginUpdate time.Time `json:"lastOriginUpdate"`
}

func Storedata(c echo.Context) error {
	dbname := c.QueryParam("db")
	colname := c.QueryParam("col")
	response, err := http.Get("https://api.rootnet.in/covid19-in/unofficial/covid19india.org/statewise")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Store the json value in the struct
	var states CovidCases
	json.Unmarshal(responseData, &states)

	// connecting with the mongodb
	client, err := mongo.NewClient(options.Client().ApplyURI(Connection))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// creating the database
	db := client.Database(dbname)
	// creating the collection
	collection := db.Collection(colname)
	for i := 0; i < len(states.Data.Statewise); i++ {
		// adding all the state in the mongodb
		res, err := collection.InsertOne(context.Background(), states.Data.Statewise[i])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.InsertedID.(primitive.ObjectID).Timestamp())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"sucess": "data added from the rest API",
	})
}
