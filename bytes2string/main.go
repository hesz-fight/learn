package main

import (
	"fmt"
	"unsafe"
)

func main() {
	chs := []byte{65, 66}
	str := *(*string)(unsafe.Pointer(&chs))
	fmt.Println(str)
}
