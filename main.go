package main // import "github.com/eriol/wp24-athletes"

import (
	"log"

	"github.com/eriol/wp24-athletes/api"
	"github.com/eriol/wp24-athletes/database"
)

func main() {

	err := database.Open("test.sqlite")
	defer database.Close()
	if err != nil {
		log.Fatal(err)
	}

	api.Serve()
}
