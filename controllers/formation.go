package controllers

import (
	"encoding/json"
	"fmt"
	"lupus-y360/models"
	"os"
)

func JFormation(results []byte, date string) {
	response := models.Response{}
	// Check json data for emptiness
	if len(results) == 32 {
		os.Exit(1)
	}

	// Decode our information from api type []byte to json format
	err := json.Unmarshal(results, &response)
	if err != nil {
		fmt.Println("Ошибка при декодировании JSON:", err)
		return
	}

	// Loop for each event in json data from API
	for _, event := range response.Events {
		EventType := event.EventType
		UniqEvent := event.UniqueID
		UserName := event.UserName
		UserLogin := event.UserLogin
		Path := event.Path
		Ip := event.ClientIP
		Adate := event.Date

		// Checking event type for notify about critical action. Actually we need only some types below
		if (EventType == "fs-copy") || (EventType == "share-change-rights") || (EventType == "fs-set-public") {
			jsonData, err := json.MarshalIndent(event, "", "  ")
			if err != nil {
				fmt.Println("Ошибка при кодировании JSON:", err)
				return
			}

			JJData := string(jsonData)
			EvStr := EventRecon(EventType)
			message := fmt.Sprintf("Framework Lupus🌸 for Yandex360\nUser-%s under loggin (%s)\nTime(%s) %s%s. Ip-%s", UserName, UserLogin, Adate, EvStr, Path, Ip)
			telegramAlerts(message)
			FileWriter(JJData, date, UniqEvent)
		}
	}
}

func EventRecon(EventType string) (EvStr string) {
	switch EventType {
	case "fs-copy":
		EvStr = "copied to local disk, file-"
	case "share-change-rights":
		EvStr = "changed the access level, file-"
	case "fs-set-public":
		EvStr = "published via the link, file-"
	}
	return EvStr
}
