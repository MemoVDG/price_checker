package main

import (
	"encoding/json"
	"fmt"
	"github.com/MemoVDG/price_checker/handlers"
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/", handlers.HandleTelegramWebhook)
		v1.POST("/test", regex)
	}

	router.Run(":8000")
}

type Test struct {
	Store string
}

func hitToAmazonURL(url string) {
	// Amazon done in amazonScrappingFile.go
}

func storeConsulting(store string) {
	fmt.Println(store)
	switch store {
	case "amazon":
		hitToAmazonURL("ssss")
	case "mercardolibre":
		fmt.Println("two")
	default:
		fmt.Println("error")
	}
}

func regex(c *gin.Context) {
	var r = c.Request
	var test Test

	err := json.NewDecoder(r.Body).Decode(&test)
	if err != nil {
		return
	}
	pat := regexp.MustCompile("\\/(.*?)\\:")
	match := pat.FindStringSubmatch(test.Store)
	if len(match) > 1 {
		storeConsulting(strings.ToLower(match[1]))
	}

	c.JSON(200, gin.H{
		"message": "Salem",
	})
}
