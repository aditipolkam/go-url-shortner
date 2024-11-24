package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type URLMapping struct {
	ID          string    `json:"id"`
	OriginalUrl string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
}

const shortCodeLength int = 6

var db *mongo.Collection

func main() {
	connectToDb()
	http.HandleFunc("/", shortenUrlHandler)
	fmt.Println("Server is running at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func shortenUrlHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetUrl(w, r)
	case "POST":
		handlePostUrl(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetUrl(w http.ResponseWriter, r *http.Request) {
	// shortCode := r.URL.Path[1:]
	// var result URLMapping

}

func handlePostUrl(w http.ResponseWriter, r *http.Request) {
	//define a valid body param structure
	var body struct {
		URL string `json:"url"`
	}

	//validate json body
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	//generate short code for url
	shortCode := generateShortCode(shortCodeLength)

	newUrl := URLMapping{
		ID:          shortCode,
		OriginalUrl: body.URL,
		CreatedAt:   time.Now(),
	}

	fmt.Print("newURl:", newUrl)

	collection := Client.Database("testdb").Collection("urlCollection")

	_, err = collection.InsertOne(context.Background(), newUrl)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to shorten URL", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"short_url": "http://localhost:8080/" + shortCode})
}
