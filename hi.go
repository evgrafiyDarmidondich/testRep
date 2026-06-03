package main

import (
	"fmt"
	"unsafe"
)

const name = "Go"

func main() {
	// Просто комент
	var num1 int64 = 15
	var num2 int = 15

	fmt.Println(num1 + int64(num2))
	fmt.Println(unsafe.Sizeof(uint8(1)))
	fmt.Println(unsafe.Sizeof(uint64(1)))
}
