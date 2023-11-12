package handler

import (
	"product"
	"product/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

type GetAllProductsResponse struct {
	Data []product.Product `json:"data"`
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		products := api.Group("/products")
		{
			products.POST("/", h.createProducts)
			products.POST("/create_in_steam", h.createProductsInStream)
			products.GET("/", h.getAllProducts)
			products.GET("/:id", h.getProduct)
		}
	}

	return router
}
