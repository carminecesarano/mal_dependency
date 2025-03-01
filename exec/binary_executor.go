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
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("Error:", err)
		}

	case "2":
		if err := cmd.Start(); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Command started successfully.")
		}

	case "3":
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Output:", string(output))
		}

	case "4":
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println("Output:", string(output))

	case "5":
		binary, err := exec.LookPath(binaryPath)
		if err != nil {
			fmt.Println("Error finding binary:", err)
			return
		}
		env := os.Environ()
		fmt.Println("Replacing process with", binary)
		if err := syscall.Exec(binary, []string{binary}, env); err != nil {
			fmt.Println("Error executing syscall.Exec:", err)
		}

	default:
		fmt.Println("Invalid choice.")
	}
}
