package exec

import (
        "fmt"
        "os/exec"
)

func ExecBinary(binaryPath string) {

        cmd := exec.Command(binaryPath)
        output, err := cmd.Output()
        if err != nil {
                fmt.Println("Error executing binary:", err)
                return
        }
        fmt.Println("Binary executed successfully. Output:", string(output))
}


