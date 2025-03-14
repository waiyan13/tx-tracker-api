package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/waiyan13/tx-tracker/config"
)

func Connect(cfg *config.Config) *sql.DB {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatalf("Error creating database handle: %s", err.Error())
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
	}

	return db
}
