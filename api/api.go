package api // import "github.com/eriol/wp24-athletes/api"

import (
	"encoding/json"
	"log"
	"net/http"
)

func Serve() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", info)
	router.HandleFunc("GET /athletes", getAthletes)
	router.HandleFunc("GET /athletes/{slug}", getAthlete)
	router.HandleFunc("GET /search", search)
	router.HandleFunc("GET /images/{slug}", images)
	router.HandleFunc("GET /random", random)

	log.Fatal(http.ListenAndServe(":3000", router))
}

func toJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if status != http.StatusOK {
		w.WriteHeader(status)
	}

	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
