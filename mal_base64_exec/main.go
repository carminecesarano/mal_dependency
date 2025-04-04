package mal_base64_exec

import (
	"encoding/base64"
	"os"
	"os/exec"
)

func Base64Exec() {

	encodedCmd := "ZWNobyBIZWxsbyBXb3JsZA=="
	decodedCmd, _ := base64.StdEncoding.DecodeString(encodedCmd)

	// PATTERN 1
	// syscall.Exec("/bin/sh", []string{"/bin/sh", "-c", string(decodedCmd)}, os.Environ())

	// PATTERN 2
	cmd1 := exec.Command("/bin/sh", "-c", string(decodedCmd))
	cmd1.Stdout = os.Stdout
	cmd1.Start()
	cmd1.Wait()

	// PATTERN 3
	cmd2 := exec.Command("/bin/sh", "-c", string(decodedCmd))
	cmd2.Stdout = os.Stdout
	cmd2.Run()

	// PATTERN 4

	proc, _ := os.StartProcess(
		"/bin/sh",
		[]string{"/bin/sh", "-c", string(decodedCmd)},
		&os.ProcAttr{
			Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
			Env:   os.Environ(),
		},
	)
	proc.Wait()

}
