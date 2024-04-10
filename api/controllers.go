package api // import "github.com/eriol/wp24-athletes/api"

import (
	"encoding/json"
	"net/http"
)

type ApiInfo struct {
	Description string `json:"description"`
	Version     string `json:"version"`
}

// Return API description
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
