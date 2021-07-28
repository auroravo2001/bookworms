package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Ctx = context.TODO()
var ReviewCollection *mongo.Collection

func DB() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))

	client, err := mongo.Connect(Ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	//check for the connection
	err = client.Ping(Ctx, nil)

	if err != nil {
		return nil, err
	}

	db := client.Database("bookworms")

	ReviewCollection = db.Collection("review")

	return client.Database("bookworms"), nil
}
