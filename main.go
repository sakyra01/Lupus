package main

import (
	"fmt"
	"lupus-y360/controllers"
	"time"
)

func main() {
	controllers.Banner()
	YTime, date := GetDate()
	for i := 0; i < 100000; i++ {
		results := controllers.Y360Connections(i, YTime)
		controllers.JFormation(results, date)
		i += 100
	}
}

func GetDate() (YTime string, date string) {
	// Get current date
	currentDate := time.Now()
	date = currentDate.Format("2006-01-02")

	// Get current time-4h
	currentTime := time.Now()
	oneHourAgo := currentTime.Add(-time.Hour * 4)
	ctime := oneHourAgo.Format("15:04:05")

	YTime = fmt.Sprintf("%sT%sZ", date, ctime)
	return YTime, date
}
