package main

import (
	"log"
	"net/http"

	"github.com/bdkamenov/gameserver/internal/routes"
)

func main() {

	// Handle routes
	routes.Handlers()

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
