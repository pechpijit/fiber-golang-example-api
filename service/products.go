package service

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pechpijit/Fiber_golang_example_api/database"
	"github.com/pechpijit/Fiber_golang_example_api/models"
	"log"
)

var products []models.Product

// Handler functions
// GetProducts godoc
// @Summary Get all products
// @Description Get details of all products
// @Tags Products
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Product
// @Router /products [get]
func GetProducts(ctx *fiber.Ctx) error {
	db, _ := database.OpenDBConnection()
	db, errInitDb := database.OpenDBConnection()
	if errInitDb != nil {
		log.Fatal("could not load database", errInitDb.Error())
	}

	products, err := db.GetProducts()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(products)
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
	productId := ctx.Params("id")

	db, errInitDb := database.OpenDBConnection()
	if errInitDb != nil {
		log.Fatal("could not load database", errInitDb.Error())
	}

	product, err := db.GetProduct(productId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.Status(fiber.StatusNotFound).SendString("Product not found")
		}
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(product)
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
	productId := ctx.Params("id")

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
	productId := ctx.Params("id")

	productUpdate := new(models.ProductRequest)
	if err := ctx.BodyParser(productUpdate); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	db, errInitDb := database.OpenDBConnection()
	if errInitDb != nil {
		log.Fatal("could not load database", errInitDb.Error())
	}

	_, errGetProduct := db.GetProduct(productId)
	if errGetProduct != nil {
		if errors.Is(errGetProduct, sql.ErrNoRows) {
			return ctx.Status(fiber.StatusNotFound).SendString("Product not found")
		}
		return ctx.Status(fiber.StatusInternalServerError).SendString(errGetProduct.Error())
	}

	err := db.UpdateProduct(productUpdate, productId)
	if err == nil {
		product, _ := db.GetProduct(productId)
		return ctx.JSON(product)
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

	productNew.ID = uuid.New().String()

	db, errInitDb := database.OpenDBConnection()
	if errInitDb != nil {
		log.Fatal("could not load database", errInitDb.Error())
	}

	err := db.CreateProduct(ctx, productNew)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(productNew)
}
