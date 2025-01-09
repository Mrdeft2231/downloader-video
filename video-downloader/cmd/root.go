package cmd

import (
	"fmt"
	"video-downloader/video-downloader/service"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "downloader",
	Short: "Менеджер скачивания видео",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Использования: downloader <URL>, <output>")
			return
		}

		url := args[0]
		output := args[1]

		err := service.DowloadVideo(url, output)
		if err != nil {
			fmt.Println("ошибка при скачивании файла: %v\n", err)
		} else {
			fmt.Println("Видео успешно скачано!")
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}
