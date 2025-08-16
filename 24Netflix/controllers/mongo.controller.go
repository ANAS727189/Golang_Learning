// mongo.controllers.go
package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ANAS727189/Netflix-Api/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Have to do :- go get go.mongodb.org/mongo-driver/mongo

var collection *mongo.Collection

func getMongoURI() string {
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		return "mongodb://localhost:27017"
	}
	return mongoURI
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file, using system environment variables")
	}
	connectionString := getMongoURI()
	dbName := getEnv("DB_NAME", "netflix")
	colName := getEnv("COLLECTION_NAME", "watchList")

	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	fmt.Println("MongoDB connection established successfully")
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready", collection)
}

func insertOneMovie(movie models.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal("Error inserting movie:", err)
	}
	fmt.Println("Inserted movie with ID:", inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal("Error updating movie:", err)
	}
	fmt.Printf("Matched %v documents and updated %v documents\n", res.MatchedCount, res.ModifiedCount)
	fmt.Println("Movie updated successfully", movieId)
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal("Error deleting movie:", err)
	}
	fmt.Printf("Deleted %v documents\n", res.DeletedCount)
	fmt.Println("Movie deleted successfully", movieId)
}

func deleteManyMovie() int64 {
	res, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal("Error deleting movies:", err)
	}
	fmt.Printf("Deleted %v documents\n", res.DeletedCount)
	fmt.Println("Watched movies deleted successfully")
	return res.DeletedCount
}

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal("Error fetching movies:", err)
	}
	defer cursor.Close(context.Background())

	// var movies []models.Netflix
	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal("Error decoding movie:", err)
		}
		movies = append(movies, movie)
	}
	err = cursor.Err()
	if err != nil {
		log.Fatal("Error iterating cursor:", err)
	}
	fmt.Println("Fetched movies:", movies)
	return movies
}

// ACTUAL CONTROLLER FUNCTIONS

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	movies := getAllMovies()
	if len(movies) == 0 {
		http.Error(w, "No movies found", http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		http.Error(w, "Error encoding movies to JSON", http.StatusInternalServerError)
		return
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "POST")

	var movie models.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "PUT")
	params := mux.Vars(r)
	updateOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "DELETE")
	deletedCount := deleteManyMovie()

	res := map[string]int64{"deletedCount": deletedCount}
	json.NewEncoder(w).Encode(res)
}
