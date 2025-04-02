package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os/exec"
)

func main() {

	encodedCmd := "ZWNobyBIZWxsb1dvcmxkCg=="

	decodedCmd, err_exec := base64.StdEncoding.DecodeString(encodedCmd)
	if err_exec != nil {
		log.Fatal(err_exec)
	}

	out, err_exec := exec.Command("sh", "-c", string(decodedCmd)).Output()
	if err_exec != nil {
		log.Fatal(err_exec)
	}

	fmt.Printf("%s\n", out)
}
