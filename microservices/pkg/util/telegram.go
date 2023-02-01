package util

import (
	"examples/microservices/pkg/setting"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	// Import all the commands you wish to use
)

func SendMessageToTelegram(text string) (string, error) {
	// log.Printf("Sending %s to chat_id: %d", text, chatId)
	var telegramApi string = "https://api.telegram.org/bot" + setting.AppSettings.GeneralSettings.TelegramToken + "/sendMessage"
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id": {setting.AppSettings.GeneralSettings.TelegramChatId},
			"text":    {text},
		})

	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("error in parsing telegram answer %s", errRead.Error())
		return "", err
	}
	bodyString := string(bodyBytes)
	log.Printf("Body of Telegram Response: %s", bodyString)

	return bodyString, nil
}
