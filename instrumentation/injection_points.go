package instrumentation

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func LogInjectionPoint() {
	pc, file, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	injectionID := fmt.Sprintf("%s:%d:%s", file, line, fn.Name())
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf("[%s] Injection triggered: %s", timestamp, injectionID)

	// Open the file
	f, err := os.OpenFile("/home/carmine/projects/workspace_goleash/injection_points.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer f.Close()

	// Lock the file
	if err := unix.Flock(int(f.Fd()), unix.LOCK_EX); err != nil {
		fmt.Printf("Error locking file: %v\n", err)
		return
	}
	defer unix.Flock(int(f.Fd()), unix.LOCK_UN)

	// Write to the file
	fmt.Fprintln(f, logMessage)
}
