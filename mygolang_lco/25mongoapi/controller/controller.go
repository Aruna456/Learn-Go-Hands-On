package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Aruna456/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://arunasubramanian456:Aru.n%40h456@godb.kdddom8.mongodb.net/?retryWrites=true&w=majority&appName=godb"
const dbName = "netflix"
const colName = "watchlist"

//MOST IMPORTANT

var collection *mongo.Collection

// Connect with mongoDB

func init() { //runs only one time

	//client options
	clientOption := options.Client().ApplyURI(connectionString)

	//Connect to db

	//context - when we are working outside the machine (database) it will be working with diff webserver,we need to provide a context for how long i could make a req,what happens if the request goes off.. etc.,

	//Type of context - Background() -> if anythings happens in bg returns a non-nil empty context
	//TODO() -> when we have no idea about what context to use

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connected Successfully!!")

	collection = client.Database(dbName).Collection(colName)

	//Collection reference is Ready
	fmt.Println("Collection Reference is Ready")
}

//MONGODB helpers - file

// insertOneRecord

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data Inserted in DB with the ID: ", inserted.InsertedID)
}

// updateOneRecord

func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	//M -> if want to have shorter and clearer result without worrying about caseInsenstivity
	//D -> Ordered element
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified Count: ", result.ModifiedCount)
}

//delete specific record or delete all  - works based on filter

func deleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	deletecnt, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Operation Succeeded!", deletecnt)
}

func deleteAllMovie() int64 {
	// filter :=
	deletecnt, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Delete Operation Succeeded: ", deletecnt)
	return deletecnt.DeletedCount
}

//Get all movies from the MongoDb

func getAllMovies() []primitive.M {

	//Cursor -> cusor.Next(context) => loop through values
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())

	return movies

}

//Actual Controllers:

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Movies")
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating Movie")
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Marking as Watched")
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting Movie")
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting All Movie")
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	// params := mux.Vars(r)

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)

}
