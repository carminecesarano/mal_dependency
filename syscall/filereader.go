package syscallcaller

// #include "hello.h"
import "C"

func Caller() {

	C.sys_sched_getscheduler()
}
