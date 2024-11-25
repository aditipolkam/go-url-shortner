package handlers

import (
	"context"
	"go-server/internal/db"
	"go-server/internal/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetRecordsHandler(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[1:]

	var result models.URLMapping
	collection := db.Collection("urlCollection")

	err := collection.FindOne(context.Background(), bson.M{"id": shortCode}).Decode(&result)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	// Redirect to the original URL
	http.Redirect(w, r, result.OriginalUrl, http.StatusFound)
}
