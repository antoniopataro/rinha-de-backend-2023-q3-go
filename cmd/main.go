package main

import (
	"fmt"

	"github.com/antoniopataro/rinha-de-backend-2023-q3-go/internal/config"
	"github.com/antoniopataro/rinha-de-backend-2023-q3-go/internal/infra/cache"
	"github.com/antoniopataro/rinha-de-backend-2023-q3-go/internal/infra/database"
	"github.com/antoniopataro/rinha-de-backend-2023-q3-go/internal/infra/http"
	"github.com/antoniopataro/rinha-de-backend-2023-q3-go/pkg/shutdown"
	"github.com/google/uuid"
)

func main() {
	defer shutdown.Gracefully()

	uuid.EnableRandPool()

	envs := config.MakeEnvs()

	cache := cache.MakeCache(envs)
	db := database.MakeDatabase(envs)

	app := http.MakeServer(cache, db)

	app.Listen(fmt.Sprintf(":%s", envs.PORT))

}
