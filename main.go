package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

const HomeEndpoint = "/"
const UsersEndpoint = "/users"
const PostsEndpoint = "/posts"

var (
	DB           *mongo.Database
	MongoContext context.Context
)

func main() {

	client, MongoContext, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	defer close(client, MongoContext, cancel)

	ping(client, MongoContext)

	DB = client.Database("insta-api-test")

	handleRequests()
}

func handleRequests() {
	http.HandleFunc(HomeEndpoint, HomeHandler)
	http.HandleFunc(UsersEndpoint, UsersHandler)
	http.HandleFunc(PostsEndpoint, PostsHandler)

	fmt.Println("Starting instaapi server at 0.0.0.0:9090 ...")

	log.Fatal(http.ListenAndServe(":9090", nil))
}
