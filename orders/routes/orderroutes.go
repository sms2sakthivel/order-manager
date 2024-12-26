package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/sms2sakthivel/order-manager/orders/model"
	"github.com/sms2sakthivel/order-manager/orders/service"
)

func RegisterRoutes(app *fiber.App, service *service.OrderService) {
	app.Get("/", OrderServiceInfo)
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/orders", func(c *fiber.Ctx) error { return GetAllOrders(c, service) })
	app.Get("/orders/:id", func(c *fiber.Ctx) error { return GetOrderByID(c, service) })
	app.Post("/orders", func(c *fiber.Ctx) error { return CreateOrder(c, service) })
	app.Put("/orders/:id", func(c *fiber.Ctx) error { return UpdateOrder(c, service) })
	app.Delete("/orders/:id", func(c *fiber.Ctx) error { return DeleteOrder(c, service) })
}

// OrderServiceInfo returns information about the Order Service
//
// @Summary      Order Service Info
// @Description  Returns basic information about the Order Service
// @Tags         General
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       / [get]
func OrderServiceInfo(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Order Service"})
}

// GetAllOrders retrieves all orders
//
// @Summary      Get All Orders
// @Description  Retrieve a list of all orders
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.OrderResponse
// @Failure      500  {object}  fiber.Error
// @Router       /orders [get]
func GetAllOrders(c *fiber.Ctx, service *service.OrderService) error {
	orders, err := service.GetOrders()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(orders)
}

// GetOrderByID retrieves a order by their ID
//
// @Summary      Get Order by ID
// @Description  Retrieve a order by their ID
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Order ID"
// @Success      200  {object}  model.OrderResponse
// @Failure      404  {object}  fiber.Error
// @Failure      500  {object}  fiber.Error
// @Router       /orders/{id} [get]
func GetOrderByID(c *fiber.Ctx, service *service.OrderService) error {
	id, _ := strconv.Atoi(c.Params("id"))
	order, err := service.GetOrder(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Order not found")
	}
	return c.JSON(order)
}

// CreateOrder adds a new order
//
// @Summary      Create a New Order
// @Description  Add a new order to the system
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        order  body      model.OrderCreateRequest  true  "Order details"
// @Success      201   {object}  model.OrderResponse
// @Failure      400   {object}  fiber.Error
// @Failure      500   {object}  fiber.Error
// @Router       /orders [post]
func CreateOrder(c *fiber.Ctx, service *service.OrderService) error {
	var order model.Order
	if err := c.BodyParser(&order); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}
	if err := service.CreateOrder(&order); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(order)
}

// UpdateOrder modifies details of an existing order
//
// @Summary      Update an Existing Order
// @Description  Modify details of an existing order
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "Order ID"
// @Param        order  body      model.OrderCreateRequest  true  "Updated order details"
// @Success      200   {object}  model.OrderResponse
// @Failure      400   {object}  fiber.Error
// @Failure      404   {object}  fiber.Error
// @Failure      500   {object}  fiber.Error
// @Router       /orders/{id} [put]
func UpdateOrder(c *fiber.Ctx, service *service.OrderService) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var order model.Order
	if err := c.BodyParser(&order); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}
	order.ID = uint(id)
	if err := service.UpdateOrder(&order); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(order)
}

// DeleteOrder removes a order by their ID
//
// @Summary      Delete a Order
// @Description  Remove a order by their ID
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        id   path  int  true  "Order ID"
// @Success      204
// @Failure      500  {object}  fiber.Error
// @Router       /orders/{id} [delete]
func DeleteOrder(c *fiber.Ctx, service *service.OrderService) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := service.DeleteOrder(uint(id)); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}
