package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {

	encodedCmd := "L2Jpbi9scw=="

	cmd, err := base64.StdEncoding.DecodeString(encodedCmd)
	if err != nil {
		fmt.Println("Error decoding Base64 command:", err)
	}

	args := []string{}
	env := os.Environ()
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Env:   env,
	}
	proc, err := os.StartProcess(string(cmd), args, attr)
	if err != nil {
		fmt.Println("Errore nell'esecuzione", err)
	}

	_, err = proc.Wait()
	if err != nil {
		fmt.Println("Errore nell'attesa", err)
	}
}
