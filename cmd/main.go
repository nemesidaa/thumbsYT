package main

import (
	"log"

	"github.com/nemesidaa/thumbsYT/internal/config"
	"github.com/nemesidaa/thumbsYT/internal/service"
)

func main() {

	cfg := config.NewConfig()
	if err := cfg.ParseFlags("config.json"); err != nil {
		log.Fatal(err)
	}
	service.ListenAndServe(cfg)
}
