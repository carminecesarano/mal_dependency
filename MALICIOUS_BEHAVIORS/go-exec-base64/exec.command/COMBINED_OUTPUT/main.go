package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os/exec"
)

func main() {

	encodedCmd := "ZWNobyBIZWxsb1dvcmxkCg=="

	decodedCmd, err := base64.StdEncoding.DecodeString(encodedCmd)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("sh", "-c", string(decodedCmd))
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", output)
}
