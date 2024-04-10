package main // import "github.com/eriol/wp24-athletes"

import (
	"github.com/eriol/wp24-athletes/api"
	"github.com/eriol/wp24-athletes/database"
)

func main() {

	database.Open("test.sqlite")
	defer database.Close()

	api.Serve()
}
