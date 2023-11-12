package repository

import (
	"product"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) Create(product_list product.CreateProductsInput) ([]product.Product, error) {
	products := make([]product.Product, 0, len(product_list.Products))

	for _, product := range product_list.Products {
		var product_item product.Product
		
		createListQuery := fmt.Sprintf("INSERT INTO %s (name, price) VALUES ($1, $2) RETURNING id, name, price", productsTable)

		row := r.db.QueryRow(createListQuery, product.Name, product.Price)
		if err := row.Scan(&product_item.Id, &product_item.Name, &product_item.Price); err != nil {
			return nil, err
		}
		products = append(products, product_item)
	}

	return products, nil
}

func (r *ProductPostgres) GetAll(search string, limit uint32) ([]product.Product, error) {
	var lists []product.Product

	query := fmt.Sprintf(`SELECT p.id, p.name, p.price FROM %s p WHERE p.name LIKE CONCAT('%%', $1::text, '%%') LIMIT $2`, productsTable)
	err := r.db.Select(&lists, query, search, limit)

	return lists, err
}

func (r *ProductPostgres) GetById(id uint32) (product.Product, error) {
	var list product.Product

	query := fmt.Sprintf("SELECT p.id, p.name, p.price FROM %s p WHERE p.id = $1", productsTable)
	err := r.db.Get(&list, query, id)

	return list, err
}
