package main

import "C"

import (
	"fmt"
)

//export PrintMessage
func PrintMessage() {
	fmt.Println("A Go function!")
}

//export Multiply
func Multiply(a, b int) int {
	return a * b
}

func main() {
}

// ❯ go build -o usedByC.o -buildmode=c-shared usedByC.go
// ❯ ls -l usedByC.*
// -rw-r--r--  1 benjamin  staff      273 Dec 21 07:07 usedByC.go
// -rw-r--r--  1 benjamin  staff     1635 Dec 21 07:07 usedByC.h
// -rw-r--r--  1 benjamin  staff  2624652 Dec 21 07:07 usedByC.o
// ❯ file usedByC.o
// usedByC.o: Mach-O 64-bit dynamically linked shared library x86_64
