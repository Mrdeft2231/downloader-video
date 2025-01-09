package utils

import (
	"errors"
	"os"
	"os/exec"
)

func CheckFFmpeg() error {
	ffmpegPath := `C:\Users\deftm\OneDrive\Рабочий стол\ffmpeg-master-latest-win64-gpl\bin\ffmpeg.exe`

	_, err := exec.LookPath(ffmpegPath)
	if err != nil {
		return errors.New("FFmpeg не установлен. Установите его перед использованием программы")
	}
	return nil
}

func RunFFmpeg(url, output string) error {
	ffmpegPath := `C:\Users\deftm\OneDrive\Рабочий стол\ffmpeg-master-latest-win64-gpl\bin\ffmpeg.exe`
	cmd := exec.Command(ffmpegPath, "-i", `"`+url+`"`, "-c", "copy", output)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
