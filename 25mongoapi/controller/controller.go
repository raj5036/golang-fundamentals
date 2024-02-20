package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raj5036/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ConnectionString string = "mongodb+srv://raj5036:UEPzMoEAikdvu4ym@cluster0.cl7i0un.mongodb.net?retryWrites=true&w=majority"
const LocalConnectionString string = "mongodb://localhost:27017"
const DBName string = "netflix"
const CollectionName string = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// Connect with MongoDB
func init() {
	//client option
	clientOption := options.Client().ApplyURI(LocalConnectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	// Most Important: Assign collection to this Collection Instance
	collection = client.Database(DBName).Collection(CollectionName)

	//collection instance
	fmt.Println("Collection instance is ready")
}

// MongoDB Helper methods
func handleErrors(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func insertOneMovie(movie model.Netflix) {
	inserted, _ := collection.InsertOne(context.Background(), movie)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	updated, err := collection.UpdateOne(context.Background(), filter, update)
	handleErrors(err)

	fmt.Println("Updated movie successfully", updated)
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}

	deleted, err := collection.DeleteOne(context.Background(), filter)
	handleErrors(err)

	fmt.Println("Deleted movie successfully", deleted)
}

func deleteAllMovies() int64 {
	deleted, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	handleErrors(err)

	fmt.Println("Deleted all movies successfully", deleted)
	return deleted.DeletedCount
}

// ref: https://youtu.be/1P5b4McGJQQ?si=YoMB2EUSbTosoLx4
func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	return movies
}

// Controllers
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()

	json.NewEncoder(w).Encode(allMovies)
}

func CreateOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)

	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deletedCount := deleteAllMovies()
	json.NewEncoder(w).Encode(deletedCount)
}
