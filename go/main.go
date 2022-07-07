package main

// #cgo CFLAGS: -I../lib
// #cgo LDFLAGS: -L../lib -lrust
// #include<librust.h>
import "C"
import (
	"fmt"
)

//export goGetNonce
func goGetNonce(address *C.char, block_number C.ulonglong) C.ulonglong {
	fmt.Println("call go goGetNonce")
	if address == C.CString("abcdefg") {
		return 1000
	}
	return 10
}

func main() {
	addr := C.CString("abcdefg")
	fmt.Println(C.execute(addr, 1000))
}
