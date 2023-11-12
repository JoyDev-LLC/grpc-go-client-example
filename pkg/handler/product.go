package handler

import (
	"product/products"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllProducts(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "limit of valid type")
		return
	}
	products, err := h.services.Product.GetProducts(c.Query("search"), uint32(limit))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetAllProductsResponse{Data: products})
}

func (h *Handler) getProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id of valid type")
		return
	}
	product, err := h.services.Product.GetProduct(c, uint32(id))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *Handler) createProducts(c *gin.Context) {
	var p_input []*products.Product
	if err := c.BindJSON(&p_input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
    
	products, err := h.services.Product.CreateProducts(c, p_input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *Handler) createProductsInStream(c *gin.Context) {
	var p_input []*products.Product
	if err := c.BindJSON(&p_input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
    
	products, err := h.services.Product.CreateProductsInStream(c, p_input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}

