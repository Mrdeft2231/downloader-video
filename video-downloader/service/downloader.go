package service

import (
	"fmt"
	"video-downloader/video-downloader/utils"
)

func DowloadVideo(url, output string) error {
	fmt.Println("начинаем скачивать видео...\nURL: %s\nOutput: %s\n", url, output)

	if err := utils.CheckFFmpeg(); err != nil {
		return fmt.Errorf("FFmpeg не найден: %v", err)
	}

	if err := utils.RunFFmpeg(url, output); err != nil {
		return fmt.Errorf("ошибка при работе FFmpge: %v", err)
	}

	fmt.Println("Скачивания завершено")
	return nil
}
