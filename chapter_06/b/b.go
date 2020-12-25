package b

import (
	"a"
	"fmt"
)

func init() {
	fmt.Println("init() b")
}

func FromB() {
	fmt.Println("fromB()")
	a.FromA()
}

// ❯ mkdir -p /usr/local/Cellar/go/current/libexec/src/b
// ❯ cp -R b /usr/local/Cellar/go/current/libexec/src
// ❯ go install b
// ❯ cd /usr/local/Cellar/go/current/libexec/pkg/darwin_amd64
// ❯ ls -l b.a
// -rw-r--r--  1 benjamin  staff  9514 Dec 26 02:09 b.a
