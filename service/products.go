package service

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/google/uuid"
	"github.com/pechpijit/Fiber_golang_example_api/database"
	"github.com/pechpijit/Fiber_golang_example_api/middleware"
	"github.com/pechpijit/Fiber_golang_example_api/models"
	"github.com/pechpijit/Fiber_golang_example_api/repository"
	"github.com/pechpijit/Fiber_golang_example_api/response"
	"log"
)

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
		return response.RespondError(ctx, fiber.StatusInternalServerError, err.Error())
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
	productId := utils.CopyString(ctx.Params("id"))

	db, errInitDb := database.OpenDBConnection()
	if errInitDb != nil {
		log.Fatal("could not load database", errInitDb.Error())
	}

	product, errGetProduct := db.GetProduct(productId)
	if errGetProduct != nil {
		if errors.Is(errGetProduct, sql.ErrNoRows) {
			return response.RespondError(ctx, fiber.StatusNotFound, "Product not found")
		}
		return response.RespondError(ctx, fiber.StatusInternalServerError, errGetProduct.Error())
	}

	return response.RespondSuccess(ctx, fiber.StatusOK, product)
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
	credentialNeed := repository.ProductDeleteCredential
	if !middleware.JWTCheckPermission(ctx, &credentialNeed) {
		return response.RespondError(ctx, fiber.StatusUnauthorized, "unauthorize access")
	}

	productId := utils.CopyString(ctx.Params("id"))

	db, errInitDb := database.OpenDBConnection()
	if errInitDb != nil {
		log.Fatal("could not load database", errInitDb.Error())
	}

	err := db.DeleteProduct(productId)
	if err != nil {
		return response.RespondError(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return ctx.SendStatus(fiber.StatusNoContent)
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
	credentialNeed := repository.ProductUpdateCredential
	if !middleware.JWTCheckPermission(ctx, &credentialNeed) {
		return response.RespondError(ctx, fiber.StatusUnauthorized, "unauthorize access")
	}

	productId := utils.CopyString(ctx.Params("id"))

	productUpdate := new(models.ProductRequest)
	if err := ctx.BodyParser(productUpdate); err != nil {
		return response.RespondError(ctx, fiber.StatusBadRequest, err.Error())
	}

	db, errInitDb := database.OpenDBConnection()
	if errInitDb != nil {
		log.Fatal("could not load database", errInitDb.Error())
	}

	_, errGetProduct := db.GetProduct(productId)
	if errGetProduct != nil {
		if errors.Is(errGetProduct, sql.ErrNoRows) {
			return response.RespondError(ctx, fiber.StatusNotFound, "Product not found")
		}
		return response.RespondError(ctx, fiber.StatusInternalServerError, errGetProduct.Error())
	}

	err := db.UpdateProduct(productUpdate, productId)
	if err == nil {
		product, _ := db.GetProduct(productId)
		return response.RespondSuccess(ctx, fiber.StatusOK, product)
	}

	return response.RespondSuccess(ctx, fiber.StatusNotFound, "")
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
	credentialNeed := repository.ProductCreateCredential
	if !middleware.JWTCheckPermission(ctx, &credentialNeed) {
		return response.RespondError(ctx, fiber.StatusUnauthorized, "unauthorize access")
	}

	productNew := new(models.Product)
	if err := ctx.BodyParser(productNew); err != nil {
		return response.RespondError(ctx, fiber.StatusBadRequest, err.Error())
	}

	productNew.ID = uuid.New().String()

	db, errInitDb := database.OpenDBConnection()
	if errInitDb != nil {
		log.Fatal("could not load database", errInitDb.Error())
	}

	err := db.CreateProduct(ctx, productNew)
	if err != nil {
		return response.RespondError(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return response.RespondSuccess(ctx, fiber.StatusCreated, productNew)
}
