package app

import (
	"context"
	"github.com/MuZaZaVr/notesService/internal/config"
	"github.com/MuZaZaVr/notesService/internal/handler"
	"github.com/MuZaZaVr/notesService/internal/repository"
	"github.com/MuZaZaVr/notesService/internal/server"
	"github.com/MuZaZaVr/notesService/internal/service"
	"github.com/MuZaZaVr/notesService/pkg/database/pg"
	"log"
	"os"
	"os/signal"
	"time"
)

func Run(configPath string) {

	ctx := context.Background()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

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

	/* Repositories layer */
	repos := repository.NewRepositories(db)

	/* Services layer */
	services := service.NewServices(service.Depends{
		Repos: repos,
	})

	newHandler := handler.NewHandler(services)

	newServer := server.NewServer(cfg, newHandler)

	go startService(ctx, newServer)

	<-stop

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := newServer.Stop(ctx); err != nil {
		log.Fatalf("Failed to stop the newServer: %s", err.Error())
	}

	log.Printf("Sgutting down newServer...")
}

func startService(ctx context.Context, server *server.Server) {
	if err := server.Run(); err != nil {
		log.Fatal(ctx, "server shutdown: ", err.Error())
	}
}
