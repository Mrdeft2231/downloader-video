package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func CheckFFmpeg() error {
	ffmpegPath := filepath.Join("..", "ffmpeg.exe") // Указываем путь к ffmpeg в папке с проектом
	_, err := os.Stat(ffmpegPath)                   // Проверяем, существует ли файл ffmpeg.exe

	if err != nil {
		return errors.New("FFmpeg не найден в папке с проектом. Убедитесь, что ffmpeg.exe находится в той же папке, что и проект")
	}
	return nil
}

func RunFFmpeg(url string) error {
	ffmpegPath := filepath.Join(".", "ffmpeg.exe") // Путь к ffmpeg в текущей папке

	output := generateOutputFileName(url)

	cmd := exec.Command(ffmpegPath, "-i", url, "-c", "copy", output)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func generateOutputFileName(url string) string {
	invalidChars := []string{":", "/", "\\", "?", "*", "<", ">", "|", "\""}
	baseName := url
	for _, char := range invalidChars {
		baseName = strings.ReplaceAll(baseName, char, "_")
	}
	return fmt.Sprintf("videos/%s.mp4", baseName)
}
