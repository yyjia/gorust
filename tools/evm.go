package tools

/*
#cgo CFLAGS: -I../lib
#cgo LDFLAGS: -L../lib -lrust
#include<librust.h>
*/
import "C"
import (
	"fmt"
)

//export goGetNonce
func goGetNonce(address *C.char, block_number C.ulonglong) C.ulonglong {
	fmt.Println("call go goGetNonce")
	if "abcdefg" == C.GoString(address) {
		return 1000
	}
	return 10
}

func GoExec(str string, num int) int {
	return int(C.execute(C.CString(str), C.int(num)))
}
