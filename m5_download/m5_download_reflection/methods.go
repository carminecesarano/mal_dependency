package m5_download_reflect

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

type MyType string

func (t MyType) UnsafeMethod() {
	fmt.Printf("Method invoked through reflection\n")
	syscall.Syscall(145, 0, 0, 0)
	cmd := exec.Command("ls")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
