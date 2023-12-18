package controllers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func LogsChecker() {
	dirs := []string{"logs", "syslogs"}

	for _, dir := range dirs {
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				if time.Since(info.ModTime()) > (time.Hour * 24 * 30 * 6) {
					err := os.Remove(path)
					if err != nil {
						log.Printf("Ошибка при удалении файла %s: %v", path, err)
					} else {
						fmt.Printf("Файл %s успешно удален\n", path)
					}
				}
			}
			return nil
		})

		if err != nil {
			log.Fatalf("Ошибка при обходе директории %s: %v", dir, err)
		}
	}
}
