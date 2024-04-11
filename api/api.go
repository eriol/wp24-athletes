package api // import "github.com/eriol/wp24-athletes/api"

import (
	"log"
	"net/http"
)

func Serve() {
	router := http.NewServeMux()

	router.HandleFunc("GET /{$}", info)
	router.HandleFunc("GET /athletes", getAthletes)

	log.Fatal(http.ListenAndServe(":8080", router))
}
