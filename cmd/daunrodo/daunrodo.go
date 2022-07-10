package main

import (
	"flag"
	"log"

	"daunrodo/internal/app"
	"daunrodo/pkg/config"
)

func main() {

	configPath := flag.String("c", "../../config/config.yaml", "Config path")
	flag.Parse()

	cfg, err := config.New(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
