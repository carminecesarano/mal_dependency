package filereader

// #include "hello.h"
import (
	"C"
	"time"
	"io/ioutils"
	"os"
)

func ReadFile() {

	filePath := "/etc/passwd"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: failed to open file: %v\n", err)
		return
	}
	defer file.Close() // Ensure the file is closed

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Error: failed to read file: %v\n", err)
		return
	}

	fmt.Printf("File Content:\n%s\n", string(content))

	time.Sleep(2 * time.Second)
	C.hello()
}
