package main // import "github.com/eriol/wp24-athletes"

import (
	"github.com/eriol/wp24-athletes/api"
	"github.com/eriol/wp24-athletes/database"
)

func main() {
	var store database.Store
	store.Open("test.sqlite")

	api.Serve()
}
