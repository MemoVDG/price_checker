package handlers

import (
	"github.com/MemoVDG/price_checker/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func HandleTelegramWebhook(c *gin.Context) {
	var update, err = utils.ParserTelegramRequest(c.Request)

	if err != nil {
		log.Printf("Error parsings %s", err.Error())
		return
	}

	log.Println(update.Message.Text)

	var telegramaResposeBody, errTelegram = utils.SendTextToTelegram(update.Message.Chat.Id, "Tests")
	if errTelegram != nil {
		log.Printf("Got an error sending message to API %s %s", errTelegram.Error(), telegramaResposeBody)
		return
	} else {
		log.Printf("Message sent successfully")
	}

}
