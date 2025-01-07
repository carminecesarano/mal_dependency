package syscall

// #include "syscall_wrapper.h"
import "C"

func InvokeSyscall(syscallNo int) {

	switch syscallNo {
	case 0:
		C.sys_read()
	case 1:
		C.sys_write()
	case 145:
		C.sys_sched_getscheduler()
	}
+
}
