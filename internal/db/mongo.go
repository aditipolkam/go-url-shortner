package db

import (
	"context"
	"log"
	"time"

	"go-server/internal/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect(cfg config.Config) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(cfg.DB_URL).SetServerAPIOptions(serverAPI)
	Client, err = mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Send a ping to confirm a successful connection
	if err := Client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	log.Println("Connected to MongoDB!")

}

func Disconnect() {
	if Client != nil {
		if err := Client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}
	log.Println("Disconnected from MongoDB")
}

func Collection(name string) *mongo.Collection {
	return Client.Database("testdb").Collection(name)
}
