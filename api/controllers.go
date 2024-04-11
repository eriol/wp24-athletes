package api // import "github.com/eriol/wp24-athletes/api"

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"github.com/eriol/wp24-athletes/database"
)

type ApiInfo struct {
	Description string `json:"description"`
	Version     string `json:"version"`
}

type ApiError struct {
	Error string `json:"error"`
}

// Return API description.
// This endpoint is the root of the API.
func info(w http.ResponseWriter, r *http.Request) {
	api := ApiInfo{
		Description: "A simple open REST API for athletes!",
		Version:     "0.1",
	}

	toJSON(w, http.StatusOK, api)
}

func getAthletes(w http.ResponseWriter, r *http.Request) {
	athletes, err := database.GetAthletes()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	toJSON(w, http.StatusOK, athletes)
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
			toJSON(w, http.StatusNotFound, ApiError{Error: "No athlete found"})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	toJSON(w, http.StatusOK, athlete)
}
