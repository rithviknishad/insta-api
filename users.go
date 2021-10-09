package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/mail"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"passsword"`
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		getUser(w, r)

	case "POST":
		createUser(w, r)

	default:
		fmt.Fprintf(w, "Whoops :/ We can do only GET and POST.")
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	// Get attributes from body.
	name := keyVal["name"]
	email := keyVal["email"]
	password := keyVal["password"]

	// Check if email is valid
	if _, err := mail.ParseAddress(email); err != nil {
		fmt.Fprintf(w, "Expected a valid email. Got %s", email)
		return
	}

	// Minimum password length
	if plen := len(password); plen < 6 {
		fmt.Fprintf(w, "Password must be atleast 6 digits. Got %d", plen)
		return
	}

	// TODO: password = securedPassword

	doc := bson.M{
		"name":     name,
		"email":    email,
		"password": password,
	}

	if res, err := DB.Collection("users").InsertOne(MongoContext, doc); err != nil {
		fmt.Fprintf(w, "Failed to create user.\n%s", err)
	} else {
		// Respond w/ created user's id.
		// TODO: convert to JSON instead of plain

		id := res.InsertedID.(primitive.ObjectID).Hex()

		fmt.Fprint(w, bson.M{"_id": id})
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	id := keyVal["_id"]

	print(id)

	result := DB.Collection("users").FindOne(MongoContext, bson.M{"_id": id})

	fmt.Fprint(w, result.Decode(&User{}))
}
