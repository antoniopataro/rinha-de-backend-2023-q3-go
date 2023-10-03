package config

import (
	"github.com/antoniopataro/rinha-go/pkg/envs"
)

type Envs struct {
	DATABASE_HOST     string
	DATABASE_NAME     string
	DATABASE_PASSWORD string
	DATABASE_PORT     string
	DATABASE_USER     string
	PORT              string
	CACHE_ADDRESS     string
	CACHE_PORT        string
}

func MakeEnvs() *Envs {
	// if err := godotenv.Load(".env"); err != nil {
	// 	log.Fatalf("error loading .env file: %v", err)
	// }

	return &Envs{
		CACHE_ADDRESS:     envs.GetEnvOrDie("CACHE_ADDRESS"),
		CACHE_PORT:        envs.GetEnvOrDie("CACHE_PORT"),
		DATABASE_HOST:     envs.GetEnvOrDie("DATABASE_HOST"),
		DATABASE_NAME:     envs.GetEnvOrDie("DATABASE_NAME"),
		DATABASE_PASSWORD: envs.GetEnvOrDie("DATABASE_PASSWORD"),
		DATABASE_PORT:     envs.GetEnvOrDie("DATABASE_PORT"),
		DATABASE_USER:     envs.GetEnvOrDie("DATABASE_USER"),
		PORT:              envs.GetEnvOrDie("PORT"),
	}
}
