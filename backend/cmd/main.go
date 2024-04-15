package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Email string `json:"email" binding:"required"`
}

func main() {
	router := gin.Default()

	router.POST("/api/v1/authenticate", func(c *gin.Context) {
		var requestBody RequestBody

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"email": requestBody.Email,
		})
	})

	router.Run(":3333")
}
