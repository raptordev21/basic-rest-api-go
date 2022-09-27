package helpers

import (
	"context"
	"fmt"
	"log"

	"github.com/raptordev21/basic-rest-api-go/config"
	"github.com/raptordev21/basic-rest-api-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertOneMovie(movie models.Netflix) {
	inserted, err := config.Collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one movie with id :", inserted.InsertedID)
}

func UpdateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	// bson.D for ordered values/representation
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := config.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified Count: ", result.ModifiedCount)
}

func DeleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	deleteCount, err := config.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Delete Count: ", deleteCount)
}

func DeleteAllMovies() int64 {
	deleteResult, err := config.Collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movies Deleted Count: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

func GetAllMovies() []primitive.M {
	curser, err := config.Collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for curser.Next(context.Background()) {
		var movie bson.M
		err := curser.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer curser.Close(context.Background())
	return movies
}
