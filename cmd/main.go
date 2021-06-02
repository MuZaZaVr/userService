package main

import (
	"github.com/MuZaZaVr/notesService/internal/app"
)

const configPath = "config/main"

func main() {
	app.Run(configPath)
}
