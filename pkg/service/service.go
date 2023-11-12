package service

import (
	"product"
	"product/pkg/repository"
	"product/products"
)


type Product interface{
	GetProduct(client products.ProductsClient, id uint32) (*products.GetProductResp, error),
	GetProducts(client products.ProductsClient, search string, limit uint32) ([]*products.Product, error),
	CreateProducts(client products.ProductsClient, product_list []*products.Product) (*products.CreateProductsResp, error),
	CreateProductsInStream(client products.ProductsClient, product_list []*products.Product) ([]*products.Product, error)
}

type Service struct{
	Product

}

func NewService(repos *repository.Repository, client products.ProductsClient) *Service{
	return &Service{
		Product: product.NewProductService(repos.Product, client),
	}
}
