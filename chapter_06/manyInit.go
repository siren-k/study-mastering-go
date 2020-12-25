package main

import (
	"a"
	"b"
	"fmt"
)

func init() {
	fmt.Println("init() manyInit")
}

func main() {
	a.FromA()
	b.FromB()
	// ❯ go run manyInit.go
	// init() a
	// init() b
	// init() manyInit
	// fromA()
	// fromB()
	// fromA()
}
