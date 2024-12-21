package main

import (
	"log"

	"github.com/stuckinforloop/dist/server"
)

func main() {
	// initialize dependencies
	srv := server.New()

	// register handlers
	srv.RegisterHandlers()

	// start maelstrom server
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
