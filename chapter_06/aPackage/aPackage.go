package aPackage

import (
	"fmt"
)

func A() {
	fmt.Println("This is function A!")
}

func B() {
	fmt.Println("privateConstant:", privateConstant)
}

const MyConstant = 123
const privateConstant = 21

// ❯ mkdir -p /usr/local/Cellar/go/current/libexec/src/aPackage
// ❯ cp -R aPackage /usr/local/Cellar/go/current/libexec/src
// ❯ go install aPackage
// ❯ cd /usr/local/Cellar/go/current/libexec/pkg/darwin_amd64
// ❯ ls -l aPackage.a
// -rw-r--r--  1 benjamin  staff  20864 Dec 26 01:50 aPackage.a

/*
 * Go 패키지를 main() 함수에서 불러오지 않고서는 직접 실행할 수 없지만, 패키지를 컴파일해서
 * 오브젝트 파일로 만들 수는 있다.
 * ❯ go tool compile aPackage.go
 * ❯ ls
 * aPackage.go aPackage.o
 * ❯ ll
 * total 48
 * -rw-r--r--  1 benjamin  staff   736B Dec 26 01:53 aPackage.go
 * -rw-r--r--  1 benjamin  staff    18K Dec 26 01:53 aPackage.o
 */
