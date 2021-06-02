package main

import (
	"github.com/MuZaZaVr/notesService/internal/config"
	"github.com/MuZaZaVr/notesService/internal/repository"
	"github.com/MuZaZaVr/notesService/pkg/database/pg"
	"log"
)

const configPath = "config/main"

func main() {

	/* Config layer */
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatalf("Error while loading config: %s", err)
	}

	/* DB layer */
	db, err := pg.NewPgConnection(cfg.Pg)
	if err != nil {
		log.Fatalf("Error init db: %s", err)
	}

	/* Repository layer */
	_ = repository.NewRepositories(db)

}