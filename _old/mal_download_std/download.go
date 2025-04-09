package mal_download_std

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile() {
	url := "http://localhost:8080/executable"
	fileName := "downloaded_file"

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	out, _ := os.Create(fileName)
	defer out.Close()

	_, _ = io.Copy(out, resp.Body)
	_ = out.Close()
	_ = os.Chmod(fileName, 0755)
}
