package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("You are using ", runtime.Compiler)
	fmt.Println("on a", runtime.GOARCH, "machine")
	fmt.Println("with Go version", runtime.Version())
}

// ❯ env GOOS=linux GOARCH=arm go build xCompile.go
// ❯ file xCompile
// xCompile: ELF 32-bit LSB executable, ARM, EABI5 version 1 (SYSV), statically linked, Go BuildID=h0DIUnpa8SzERkAgDco-/8Cp73-CgFKnHRul9QWao/Evjz1_9srjJa6uUnDErk/qJen2nkNblkSuflMPsdS, not stripped
// ❯ ./xCompile
// zsh: exec format error: ./xCompile
