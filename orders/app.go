package orders

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sms2sakthivel/order-manager/orders/database"
	"github.com/sms2sakthivel/order-manager/orders/model"
	"github.com/sms2sakthivel/order-manager/orders/repository"
	"github.com/sms2sakthivel/order-manager/orders/routes"
	"github.com/sms2sakthivel/order-manager/orders/service"
)

func NewApp() *fiber.App {
	// Step 1: Connect to the database
	database.Connect()

	// Step 2: Auto-migrate Order schema
	model.Automigrate(database.DB)

	// Step 3: Initialize repository, service, and app
	repo := &repository.GormOrderRepository{DB: database.DB}
	service := &service.OrderService{Repo: repo}
	service.InitCartService(database.DB)

	app := fiber.New()

	// Step 4: Enable Logger middleware with timing
	app.Use(logger.New(logger.Config{
		Format: "${time} - ${latency} - ${status} - ${method} ${path}\n",
	}))

	// Step 5: Register routes
	routes.RegisterRoutes(app, service)
	return app
}
