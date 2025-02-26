package reflection

import (
	"fmt"
	"os"
	"os/exec"
)

type MyType string

func (t MyType) UnsafeMethod() {
	fmt.Printf("Malicious method invoked\n")

	cmd := exec.Command("ls")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
	fmt.Println("lib.init()")
}
