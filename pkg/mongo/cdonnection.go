package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	"ulab3/config"
)

func Connection(config config.Config) (*mongo.Client, error) {
	// Build the MongoDB connection string
	var uri string

	if config.DB_USER != "" && config.DB_PASS != "" {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
			config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_PORT, config.DB_NAME)
	} else {
		uri = fmt.Sprintf("mongodb://%s:%s/%s",
			config.DB_HOST, config.DB_PORT, config.DB_NAME)
	}

	// Create a MongoDB client with the connection string
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB with a timeout
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB successfully!")
	return client, nil
}
