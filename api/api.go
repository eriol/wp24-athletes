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

	log.Fatal(http.ListenAndServe(":8080", router))
}

func toJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
