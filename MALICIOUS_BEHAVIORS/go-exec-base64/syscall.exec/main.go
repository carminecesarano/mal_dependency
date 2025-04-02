package main

import (
	"encoding/base64"
	"os"
	"syscall"
	"log"
)

func main() {
	
	encodedCmd := "ZWNobyBIZWxsbyBXb3JsZA=="

	decodedCmd, err := base64.StdEncoding.DecodeString(encodedCmd)
	if err != nil {
		log.Fatalf("Errore durante la decodifica: %v", err)
	}

	err = syscall.Exec("/bin/sh", []string{"/bin/sh", "-c", string(decodedCmd)}, os.Environ())
	if err != nil {
		log.Fatalf("Errore durante l'esecuzione del comando: %v", err)
	}
}

