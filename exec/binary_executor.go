package exec

import (
        "fmt"
        "os/exec"
)

func ExecBinary(string binaryPath) {

        cmd := exec.Command(binaryPath)
        output, err := cmd.Output()
        if err != nil {
                fmt.Println("Error executing binary:", err)
                return
        }
        fmt.Println("Binary executed successfully. Output:", string(output))
}


