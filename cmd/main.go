package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"video-downloader/utils"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите URL видео для скачивания:")
	scanner.Scan()             // Ожидаем ввод от пользователя
	videoURL := scanner.Text() // Получаем введённую строку

	downloadsDir := "./downloads"

	// Проверяем, существует ли папка, если нет то создаём её
	if _, err := os.Stat(downloadsDir); os.IsNotExist(err) {
		err := os.Mkdir(downloadsDir, 0755)
		if err != nil {
			log.Fatal("Не удалось создать папку для видео: ", err)
		}
	}

	// Генерируем имя файла на основе URL
	// Извлекаем имя файла из URL, заменяем символы, которые могут быть недопустимыми в имени файла
	fileName := strings.TrimSuffix(filepath.Base(videoURL), filepath.Ext(videoURL)) + ".mp4"

	// Путь до сохранения видео
	outputPath := filepath.Join(downloadsDir, fileName)

	if err := utils.CheckFFmpeg(); err != nil {
		log.Fatal(err)
	}

	if err := utils.RunFFmpeg(videoURL, outputPath); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Видео успешно скачанно в файл", outputPath)
	fmt.Println("Программа завершена. Нажмите Enter для выхода.")
	fmt.Scanln() // Ждём ввода от пользователя

}
