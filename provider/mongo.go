package provider

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(uri string, databaseName string) (*mongo.Client, *mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		fmt.Println("FAILED TO CREATE CLIENT")
		return nil, nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("FAILED TO CONNECT")
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("FAILED TO PING")
		return nil, nil, err
	}
	database := client.Database(databaseName)
	return client, database, nil
}
