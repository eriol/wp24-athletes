package api // import "github.com/eriol/wp24-athletes/api"

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/eriol/wp24-athletes/database"
)

type ApiInfo struct {
	Description string `json:"description"`
	Version     string `json:"version"`
}

// Return API description.
// This endpoint is the root of the API.
func info(w http.ResponseWriter, r *http.Request) {
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

func getAthlete(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimSpace(r.PathValue("slug"))
	if slug == "" {
		http.NotFound(w, r)
		return
	}

	athlete, err := database.GetAthlete(slug)

	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(athlete); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
