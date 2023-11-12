package repository

import (
	"product"
	"github.com/jmoiron/sqlx"
)

type Product interface {
	Create(product product.CreateProductsInput) ([]product.Product, error)
	GetAll(search string, limit uint32) ([]product.Product, error)
	GetById(id uint32) (product.Product, error)
}

type Repository struct {
	Product
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Product: NewProductPostgres(db),
	}
}
