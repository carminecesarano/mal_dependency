package syscall

// #include "syscall_wrapper.h"
import "C"
import "fmt"

func InvokeSyscall(syscallNo int) {

	switch syscallNo {
	case 0:
		C.sys_read()
		// C.sys_read_clib()
	case 1:
		C.sys_write()
		// sys_write_assembly
	case 145:
		C.sys_sched_getscheduler()
	case 170:
		C.sys_sethostname()
	case 204:
		C.sys_sched_getaffinity()
	case 252:
		C.sys_ioprio_get()
	}

}

func InvokeTest() {
	fmt.Printf("test")
}
