package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	MongodbUri := "mongodb+srv://sourav:Iamdeveloper123$@cluster0.yofmoqr.mongodb.net/resturant?retryWrites=true&w=majority"
	fmt.Print(MongodbUri)

	client, err := mongo.NewClient(options.Client().ApplyURI(MongodbUri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connected successfully...")

	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("resturant").Collection(collectionName)
	return collection
}
