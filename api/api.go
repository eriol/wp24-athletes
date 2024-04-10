package api // import "github.com/eriol/wp24-athletes/api"

import (
	"log"
	"net/http"
)

func Serve() {
	http.HandleFunc("/", info)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
