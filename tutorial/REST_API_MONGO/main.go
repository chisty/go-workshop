package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

var client *mongo.Client

func dbsetup() (*mongo.Collection, context.Context) {
	collection := client.Database("chisty").Collection("person")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	return collection, ctx
}

func CreatePersonEndPoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var person Person
	error := json.NewDecoder(request.Body).Decode(&person)
	if error != nil {
		fmt.Println("Error in decode:= " + error.Error())
	}

	// collection := client.Database("chisty").Collection("person")
	// ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	collection, ctx := dbsetup()

	result, error := collection.InsertOne(ctx, person)
	if error != nil {
		fmt.Println("Error in insert: ", error.Error())
	}

	json.NewEncoder(response).Encode(result)
}

//GetPeopleEndPoint request uri localhost:12345/person
func GetPeopleEndPoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var people []Person

	// collection := client.Database("chisty").Collection("person")
	// ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	collection, ctx := dbsetup()
	cursor, error := collection.Find(ctx, bson.M{})
	if error != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + error.Error() + `"}`))
		return
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var person Person
		cursor.Decode(&person)
		people = append(people, person)
	}

	if error := cursor.Err(); error != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + error.Error() + `"}`))
		return
	}

	json.NewEncoder(response).Encode(people)
}

//GetPersonEndPoint request uri localhost:12345/person/5d85e47e1ef510acb93fdb99
func GetPersonEndPoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var person Person
	// collection := client.Database("chisty").Collection("person")
	// ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	collection, ctx := dbsetup()

	error := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
	if error != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + error.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(person)
}

func main() {
	fmt.Println("Starting the Server ...")

	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	router := mux.NewRouter()
	router.HandleFunc("/person", CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/person", GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/person/{id}", GetPersonEndPoint).Methods("GET")
	http.ListenAndServe(":12345", router)
}
