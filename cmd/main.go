package main

import (
	"log"

	"github.com/codepnw/live-streaming/internal/server"
)

func main() {
	err := server.Run()
	if err != nil {
		log.Fatalf("start server error: %v", err)
	}
}
