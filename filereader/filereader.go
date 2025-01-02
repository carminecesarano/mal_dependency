package filereader

// #include "hello.h"
import "C"
import "time"

func ReadFile() {
	time.Sleep(2 * time.Second)
	C.hello()
}
