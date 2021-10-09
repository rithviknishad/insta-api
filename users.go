package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/mail"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TODO: check if there's late initialization instead of using function to get collection
func collection() *mongo.Collection {
	return DB.Collection("users")
}

type UserPost struct {
	Name     string `bson:"name,omitempty"`
	Email    string `bson:"email,omitempty"`
	Password string `bson:"passsword,omitempty"`
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		GetUser(w, r)

	case "POST":
		CreateUser(w, r)

	default:
		fmt.Fprintf(w, "Whoops :/ We can do only GET and POST.")
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	user := bson.D{
		{"name", name},
		{"email", email},
		{"password", password},
	}

	if res, err := collection().InsertOne(MongoContext, user); err != nil {
		fmt.Fprintf(w, "Failed to create user.\n%s", err)
	} else {
		// Respond w/ created user's id.
		// TODO: convert to JSON instead of plain
		fmt.Fprint(w, res.InsertedID.(primitive.ObjectID).Hex())
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	name := keyVal["name"]

	fmt.Fprintf(w, "%s", name)
}
