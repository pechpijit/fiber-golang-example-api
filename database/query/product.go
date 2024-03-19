package query

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/pechpijit/Fiber_golang_example_api/models"
)

type ProductQuery struct {
	*sql.DB
}

func (db *ProductQuery) GetProducts() ([]models.Product, error) {
	rows, err := db.Query("SELECT id, name, price, discount FROM products_table")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Discount)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if products == nil {
		products = []models.Product{}
	}

	return products, nil
}

func (db *ProductQuery) GetProduct(productId string) (models.Product, error) {
	var product models.Product
	row := db.QueryRow(`SELECT id, name, price, discount FROM products_table WHERE id = $1;`, productId)
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Discount)
	return product, err
}

func (db *ProductQuery) CreateProduct(ctx *fiber.Ctx, productNew *models.Product) error {
	_, err := db.Exec(
		"INSERT INTO products_table (id, name, price, discount) VALUES($1, $2, $3, $4)",
		productNew.ID, productNew.Name, productNew.Price, productNew.Discount)
	return err
}

func (db *ProductQuery) UpdateProduct(productUpdate *models.ProductRequest, productId string) error {
	_, err := db.Exec("UPDATE products_table SET price = $1, discount = $2 WHERE id = $3", productUpdate.Price, productUpdate.Discount, productId)
	return err
}

func (db *ProductQuery) DeleteProduct(productId string) error {
	_, err := db.Exec("DELETE FROM products_table WHERE id = $1", productId)
	return err
}
