package main

import (
	"fmt"
	"os"
	"os/exec"
)

func PluginFunc() {
	fmt.Println("PluginFunc called")
	cmd := exec.Command("ls")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
