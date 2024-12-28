package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sms2sakthivel/order-manager/orders/model"
	"github.com/sms2sakthivel/order-manager/orders/service"
)

// GetAllCarts retrieves all carts
//
// @Summary      Get All Carts
// @Description  Retrieve a list of all carts
// @Tags         Carts
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.CartResponse
// @Failure      500  {object}  fiber.Error
// @Router       /carts [get]
func GetAllCarts(c *fiber.Ctx, service *service.CartService) error {
	carts, err := service.GetCarts()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(carts)
}

// GetCartByID retrieves a cart by their ID
//
// @Summary      Get Cart by ID
// @Description  Retrieve a cart by their ID
// @Tags         Carts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Cart ID"
// @Success      200  {object}  model.CartResponse
// @Failure      404  {object}  fiber.Error
// @Failure      500  {object}  fiber.Error
// @Router       /carts/{id} [get]
func GetCartByID(c *fiber.Ctx, service *service.CartService) error {
	id, _ := strconv.Atoi(c.Params("id"))
	cart, err := service.GetCart(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Cart not found")
	}
	return c.JSON(cart)
}

// CreateCart adds a new cart
//
// @Summary      Create a New Cart
// @Description  Add a new cart to the system
// @Tags         Carts
// @Accept       json
// @Produce      json
// @Param        cart  body      model.CartCreateRequest  true  "Cart details"
// @Success      201   {object}  model.CartResponse
// @Failure      400   {object}  fiber.Error
// @Failure      500   {object}  fiber.Error
// @Router       /carts [post]
func CreateCart(c *fiber.Ctx, service *service.CartService) error {
	var cartReq model.CartCreateRequest
	if err := c.BodyParser(&cartReq); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}
	cart, err := service.CreateCart(&cartReq)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(cart)
}

// UpdateCart modifies details of an existing cart
//
// @Summary      Update an Existing Cart
// @Description  Modify details of an existing cart
// @Tags         Carts
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "Cart ID"
// @Param        cart  body      model.CartUpdateRequest  true  "Updated cart details"
// @Success      200   {object}  model.CartResponse
// @Failure      400   {object}  fiber.Error
// @Failure      404   {object}  fiber.Error
// @Failure      500   {object}  fiber.Error
// @Router       /carts/{id} [put]
func UpdateCart(c *fiber.Ctx, service *service.CartService) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var cartReq model.CartUpdateRequest
	if err := c.BodyParser(&cartReq); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}
	cartReq.ID = uint(id)
	cart, err := service.UpdateCart(&cartReq)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(cart)
}

// DeleteCart removes a cart by their ID
//
// @Summary      Delete a Cart
// @Description  Remove a cart by their ID
// @Tags         Carts
// @Accept       json
// @Produce      json
// @Param        id   path  int  true  "Cart ID"
// @Success      204
// @Failure      500  {object}  fiber.Error
// @Router       /carts/{id} [delete]
func DeleteCart(c *fiber.Ctx, service *service.CartService) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := service.DeleteCart(uint(id)); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}
