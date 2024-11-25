package routes

import (
	"go-server/internal/handlers"
	"net/http"
)

func shortenUrlHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handlers.GetRecordsHandler(w, r)
	case "POST":
		handlers.PostRecordHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func Setup() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", shortenUrlHandler)
	return router
}
