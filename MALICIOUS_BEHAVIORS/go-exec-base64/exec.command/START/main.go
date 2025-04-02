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

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Attesa fine...")

	err = cmd.Wait()

	log.Printf("Codice eseguito e terminato con errore: %v", err)
}
