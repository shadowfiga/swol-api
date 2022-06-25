package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"swol-api/config"
)

func newDbConnection() *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Conf.DbUrl))

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to DB")
	return client.Database("swol")
}

var (
	Db = newDbConnection()
)
