package controllers

import (
	"fmt"
	"time"
)

func Notify(EventType, UserName, UserLogin, Path, ClientIP string, ADate time.Time) {
	if (EventType == "fs-copy") || (EventType == "share-change-rights") || (EventType == "fs-set-public") || (EventType == "share-activate-invite") {
		EvStr := EventRecon(EventType)
		messageTelegram := fmt.Sprintf("Framework Lupus🌸 for Yandex360\nUser-%s under loggin (%s)\nTime(%v) %s%s. Ip-%s", UserName, UserLogin, ADate, EvStr, Path, ClientIP)
		telegramAlerts(messageTelegram)
	}
}

// Checking event type for notify about critical action

func EventRecon(EventType string) (EvStr string) {
	switch EventType {
	case "fs-copy":
		EvStr = "copied to local disk, file-"
	case "share-change-rights":
		EvStr = "changed the access level, file-"
	case "fs-set-public":
		EvStr = "published via the link, file-"
	case "share-activate-invite":
		EvStr = "accepted the invitation, file-"
	}
	return EvStr
}
