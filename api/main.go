package main

import (
	"github.com/aweglteo/fullstack_development/api/external/router"
)

func main() {
	// listening gin server
	router.Router.Run()
}
