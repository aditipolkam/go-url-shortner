package handlers

import (
	"context"
	"encoding/json"
	"go-server/internal/db"
	"go-server/internal/models"
	pkg "go-server/pkg/helpers"
	"log"
	"net/http"
	"time"
)

const shortCodeLength int = 6

func PostRecordHandler(w http.ResponseWriter, r *http.Request) {
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
	shortCode := pkg.GenerateShortCode(shortCodeLength)

	newUrl := models.URLMapping{
		ID:          shortCode,
		OriginalUrl: body.URL,
		CreatedAt:   time.Now(),
	}

	collection := db.Client.Database("testdb").Collection("urlCollection")

	_, err = collection.InsertOne(context.Background(), newUrl)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Failed to shorten URL", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"short_url": "http://localhost:8080/" + shortCode})
}
