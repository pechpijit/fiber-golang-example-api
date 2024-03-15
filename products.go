package main

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
}

type ProductRequest struct {
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
}

var products []Product

// Handler functions
// getProducts godoc
// @Summary Get all products
// @Description Get details of all products
// @Tags Products
// @Accept  json
// @Produce  json
// @Success 200 {array} Product
// @Router /products [get]
func getProducts(c *fiber.Ctx) error {
	return c.JSON(products)
}

// Handler functions
// getProduct godoc
// @Summary Get product by id
// @Description Get details of product
// @Tags Products
// @Accept  json
// @Produce  json
// @Success 200 {object} Product
// @Router /products/{productId} [get]
// @Param productId path int true "Product id"
func getProduct(ctx *fiber.Ctx) error {
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
// deleteProduct godoc
// @Summary Update product
// @Description Update details of product
// @Tags Products
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 204
// @Router /products/{productId} [delete]
// @Param productId path int true "Product id"
func deleteProduct(ctx *fiber.Ctx) error {
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
// updateProduct godoc
// @Summary Update product
// @Description Update details of product
// @Tags Products
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} Product
// @Router /products/{productId} [put]
// @Param productId path int true "Product id"
// @Param json body ProductRequest true "Product detail"
func updateProduct(ctx *fiber.Ctx) error {
	productId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	productUpdate := new(ProductRequest)
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
// createProduct godoc
// @Summary Create product
// @Description Create details of product
// @Tags Products
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} Product
// @Router /products [post]
// @Param json body Product true "Product detail"
func createProduct(ctx *fiber.Ctx) error {
	productNew := new(Product)
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
