package routes

import (
	"golang-websocket/api/handler/customer"

	"github.com/gofiber/fiber/v2"
)

func RouteCustomer(route fiber.Router) {

	handlerCustomer := customer.NewCustomerHandler()
	r := route.Group("/customer")
	{
		// r.Use(middleware.MiddlewareAuthentication)
		{
			r.Get("/list", func(c *fiber.Ctx) error {
				return handlerCustomer.List(c)
			})
			r.Get("/detail/:id", func(c *fiber.Ctx) error {
				return handlerCustomer.Detail(c)
			})
			r.Post("/insert", func(c *fiber.Ctx) error {
				return handlerCustomer.Insert(c)
			})
			r.Put("/update/:id", func(c *fiber.Ctx) error {
				return handlerCustomer.Update(c)
			})
			r.Delete("/delete/:id", func(c *fiber.Ctx) error {
				return handlerCustomer.Delete(c)
			})
		}
	}
}
