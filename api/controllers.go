package api // import "github.com/eriol/wp24-athletes/api"

import (
	"database/sql"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/eriol/wp24-athletes/database"
	"golang.org/x/image/draw"
)

type ApiInfo struct {
	Description string `json:"description"`
	Version     string `json:"version"`
}

type ApiError struct {
	Error string `json:"error"`
}

func preflight(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Method", "GET, OPTIONS")
	(*w).Header().Set("Access-Control-Max-Age", "86400")
}

// Return API description.
// This endpoint is the root of the API.
func info(w http.ResponseWriter, r *http.Request) {
	preflight(&w)

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

// Return an array with all the athletes.
func getAthletes(w http.ResponseWriter, r *http.Request) {
	preflight(&w)

	athletes, err := database.GetAthletes()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	toJSON(w, http.StatusOK, athletes)
}

// Return the specified (in the path) athlete.
func getAthlete(w http.ResponseWriter, r *http.Request) {
	preflight(&w)

	slug := strings.TrimSpace(r.PathValue("slug"))
	if slug == "" {
		toJSON(w, http.StatusBadRequest, ApiError{Error: "No athlete slug provided"})
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

// Search for an athlete.
func search(w http.ResponseWriter, r *http.Request) {
	preflight(&w)
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
		return
	}

	athletes, err := database.Search(searchType, q)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	toJSON(w, http.StatusOK, athletes)
}

// Return the image of the specified athlete.
func images(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimSpace(r.PathValue("slug"))
	if slug == "" {
		toJSON(w, http.StatusBadRequest, ApiError{Error: "No athlete slug provided"})
		return
	}

	p := path.Join("./images", fmt.Sprintf("%s%s", slug, ".jpg"))
	f, err := os.Open(p)
	defer f.Close()
	if err != nil {
		toJSON(w, http.StatusNotFound, ApiError{Error: "No image found"})
		return
	}

	i, _, err := image.Decode(f)
	if err != nil {
		toJSON(w, http.StatusInternalServerError, ApiError{Error: "Can't decode the image"})
		return
	}

	var outputImage image.Image
	queryParams := r.URL.Query()
	size := queryParams.Get("size")

	scaleBy := 1
	if size == "M" {
		scaleBy = 2
	}
	if size == "S" {
		scaleBy = 4
	}

	if size == "M" || size == "S" {
		resized := image.NewRGBA(image.Rect(0, 0, i.Bounds().Max.X/scaleBy, i.Bounds().Max.Y/scaleBy))
		draw.BiLinear.Scale(resized, resized.Rect, i, i.Bounds(), draw.Over, nil)
		outputImage = resized
	} else {
		outputImage = i
	}

	w.Header().Set("Content-Type", "image/jpeg")
	jpeg.Encode(w, outputImage, &jpeg.Options{Quality: jpeg.DefaultQuality})
}

// Return a random athlete.
func random(w http.ResponseWriter, r *http.Request) {
	preflight(&w)

	athlete, err := database.GetRandomAthlete()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	toJSON(w, http.StatusOK, athlete)
}
