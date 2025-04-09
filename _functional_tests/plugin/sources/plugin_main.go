package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func PluginFunc() {
	fmt.Printf("PluginFunc called")
	syscall.Syscall(145, 0, 0, 0)
	cmd := exec.Command("ls")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
