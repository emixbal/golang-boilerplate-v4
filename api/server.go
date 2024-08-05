package api

import (
	"golang-websocket/api/routes"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	app := SetupRouter()
	app.Listen(":3000")
}

func SetupRouter() *fiber.App {
	app := fiber.New()
	v1 := app.Group("/v1")
	{
		routes.RouteCustomer(v1)
	}
	return app
}
