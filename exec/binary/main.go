package main

import (
        "os"
        "os/exec"
)

func main() {
        cmd := exec.Command("ls")
        cmd.Stdout = os.Stdout
        _ = cmd.Run()
}





