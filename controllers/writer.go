package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FileRuler(JJData, date, UniqEvent, messageSyslog string) {
	// Get Filepath and make name for file for writing mode
	logfilePath := fmt.Sprintf("logs/logfile-%s.log", date)
	syslogPath := fmt.Sprintf("syslogs/syslog-%s.txt", date)

	// Checking file existence
	Status := FileExist(logfilePath, UniqEvent)
	FileExist(syslogPath, UniqEvent)
	if Status == false {
		return
	}

	FileWriter(logfilePath, JJData)
	FileWriter(syslogPath, messageSyslog)
}

func FileWriter(filePath string, fileContent string) {
	// Open file for extra writing
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close() // Close file at the end
	// Add information in file
	_, err = file.WriteString(fileContent)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}
}

func FileExist(filePath, UniqEvent string) (LogfileStatus bool) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Создание файла, если он не существует
		_, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Ошибка при создании файла:", err)
		}
		LogfileStatus = true
		return LogfileStatus
	} else {
		LogfileStatus = UniqIDCompare(filePath, UniqEvent)
		return LogfileStatus
	}
}

func UniqIDCompare(filePath, UniqEvent string) (Status bool) {
	Status = true

	// String we need compare for
	searchString := UniqEvent

	// Open file wor reading
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
	}
	defer file.Close() // Close file at the end

	// Make scanner for reading file
	scanner := bufio.NewScanner(file)

	// String enumeration in file
	for scanner.Scan() {
		line := scanner.Text()

		// Check value "uniqId" in string
		if strings.Contains(line, "uniqId") {
			// Compare value "uniqId" with current string
			if strings.Contains(line, searchString) {
				fmt.Println("Лог уже есть в записи")
				Status = false
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
	}
	return Status
}
