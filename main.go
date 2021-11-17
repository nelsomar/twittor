package main

import (
	"log"

	"github.com/nelsomar/twittor/bd"
	"github.com/nelsomar/twittor/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("No connection to database")
		return
	}
	handlers.HandlersAPI()
}
