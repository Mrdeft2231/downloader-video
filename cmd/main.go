package main

import (
	"fmt"
	"os"
	"sync"
	"video-downloader/utils"
)

func main() {

	createOutputDirectory()

	var links []string = []string{
		"https://v.ozone.ru/vod/video-33/01HGM1M67HMBS3JSHBWGX5WBHC/master.m3u8",
		"https://v.ozone.ru/vod/video-35/01HGM1HNY0THZ3T9PS7220Z0N2/master.m3u8",
		"https://v.ozone.ru/vod/video-32/01HGM1G25P13DWA5PKJMMYA9GJ/master.m3u8",
		"https://v.ozone.ru/vod/video-30/01HGM1CPYYGZSVFXRDV6FNDZ2Y/master.m3u8",
		"https://v.ozone.ru/vod/video-30/01HGM1AV4550TDWX4Q4B89TFWV/master.m3u8",
		"https://v.ozone.ru/vod/video-32/01HGM18EVVD0DJ8F9WHYQTAJKZ/master.m3u8",
		"https://v.ozone.ru/vod/video-32/01HGM13K7P03YQGMVBSQDXPPKS/master.m3u8",
		"https://v.ozone.ru/vod/video-34/01HGM113WQ7V13NK2TBS07HAM1/master.m3u8",
		"https://v.ozone.ru/vod/video-35/01HGM0YSMXY8FHS63ZVKNAE187/master.m3u8",
		"https://v.ozone.ru/vod/video-29/01HGM0Q8VWPZ09HC8KADSNQ6SE/master.m3u8",
		"https://v.ozone.ru/vod/video-33/01HGM0N74W46R51GXJNNDB48P6/master.m3u8",
		"https://v.ozone.ru/vod/video-28/01HGM0K0AJQKRV0KEEH0CXTDNT/master.m3u8",
		"https://v.ozone.ru/vod/video-31/01HGM0FRE4BXH6GZZJNE7KGX2H/master.m3u8",
		"https://v.ozone.ru/vod/video-33/01HGM0DXYEK2TP7Z3NZ9Q221KA/master.m3u8",
		"https://v.ozone.ru/vod/video-32/01HGM0BVDVZVKJ67JQMH5WTN2T/master.m3u8",
		"https://v.ozone.ru/vod/video-33/01HGM07XAYBE5RBR41FEG3DDF6/master.m3u8",
		"https://v.ozone.ru/vod/video-28/01HGM068NMZVZ5F7VSP7R4YJH0/master.m3u8",
		"https://v.ozone.ru/vod/video-34/01HGM04P09D7CJWSP843Z7EYBN/master.m3u8",
		"https://v.ozone.ru/vod/video-29/01HGM01589W4YGSRWSG13W3Q9W/master.m3u8",
		"https://v.ozone.ru/vod/video-31/01HGKZZYE9JTZSF0SH8GRNVZGB/master.m3u8",
		"https://v.ozone.ru/vod/video-29/01HGKZXMEG3XCPXS3V30J3JSBB/master.m3u8",
		"https://v.ozone.ru/vod/video-34/01HGKZPB41FQQGEZCHXX06D8AV/master.m3u8",
		"https://v.ozone.ru/vod/video-29/01HGKZN50KKC59K3C770EAVNJ1/master.m3u8",
		"https://v.ozone.ru/vod/video-31/01HGKZJQK8XSNT5N1Y54GEMN4G/master.m3u8",
	}

	fmt.Printf("Всего ссылок на скачивание: %d\n", len(links))

	DownloadVideo(links)
}

func DownloadVideo(links []string) {
	const workerCount = 3
	tasks := make(chan string, len(links))
	var wg sync.WaitGroup

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for links := range tasks {
				fmt.Printf("[Воркер %d] Начинаем скачивание: %s\n", workerID, links)
				err := utils.RunFFmpeg(links)
				if err != nil {
					fmt.Printf("[Воркер %d] Ошибка при скачивании $s: %v\n", workerID, links)
				} else {
					fmt.Printf("[Воркер %d] Успешно скачанно: %s\n", workerID, links)
				}
			}
		}(i)
	}

	for _, link := range links {
		tasks <- link
	}

	close(tasks)

	wg.Wait()
	fmt.Println("Все загрузки завершены.")
}

// Проверяет, существует ли папка, и создаёт её, если нет
func createOutputDirectory() {
	outputDir := "videos"
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.Mkdir(outputDir, 0755)
		if err != nil {
			fmt.Printf("Ошибка создания директории: %v\n", err)
			os.Exit(1) // Завершаем программу, если не удалось создать папку
		}
		fmt.Println("Папка 'videos' создана.")
	}
}
