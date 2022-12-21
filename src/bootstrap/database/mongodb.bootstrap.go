package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnection struct {
	Conn *mongo.Database
}

var mongoSingleton *MongoConnection

func GetMongoConnection() *MongoConnection {
	if mongoSingleton == nil {
		os.Setenv("MONGODB_URI", "mongodb+srv://sembenavente:yLYny36GbMfCehL@clusterforlearning.6ihmn20.mongodb.net/?retryWrites=true&w=majority")
		uri := os.Getenv("MONGODB_URI")
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
			panic(err)
		}
		conn := client.Database("database01")
		mongoSingleton = &MongoConnection{
			Conn: conn,
		}
		log.Println("Mongo connection created")
	}
	return mongoSingleton
}
