package controller

import (
	"context"
	"fmt"
	"log"
	"runtime/trace" 

	"github.com/mastermuskan22/mongodbapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://mastermuskan:Master@22@cluster0.v5pmjuh.mongodb.net/"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB Database connection is successful!")
	
	collection = client.Database(dbName).Collection(colName)

	//collection instance

	fmt.Println("Collection instance is ready!")
}

func insertOneMovie(movie model.Netflix){
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Inserted one movie with ID ", inserted.InsertedID)
}

func updateOneMovie(movieId string){
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil{
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Modified count: ", result.ModifiedCount)
}

func deleteOneMovie(movieId string){
	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil{
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Movie got deleted with delete count ", deleteCount)
}

func deleteAllMovies() int64{
	delelteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("The count of movies deleted ", delelteResult.DeletedCount)
	return delelteResult.DeletedCount
}

func getAllMovies() []primitive.M{
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil{
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()){
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil{
			log.Fatal(err)
		}
		movies=append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}