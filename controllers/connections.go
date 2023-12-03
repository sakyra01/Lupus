package controllers

import (
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func Y360Connections(step int, YTime string) (results []byte) {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	orgID := os.Getenv("orgID")
	Token := os.Getenv("Bearer")
	uri := fmt.Sprintf("https://api360.yandex.net/security/v1/org/%s/audit_log/disk?pageSize=100&pageToken=%d&afterDate=%s", orgID, step, YTime)

	// Build request to yandex360 api source
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		os.Exit(1)
	}
	req.Header.Add("Authorization", Token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	results, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		os.Exit(1)
	}
	return results
}

func telegramAlerts(message string) {
	botToken := os.Getenv("Telegram_Token")
	channelID := os.Getenv("Telegram_chanel_ID")

	telegramURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	params := url.Values{
		"chat_id": {channelID},
		"text":    {message},
	}

	_, err := http.PostForm(telegramURL, params)
	if err != nil {
		fmt.Println("Ошибка при отправке сообщения в Telegram:", err)
		return
	}
}
