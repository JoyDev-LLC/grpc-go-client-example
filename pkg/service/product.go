package service

import (
	"product/pkg/repository"
	"product/products"
	"context"
	"io"
	"log"
)

type ProductService struct {
	repo repository.Product
	client products.ProductsClient
}

func NewProductService(repo repository.Product, client products.ProductsClient) *ProductService{
	return &ProductService{repo :repo, client: client}
}

func GetProduct(client products.ProductsClient, id uint32) (*products.GetProductResp, error) {
	product, err := client.GetProduct(context.Background(), &products.GetProductReq{Id: id})
	if err != nil {
		 log.Printf("failed to get product: %v", err)
		 return product, err
	 }
	return product, nil
 }

func GetProducts(client products.ProductsClient, search string, limit uint32) ([]*products.Product, error) {
  	productList, err := client.GetProducts(context.Background(), &products.GetProductsReq{Search: search, Limit: limit})
	if err != nil {
		 log.Printf("failed to get book list: %v", err)
		 return nil, err
	}

 	prods := make([]*products.Product, 0)
	for {
		product, err := productList.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("failed to get product: %v", err)
			return nil, err
		}
		prods = append(prods, product)
	}
 
	return prods, nil
}
 
func CreateProducts(client products.ProductsClient, product_list []*products.Product) (*products.CreateProductsResp, error) {
	stream, err := client.CreateProducts(context.Background())
	if err != nil {
		 log.Printf("%v.CreateProducts(_) = _, %v", client, err)
	 return nil, err
	}

	for _, product := range product_list {
		if err := stream.Send(product); err != nil {
			log.Printf("%v.Send(%v) = %v", stream, product_list, err)
		 return nil, err
		}
	}

	productList, err := stream.CloseAndRecv()
	if err != nil {
		 log.Printf("failed to create product list: %v", err)
		return productList, err
	}
 
    return productList, nil
}

 
 func CreateProductsInStream(client products.ProductsClient, product_list []*products.Product) ([]*products.Product, error) {
	stream, err := client.CreateProducts(context.Background())
	if err != nil {
		log.Printf("%v.CreateProducts(_) = _, %v", client, err)
		return nil, err
	}
	prods := make([]*products.Product, 0)
	for {
		prductsList, err := stream.Recv()
		if err == io.EOF {
		  return nil, err
		}
		if err != nil {
		  return nil, err
		}

		for _, product := range prductsList {
		  if err := stream.Send(product); err != nil {
			return nil, err
		  }
		  prods = append(prods, product)
		}		
	}

	return prods, nil

}
	

 