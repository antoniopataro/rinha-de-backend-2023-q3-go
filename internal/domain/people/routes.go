package people

import "github.com/gofiber/fiber/v2"

func MakeRoutes(app *fiber.App, controller *Controller) {
	app.Post("/pessoas", controller.Create)
	app.Get("/contagem-pessoas", controller.Count)
	app.Get("/pessoas/:id", controller.Find)
	app.Get("/pessoas", controller.Search)
}
