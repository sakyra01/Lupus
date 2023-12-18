package controllers

import (
	"encoding/json"
	"fmt"
	"lupus-y360/models"
	"os"
)

func JFormation(results []byte, date string, FlagStatus bool) {
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
		OrgID := event.OrgID
		UserUID := event.UserUID
		UserLogin := event.UserLogin
		UserName := event.UserName
		OwnerUID := event.OwnerUID
		OwnerLogin := event.OwnerLogin
		OwnerName := event.OwnerName
		ResourceFileID := event.ResourceFileID
		Path := event.Path
		Size := event.Size
		Rights := event.Rights
		RequestID := event.RequestID
		UniqueID := event.UniqueID
		ClientIP := event.ClientIP
		ADate := event.Date
		LastModificationDate := event.LastModificationDate

		//	Standard type of events which we work with
		EventsList := []string{"fs-copy", "fs-mkdir", "fs-move", "fs-set-public", "fs-store", "fs-trash-append",
			"fs-trash-drop", "fs-trash-drop-all", "share-activate-invite", "share-change-rights",
			"share-change-invite-rights", "share-create-group", "share-invite-user", "fs-rm"}
		//	Searching new type of events which could be added in API Yandex360
		NewEvent(EventsList, EventType)

		// Notification on telegram
		if FlagStatus == true {
			Notify(EventType, UserName, UserLogin, Path, ClientIP, ADate)
		}

		// Encode struct information to json format
		jsonData, err := json.MarshalIndent(event, "", "  ")
		if err != nil {
			fmt.Println("JSON Encode error:", err)
			return
		}

		JJData := string(jsonData)
		messageSyslog := fmt.Sprintf("CEF:0|event=%s|orgID=%d|userUID=%s|userLogin=%s|username=%s|ownerUID=%s|ownerLogin=%s|ownerName=%s|resourceFieldID=%s|path=%s|size=%s|rights=%s|requestID=%s|uniqueID=%s|clientIP=%s|date=%v|lastModificationDate=%v\n", EventType, OrgID, UserUID, UserLogin, UserName, OwnerUID, OwnerLogin, OwnerName, ResourceFileID, Path, Size, Rights, RequestID, UniqueID, ClientIP, ADate, LastModificationDate)
		FileRuler(JJData, date, UniqueID, messageSyslog)
	}
}

func NewEvent(EventsList []string, CurrentEvent string) {
	found := false
	for _, event := range EventsList {
		if event == CurrentEvent {
			found = true
			break
		}
	}
	if !found {
		newMessage := fmt.Sprintf("A new Type Event has appeard in Yandex360 - %s !", CurrentEvent)
		//	Notify about new event type
		telegramAlerts(newMessage)
	}
}
