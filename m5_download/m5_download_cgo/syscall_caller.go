package m5_download_cgo

// #include "syscall_wrapper.h"
import "C"

func InvokeDownloadCGO() {

	C.sys_download()

}
