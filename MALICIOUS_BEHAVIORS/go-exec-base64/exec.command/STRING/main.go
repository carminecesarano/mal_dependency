package main

import (
	"fmt"
	"os/exec"
)

func main() {
	
	byteSequence := []byte{101, 99, 104, 111, 32, 72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}

	commandString := string(byteSequence)

	cmd := exec.Command("sh", "-c", commandString)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Errore nell'esecuzione del comando:", err)
	}

	fmt.Println(string(output))
}
