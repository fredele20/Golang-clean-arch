package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const connectionString = "mongodb://localhost:27017"

const DBNAME = "playground"

const USERCOLLECTION = "users"
const POSTCOLLECTION = "posts"

var userCollection *mongo.Collection
var postCollection *mongo.Collection

func init () {
	// set the client options
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return
	}

	// check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connection to the database established")
	userCollection = client.Database(DBNAME).Collection(USERCOLLECTION)
	postCollection = client.Database(DBNAME).Collection(POSTCOLLECTION)

	fmt.Println("connection instance created...")
}