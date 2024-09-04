package main

import (
	"log"

	"github.com/nemesidaa/thumbsYT/internal/config"
	"github.com/nemesidaa/thumbsYT/internal/service"
)

func main() {
	service.SafeExecution()
	cfg := config.NewConfig()
	log.Println("cfg: ", cfg)
	if err := cfg.ParseFlags("server.json"); err != nil {
		log.Fatal(err)
	}
	log.Println("Flags Parsed!")

	service.ListenAndServe(cfg)
	log.Println("Server stopped!")
}
