package main

// #cgo CFLAGS: -I${SRCDIR}/callClib
// #cgo LDFLAGS: ${SRCDIR}/callC.a
// #include <stdlib.h>
// #include <callC.h>
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Going to call a C function!")
	C.Hello()

	fmt.Println("Going to call another C function!")
	myMessage := C.CString("This is Mihalis!")
	defer C.free(unsafe.Pointer(myMessage))
	C.printMessage(myMessage)

	fmt.Println("All perfectly done!")

	// ❯ ls -l callClib
	// total 16
	// -rw-r--r--  1 benjamin  staff  165 Dec 21 06:25 callC.c
	// -rw-r--r--  1 benjamin  staff   87 Dec 21 06:23 callC.h
	// ❯ gcc -c callClib/*.c
	// ❯ ls -l callC.o
	// -rw-r--r--  1 benjamin  staff  952 Dec 21 06:35 callC.o
	// ❯ file callC.o
	// callC.o: Mach-O 64-bit object x86_64
	// ❯ ar rs callC.a *.o
	// ar: creating archive callC.a
	// ❯ ls -l callC.a
	// -rw-r--r--  1 benjamin  staff  1160 Dec 21 06:35 callC.a
	// ❯ file callC.a
	// callC.a: current ar archive random library
	// ❯ rm callC.o
	// ❯ go build callC.go
	// ❯ ls -l callC
	// -rwxr-xr-x  1 benjamin  staff  2610440 Dec 21 06:46 callC
	// ❯ file callC
	// callC: Mach-O 64-bit executable x86_64
	// ❯ ./callC
	// Going to call a C function!
	// Hello from C!
	// Going to call another C function!
	// Go send me This is Mihalis!
	// All perfectly done!
}
