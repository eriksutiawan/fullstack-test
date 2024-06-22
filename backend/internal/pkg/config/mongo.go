package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongo() *mongo.Database {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	var err error
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(os.Getenv("COLLECTION_NAME"))
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error : ", err)
	}
}
