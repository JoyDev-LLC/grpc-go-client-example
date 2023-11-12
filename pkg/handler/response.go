package handler

import (
	"log"
	"github.com/gin-gonic/gin"
)

type Error struct {
	Message string `json:"message"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Printf("error: %s", message)
	c.AbortWithStatusJSON(statusCode, Error{ message })	
}
