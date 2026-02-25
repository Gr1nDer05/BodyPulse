package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB(connectionString string) error {
	var err error

	config, err := pgxpool.ParseConfig(connectionString)

	if err != nil {
		return err
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return err
	}

	DB = pool

	log.Println("Database succefully initialized")
	return nil
}
