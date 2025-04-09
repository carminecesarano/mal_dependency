package m5_download_cgo

// #include "syscall_wrapper.h"
import "C"

func InvokeDownloadCGO(mode int) {

	if mode == 1 {
		C.sys_download()
	} else {
		C.sys_download_lib()
	}

}
