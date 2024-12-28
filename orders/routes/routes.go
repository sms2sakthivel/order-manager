package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/sms2sakthivel/order-manager/orders/service"
)

func RegisterRoutes(app *fiber.App, service *service.OrderService) {
	app.Get("/", OrderServiceInfo)
	app.Get("/swagger/*", swagger.HandlerDefault)

	// order endpoints
	app.Get("/orders", func(c *fiber.Ctx) error { return GetAllOrders(c, service) })
	app.Get("/orders/:id", func(c *fiber.Ctx) error { return GetOrderByID(c, service) })
	app.Post("/orders", func(c *fiber.Ctx) error { return CreateOrder(c, service) })
	app.Put("/orders/:id", func(c *fiber.Ctx) error { return UpdateOrder(c, service) })
	app.Delete("/orders/:id", func(c *fiber.Ctx) error { return DeleteOrder(c, service) })

	// cart endpoints
	app.Get("/carts", func(c *fiber.Ctx) error { return GetAllCarts(c, &service.CartService) })
	app.Get("/carts/:id", func(c *fiber.Ctx) error { return GetCartByID(c, &service.CartService) })
	app.Post("/carts", func(c *fiber.Ctx) error { return CreateCart(c, &service.CartService) })
	app.Put("/carts/:id", func(c *fiber.Ctx) error { return UpdateCart(c, &service.CartService) })
	app.Delete("/carts/:id", func(c *fiber.Ctx) error { return DeleteCart(c, &service.CartService) })
}
