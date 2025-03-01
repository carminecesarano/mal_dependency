package exec

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func ExecBinary(binaryPath string, choice string) {
	cmd := exec.Command(binaryPath)

	switch choice {
	case "1":
		fmt.Printf("Binary Executor 1\n")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("Error:", err)
		}

	case "2":
		fmt.Printf("Binary Executor 2\n")
		if err := cmd.Start(); err != nil {
			fmt.Println("Error:", err)
		}

	case "3":
		fmt.Printf("Binary Executor 1\n")
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Output:", string(output))
		}

	case "4":
		fmt.Printf("Binary Executor 1\n")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println("Output:", string(output))

	case "5":
		fmt.Printf("Binary Executor 1\n")
		binary, err := exec.LookPath(binaryPath)
		if err != nil {
			fmt.Println("Error finding binary:", err)
			return
		}
		env := os.Environ()
		if err := syscall.Exec(binary, []string{binary}, env); err != nil {
			fmt.Println("Error executing syscall.Exec:", err)
		}

	default:
		fmt.Println("Invalid choice.")
	}
}
