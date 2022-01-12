package main

import (
	"github.com/MemoVDG/price_checker/handlers"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/test", ch)
		v1.GET("/", hello)
		v1.POST("/", handlers.HandleTelegramWebhook)
	}

	router.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Salem",
	})
}

func ch(c *gin.Context) {
	log.Printf(os.Getenv("TELEGRAM_BOT_TOKEN"))
	c.JSON(200, gin.H{
		"message": "Chaleeee",
	})
}
