package main

import (
	"fmt"
	"video-downloader/video-downloader/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println("Ошибка: ", err)
	}
}
