package main

import (
	"github.com/MemoVDG/price_checker/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/", handlers.HandleTelegramWebhook)
		v1.POST("/product", handlers.ProductRequest)
	}

	router.Run(":8000")
}
