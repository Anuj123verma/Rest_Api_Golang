package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Anuj123Verma/Rest_Api_Golang/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Connection = "mongodb://0.0.0.0:27017"
)

func Createdb(states entity.CovidCases, dbname string, collectionname string) {
	// creating the mongo client and opening the connection
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
	collection := db.Collection(collectionname)
	for i := 0; i < len(states.Data.Statewise); i++ {
		// adding all the state in the mongodb
		res, err := collection.InsertOne(context.Background(), states.Data.Statewise[i])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.InsertedID.(primitive.ObjectID).Timestamp())
	}
}
