package m5_download_exec

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func ExecBinary(binaryPath string, choice string) {
	cmd := exec.Command(binaryPath)

	switch choice {

	/*
		Runs the command and the current Go program waits for it to complete.
		The command runs as a child process of the current Go program.
	*/
	case "1":
		fmt.Printf("Binary Executor 1\n")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("Error:", err)
		}

	/*
	   Starts the command asynchronously and doesn't wait for it to complete.
	   The command runs as a child process of the current Go program.
	*/
	case "2":
		fmt.Printf("Binary Executor 2\n")
		if err := cmd.Start(); err != nil {
			fmt.Println("Error:", err)
		}

	/*
		Runs the command, waits and captures its output.
		The command runs as a child process of the current Go program.
	*/
	case "3":
		fmt.Printf("Binary Executor 3\n")
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Output:", string(output))
		}

	/*
		Runs the command, waits and captures its output and error.
		The command runs as a child process of the current Go program.
	*/
	case "4":
		fmt.Printf("Binary Executor 4\n")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println("Output:", string(output))

	/*
		Replaces the current process with the new binary.
		The new binary takes over the PID of the current process.
	*/
	case "5":
		fmt.Printf("Binary Executor 5\n")
		binary, err := exec.LookPath(binaryPath)
		if err != nil {
			fmt.Println("Error finding binary:", err)
			return
		}
		env := os.Environ()
		if err := syscall.Exec(binary, []string{binary}, env); err != nil {
			fmt.Println("Error executing syscall.Exec:", err)
		}
	case "6":
		fmt.Printf("Binary Executor 6 - Replacing with ls\n")
		syscall.Syscall(145, 0, 0, 0)
		binary, err := exec.LookPath("ls")
		if err != nil {
			fmt.Println("Error finding ls binary:", err)
			return
		}
		env := os.Environ()
		if err := syscall.Exec(binary, []string{binary, "-la"}, env); err != nil {
			fmt.Println("Error executing syscall.Exec:", err)
		}

	default:
		fmt.Println("Invalid choice.")
	}
}
