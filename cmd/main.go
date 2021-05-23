package main

import (
	"fmt"
	"github.com/MuZaZaVr/notesService/internal/config"
	"log"
)

const configPath = "config/main"

func main() {

	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatalf("Error while loading config: %s", err)
	}
	fmt.Printf("Config: %v", cfg)	// remove after next step

}