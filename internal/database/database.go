package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DbName = "Main"
)

var MongoClient *mongo.Client

func ConnectToMongoDB(uri string) error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	MongoClient = client
	log.Println("Connected to MongoDB!")

	return nil
}

func GetCollection(databaseName, collectionName string) *mongo.Collection {
	if MongoClient == nil {
		log.Fatal("MongoDB client not initialized")
	}
	return MongoClient.Database(databaseName).Collection(collectionName)
}
