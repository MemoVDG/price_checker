package utils

import (
	"encoding/json"
	"fmt"
	"github.com/MemoVDG/price_checker/models"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func ParserTelegramRequest(r *http.Request) (*models.Update, error) {
	var update models.Update
	fmt.Print(r)
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("Could not decode incoming update %s", err.Error())
		return nil, err
	}

	return &update, nil
}

func SendTextToTelegram(chatId int, text string) (string, error) {
	log.Printf("Sending %s to chat_id: %d", text, chatId)
	var telegramApi string = "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_BOT_TOKEN") + "/sendMessage"
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id": {strconv.Itoa(chatId)},
			"text":    {text},
		})

	if err != nil {
		log.Printf("Error when posting text to chat %s", err.Error())
		return "", err
	}

	defer response.Body.Close()
	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("Error in parsing telegram answer %s", errRead.Error())
		return "", err
	}
	bodyString := string(bodyBytes)
	log.Printf("Body of Telegram response %s", bodyString)
	return bodyString, nil
}
