package api // import "github.com/eriol/wp24-athletes/api"

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/eriol/wp24-athletes/database"
)

type ApiInfo struct {
	Description string `json:"description"`
	Version     string `json:"version"`
}

// Return API description.
// This endpoint is the root of the API.
func info(w http.ResponseWriter, r *http.Request) {
	// The "/" pattern matches everything, so check if we are at the
	// root and return a 404 otherwise.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	api := ApiInfo{
		Description: "A simple open REST API for athletes!",
		Version:     "0.1",
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(api); err != nil {
		panic(err)
	}
}

func getAthletes(w http.ResponseWriter, r *http.Request) {
	athletes, err := database.GetAthletes()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(athletes); err != nil {
		panic(err)
	}
}
