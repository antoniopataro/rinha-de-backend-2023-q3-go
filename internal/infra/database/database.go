package database

import (
	"context"
	"fmt"
	"log"

	"github.com/antoniopataro/rinha-go/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Client *pgxpool.Pool
}

func MakeDatabase(envs *config.Envs) *Database {
	USER := envs.DATABASE_USER
	PASSWORD := envs.DATABASE_PASSWORD
	HOST := envs.DATABASE_HOST
	PORT := envs.DATABASE_PORT
	NAME := envs.DATABASE_NAME

	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		USER,
		PASSWORD,
		HOST,
		PORT,
		NAME,
	)

	cfg, err := pgxpool.ParseConfig(url)

	if err != nil {
		log.Fatalln("could not parse config", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)

	if err != nil {
		log.Fatalln("could create connection pool", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalln("could not ping connection pool", err)
	}

	return &Database{
		Client: pool,
	}
}
