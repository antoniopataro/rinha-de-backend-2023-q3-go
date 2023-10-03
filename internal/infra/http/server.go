package http

import (
	"github.com/antoniopataro/rinha-go/internal/domain/people"
	"github.com/antoniopataro/rinha-go/internal/infra/cache"
	"github.com/antoniopataro/rinha-go/internal/infra/database"
	"github.com/gofiber/fiber/v2"
)

func MakeServer(cache *cache.Cache, database *database.Database) (app *fiber.App) {
	app = fiber.New()

	app.Use("*", func(ctx *fiber.Ctx) error {
		ctx.Set("Content-Type", "application/json")

		return ctx.Next()
	})

	peopleReposirory := people.MakeReposirory(cache, database)
	peopleController := people.MakeController(peopleReposirory)
	people.MakeRoutes(app, peopleController)

	return app
}
