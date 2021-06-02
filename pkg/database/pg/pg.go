package pg

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/MuZaZaVr/notesService/internal/config"
	_ "github.com/lib/pq"
)

func NewPgConnection(pgConfig config.PgConfig) (*sql.DB, error) {
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s port=%d",
		pgConfig.Host, pgConfig.User, pgConfig.Name, pgConfig.SSLMode, pgConfig.Password, pgConfig.Port)
	log.Println("DbURI: " + dbURI)

	db, err := sql.Open(pgConfig.Dialect, dbURI)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	pgConfig.URI = dbURI

	return db, nil
}