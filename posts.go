package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id              int    `json:"_id" bson:"_id"`
	Caption         string `json:"caption" bson:"caption"`
	ImageURL        string `json:"imageURL" bson:"imageURL"`
	PostedTimestamp string `json:"postedTimestamp" bson:"postedTimestamp"`
}

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		getPost(w, r)

	case "POST":
		createPost(w, r)

	default:
		fmt.Fprintf(w, "Whoops :/ We can do only GET and POST.")
	}
}

func createPost(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	// Get attributes from body.
	caption := keyVal["caption"]
	imageURL := keyVal["imageURL"]
	timestamp := keyVal["postedTimestamp"]

	// TODO: password = securedPassword

	doc := bson.M{
		"caption":         caption,
		"imageURL":        imageURL,
		"postedTimestamp": timestamp,
	}

	if res, err := DB.Collection("posts").InsertOne(MongoContext, doc); err != nil {
		fmt.Fprintf(w, "Failed to create post.\n%s", err)
	} else {
		// Respond w/ created user's id.
		// TODO: convert to JSON instead of plain
		fmt.Fprint(w, res.InsertedID.(primitive.ObjectID).Hex())
	}
}

func getPost(w http.ResponseWriter, r *http.Request) {

}
