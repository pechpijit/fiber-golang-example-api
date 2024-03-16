package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pechpijit/Fiber_golang_example_api/models"
	"strconv"
)

var products []models.Product

func AddMockUpData() {
	products = append(products, models.Product{
		ID:       1,
		Name:     "cc_item_health",
		Price:    500,
		Discount: 10,
	})

	products = append(products, models.Product{
		ID:       2,
		Name:     "cc_target_farm",
		Price:    900,
		Discount: 15,
	})
}

// Handler functions
// GetProducts godoc
// @Summary Get all products
// @Description Get details of all products
// @Tags Products
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Product
// @Router /products [get]
func GetProducts(c *fiber.Ctx) error {
	return c.JSON(products)
}

// Handler functions
// GetProduct godoc
// @Summary Get product by id
// @Description Get details of product
// @Tags Products
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Product
// @Router /products/{productId} [get]
// @Param productId path int true "Product id"
func GetProduct(ctx *fiber.Ctx) error {
	productId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for _, product := range products {
		if product.ID == productId {
			return ctx.JSON(product)
		}
	}

	return ctx.SendStatus(fiber.StatusNotFound)
}

// Handler functions
// DeleteProduct godoc
// @Summary Update product
// @Description Update details of product
// @Tags Products
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 204
// @Router /products/{productId} [delete]
// @Param productId path int true "Product id"
func DeleteProduct(ctx *fiber.Ctx) error {
	productId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, product := range products {
		if product.ID == productId {
			products = append(products[:i], products[i+1:]...)
			return ctx.SendStatus(fiber.StatusNoContent)
		}
	}

	return ctx.SendStatus(fiber.StatusNotFound)
}

// Handler functions
// UpdateProduct godoc
// @Summary Update product
// @Description Update details of product
// @Tags Products
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} models.Product
// @Router /products/{productId} [put]
// @Param productId path int true "Product id"
// @Param json body models.ProductRequest true "Product detail"
func UpdateProduct(ctx *fiber.Ctx) error {
	productId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	productUpdate := new(models.ProductRequest)
	if err := ctx.BodyParser(productUpdate); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, product := range products {
		if product.ID == productId {
			products[i].Price = productUpdate.Price
			products[i].Discount = productUpdate.Discount
			return ctx.JSON(products[i])
		}
	}

	return ctx.SendStatus(fiber.StatusNotFound)
}

// Handler functions
// CreateProduct godoc
// @Summary Create product
// @Description Create details of product
// @Tags Products
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} models.Product
// @Router /products [post]
// @Param json body models.Product true "Product detail"
func CreateProduct(ctx *fiber.Ctx) error {
	productNew := new(models.Product)
	if err := ctx.BodyParser(productNew); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for _, product := range products {
		if product.ID == productNew.ID {
			return ctx.SendStatus(fiber.StatusUnprocessableEntity)
		}
	}

	products = append(products, *productNew)

	return ctx.Status(fiber.StatusCreated).JSON(productNew)
}
