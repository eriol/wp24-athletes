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
	// The "/" pattern matches everything, so check if we are at the
	// root and return a 403 otherwise (we blame the client for endpoints that
	// don't exist!).
	//
	// It's not possible to specify a custom NotFound(), because in
	// https://golang.org/src/pkg/net/http/server.go NotFoundHandler()
	// returns a hardcoded function called NotFound(). So we need to do this to
	// use JSON instead.
	if r.URL.Path != "/" {
		toJSON(w, http.StatusForbidden, ApiError{Error: "Forbidden"})
		return
	}

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

func search(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	var searchType database.SearchType
	var q string
	name := queryParams.Get("name")
	sport := queryParams.Get("sport")
	famousFor := queryParams.Get("famous_for")

	if name != "" {
		searchType = database.SearchByName
		q = strings.TrimSpace(name)
	} else if sport != "" {
		searchType = database.SearchBySport
		q = strings.TrimSpace(sport)
	} else if famousFor != "" {
		searchType = database.SearchByFamousFor
		q = strings.TrimSpace(famousFor)
	} else {
		toJSON(w, http.StatusBadRequest, ApiError{Error: "You have to search by name, sport or famous_for."})
	}

	athletes, err := database.Search(searchType, q)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	toJSON(w, http.StatusOK, athletes)
}
