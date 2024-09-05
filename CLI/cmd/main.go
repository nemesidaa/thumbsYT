package main

import (
	"log"

	"github.com/nemesidaa/thumbsYT/CLI/internal/config"
	"github.com/nemesidaa/thumbsYT/CLI/internal/logic"
)

func main() {
	defer logic.SafeExecution()

	cfg := config.NewConfig()
	if err := cfg.ParseFlags("client.json"); err != nil {
		log.Fatal(err)
	}
	logic.Start(cfg)

	log.Println("Service stopped!")
}
