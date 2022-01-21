package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/MemoVDG/price_checker/structs"
	"github.com/MemoVDG/price_checker/utils/scrappers"
	"github.com/gin-gonic/gin"
)

func storeConsulting(store string, url string) (int, string) {
	switch store {
	case "amazon":
		return scrappers.AmazonScrapping(url)
	case "mercardolibre":
		fmt.Println("two")
	default:
		fmt.Println("error")
	}
	return 0, "error"
}

func ProductRequest(c *gin.Context) {
	var r = c.Request
	var body structs.AddProductRequest

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		//TODO: Handle error
		return
	}
	price, productName := storeConsulting(body.Store, body.Url)

	//This is for telegram
	//pat := regexp.MustCompile("\\/(.*?)\\:")
	//match := pat.FindStringSubmatch(body.Url)
	//storeConsulting(strings.ToLower(match[1]))

	c.JSON(200, gin.H{
		"product": productName,
		"price":   price,
	})
}
