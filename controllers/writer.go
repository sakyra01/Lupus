package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FileWriter(JJData, date, UniqEvent string) {
	// Установите путь и имя файл, а для записи
	filePath := fmt.Sprintf("logs/logfile-%s.log", date)

	// Checking file existence
	Status := FileExist(filePath, UniqEvent)

	if Status == false {
		return
	}
	// Открытие файла в режиме дозаписи
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	// Информация, которую нужно записать
	fileContent := JJData

	// Запись информации в файл
	_, err = file.WriteString(fileContent)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}
}

func FileExist(filePath, UniqEvent string) (WriterStatus bool) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Создание файла, если он не существует
		_, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Ошибка при создании файла:", err)
		}
		WriterStatus = true
		return WriterStatus
	} else {
		WriterStatus = UniqIDCompare(filePath, UniqEvent)
		return WriterStatus
	}
}

func UniqIDCompare(filePath, UniqEvent string) (Status bool) {
	Status = true

	// Строка, которую нужно сравнить
	searchString := UniqEvent

	// Открытие файла для чтения
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
	}
	defer file.Close()

	// Создание сканнера для чтения файла построчно
	scanner := bufio.NewScanner(file)

	// Перебор каждой строки файла
	for scanner.Scan() {
		line := scanner.Text()

		// Проверка наличия поля "uniqId" в строке
		if strings.Contains(line, "uniqId") {
			// Сравнение переданной строки с полем "uniqId" в текущей строке
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
