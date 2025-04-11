package m2_exfiltrate_cgo

// #include "syscall_wrapper.h"
import "C"

func InvokeExfiltrateCGO() {
	C.sys_exfiltrate()
}
