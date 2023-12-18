package main

import (
	"flag"
	"fmt"
	"lupus-y360/controllers"
	"time"
)

func main() {
	FlagStatus := FlagCatcher() // Check flag existing on input
	controllers.Banner()        // Show banner
	YTime, date := GetDate()    // Get current date and time
	controllers.LogsChecker()   // Check logs for date and remove old files

	for i := 0; i < 100000; i++ {
		results := controllers.Y360Connections(i, YTime)
		controllers.JFormation(results, date, FlagStatus)
		i += 100
	}
}

func GetDate() (YTime string, date string) {
	// Get current date
	currentDate := time.Now()
	date = currentDate.Format("2006-01-02")

	// Get current time-4h. Why 4h? Cause Yandex360 Disk have little trouble with refresh data.
	// Actually for last 3h there will be not any information
	currentTime := time.Now()
	oneHourAgo := currentTime.Add(-time.Hour * 4)
	ctime := oneHourAgo.Format("15:04:05")

	YTime = fmt.Sprintf("%sT%sZ", date, ctime)
	return YTime, date
}

func FlagCatcher() (FlagStatus bool) {
	// Definition flag --report or -r
	report := flag.Bool("report", false, "On notification to telegram")
	reportShort := flag.Bool("r", false, "On notification to telegram")

	// CMD command line argument parsing
	flag.Parse()

	// Checking argument for (-r/--report)
	if *report || *reportShort {
		fmt.Println("Active notification")
		FlagStatus = true
	} else {
		fmt.Println("Inactive notification")
		FlagStatus = false
	}
	return FlagStatus
}
