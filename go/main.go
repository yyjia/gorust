package main

import (
	"fmt"
	"gorust/tools"
)

func main() {
	fmt.Println(tools.GoExec("abcdefg", 1234))
}
