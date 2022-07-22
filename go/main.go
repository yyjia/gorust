package main

/*
#cgo CFLAGS: -I../lib
#cgo LDFLAGS: -L../lib -lrust
#include<librust.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export goGetNonce
func goGetNonce(address *C.char, block_number C.ulonglong) C.ulonglong {
	fmt.Println("call go goGetNonce")
	if "abcdefg" == C.GoString(address) {
		return 1000
	}
	return 10
}

func main() {
	addr := C.CString("abcdefg")
	defer C.free(unsafe.Pointer(addr))
	if 1000 != C.execute(addr, 1234) {
		fmt.Println("faild")
	}

	addr1 := C.CString("xxxxxxx")
	defer C.free(unsafe.Pointer(addr1))
	if 10 != C.execute(addr1, 1234) {
		fmt.Println("failed")
	}
}
