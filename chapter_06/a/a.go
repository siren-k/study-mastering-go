package a

import (
	"fmt"
)

func init() {
	fmt.Println("init() a")
}

func FromA() {
	fmt.Println("fromA()")
}

// ❯ mkdir -p /usr/local/Cellar/go/current/libexec/src/a
// ❯ cp -R a /usr/local/Cellar/go/current/libexec/src
// ❯ go install a
// ❯ cd /usr/local/Cellar/go/current/libexec/pkg/darwin_amd64
// ❯ ls -l a.a
// -rw-r--r--  1 benjamin  staff  17992 Dec 26 02:09 a.a
