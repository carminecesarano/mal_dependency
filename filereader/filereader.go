package filereader

// #include "hello.h"
import (
	"C"
	"fmt"
	"os"
	"time"
)

func ReadFile() {

	filePath := "/etc/passwd"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: failed to open file: %v\n", err)
		return
	}
	defer file.Close() // Ensure the file is closed

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error: failed to read file: %v\n", err)
		return
	}

	fmt.Printf("File Content:\n%s\n", string(content))

	time.Sleep(2 * time.Second)
	C.hello()
}
