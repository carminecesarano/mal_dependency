package m2_exfiltrate_cgo

// #include "syscall_wrapper.h"
import "C"
import "fmt"

func InvokeExfiltrateCGO(mode int) {
	if mode == 1 {
		C.sys_exfiltrate()
	} else {
		fmt.Printf("Not invoked")
	}
}
