package product

type Product struct {
	Id uint32 `json:"id" db:"id"`
	Name string `json:"name" binding:"required"`
	Price float32 `json:"price" binding:"required"`
}

type ProductInput struct {
	Name string `json:"name"`
	Price float32 `json:"price"`
}

type CreateProductsInput struct {
	Products []ProductInput
}