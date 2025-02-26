package plugin

import (
	"os"
	"os/exec"
)

func PluginFunc() {
	cmd := exec.Command("ls")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
