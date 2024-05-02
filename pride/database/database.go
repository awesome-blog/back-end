package database

import (
	"context"
	"fmt"
	"log"
	"pride/configs"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {

	uri := fmt.Sprintf("mongodb://%s:%s",
		configs.GetEnv("MONGO_HOST"),
		configs.GetEnv("MONGO_PORT"))

	credential := options.Credential{
		Username: configs.GetEnv("MONGO_USER"),
		Password: configs.GetEnv("MONGO_PASS"),
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(uri).SetAuth(credential))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	return client
}

// Client instance
var DB *mongo.Client = ConnectDB()

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(configs.GetEnv("MONGO_DB")).Collection(collectionName)

	return collection
}
