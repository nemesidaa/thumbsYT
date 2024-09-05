package main

import (
	"log"

	"github.com/nemesidaa/thumbsYT/internal/config"
	"github.com/nemesidaa/thumbsYT/internal/service"
)

func main() {
	defer service.SafeExecution()
	cfg := config.NewConfig()
	if err := cfg.ParseFlags("server.json"); err != nil {
		log.Fatal(err)
	}
	log.Println("Flags Parsed!")

	if err := service.ListenAndServe(cfg); err != nil {
		log.Fatal(err)
	}
	log.Println("Server stopped!")
}
