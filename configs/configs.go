package configs

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConfigDB() *mongo.Client {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	url := "mongodb://" + os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT")

	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection by pinging the server
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("MongoDB server not reachable:", err)
	}

	fmt.Println("\nConnected to MongoDB! Success \n")

	return client
}

func GetCollention(client *mongo.Client, collectionName string) *mongo.Collection {
	fmt.Println("DB_NAME: ", os.Getenv("DATABASE_NAME"))
	collection := client.Database(os.Getenv("DATABASE_NAME")).Collection((collectionName))
	return collection
}
