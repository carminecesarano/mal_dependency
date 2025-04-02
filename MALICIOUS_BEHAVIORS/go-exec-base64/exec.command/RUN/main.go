package main

import (
	"encoding/base64"
	"log"
	"os/exec"
)

func main() {

	encodedCmd := "ZWNobyBIZWxsb1dvcmxkID4gaGVsbG8udHh0Cg=="

	decodedCmd, err := base64.StdEncoding.DecodeString(encodedCmd)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("sh", "-c", string(decodedCmd))
	log.Printf("codice eseguito, attesa della fine...")

	err = cmd.Run()

	log.Printf("codice eseguito con errore: %v", err)
}
