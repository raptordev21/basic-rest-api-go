package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost/netflix"
const dbname = "netflix"
const colName = "watchlist"

var Collection *mongo.Collection

// init only run once
func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB Connected Successfully")

	Collection = (*mongo.Collection)(client.Database(dbname).Collection(colName))

	fmt.Println("Collection reference is ready")
}
