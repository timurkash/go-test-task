package main

import (
	"github.com/timurkash/go-test-task/internal/app"
	"log"
)

func main() {
	cfg := app.ParseFlags()
	application := app.New(cfg)

	log.Printf("Starting go-test-task on port %d", cfg.Port)
	if err := application.Run(); err != nil {
		log.Fatal(err)
	}

}
