package main

import "net/http"

type Post struct {
	Id              int    `json:"Id" bson:"Id"`
	Caption         string `json:"Caption" bson:"Caption"`
	ImageURL        string `json:"ImageURL" bson:"ImageURL"`
	PostedTimestamp string `json:"PostedTimestamp" bson:"PostedTimestamp"`
}

func PostsHandler(responseWriter http.ResponseWriter, request *http.Request) {

}
