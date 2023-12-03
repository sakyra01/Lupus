package models

import "time"

type Event struct {
	EventType            string    `json:"eventType"`
	OrgID                int       `json:"orgId"`
	UserUID              string    `json:"userUid"`
	UserLogin            string    `json:"userLogin"`
	UserName             string    `json:"userName"`
	OwnerUID             string    `json:"ownerUid"`
	OwnerLogin           string    `json:"ownerLogin"`
	OwnerName            string    `json:"ownerName"`
	ResourceFileID       string    `json:"resourceFileId"`
	Path                 string    `json:"path"`
	Size                 string    `json:"size"`
	Rights               string    `json:"rights"`
	RequestID            string    `json:"requestId"`
	UniqueID             string    `json:"uniqId"`
	ClientIP             string    `json:"clientIp"`
	Date                 time.Time `json:"date"`
	LastModificationDate time.Time `json:"lastModificationDate"`
}

type Response struct {
	Events        []Event `json:"events"`
	NextPageToken string  `json:"nextPageToken"`
}
