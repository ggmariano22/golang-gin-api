package repository

import (
	"context"
	"database/sql"
	"fmt"
	"golang/api-go-routine/config"
	"golang/api-go-routine/models"
	"log"
)

var repository *sql.DB
var erro error

func Insert(ctx context.Context, p models.Product) (int64, error) {
	query := "INSERT INTO products (description, un, value) VALUES (?, ?, ?)"
	res, err := repository.ExecContext(ctx, query, p.Description, p.UN, p.Value)

	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func Select(ctx context.Context, page int, limit int) models.Products {
	offset := (page - 1) * limit
	query := fmt.Sprintf("SELECT * FROM products LIMIT %d, %d", offset, limit)

	res, err := repository.QueryContext(ctx, query)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()

	var product models.Product
	var products models.Products
	for res.Next() {

		res.Scan(&product.ID, &product.Description, &product.UN, &product.Value)
		products.Data = append(products.Data, product)
	}

	return products
}

func init() {
	repository, erro = config.GetConnection()

	if erro != nil {
		fmt.Println(erro)
	}
}
