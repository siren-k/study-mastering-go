package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, "This is Standard output\n")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stderr, "\n")

	// $ go run stdERR.go
	// This is Standard output
	// Please give me one argument!

	// 표준 에러를 특정한 파일로 리다이렉션
	// $ go run stdERR.go 2>/tmp/stdErr
	// This is Standard output
	// cat /tmp/stdErr
	// Please give me one argument!

	// 에러 출력을 무시
	// $ go run stdERR.go 2>/dev/null
	// This is Standard output

	// 표준 출력과 표준 에러를 동일한 파일로 리다이렉션
	// $ go run stdERR.go >/tmp/output 2>&1
	// cat /tmp/output
	// This is Standard output
	// Please give me one argument!

	// 표준 출력과 표준 에러를 모두 무시
	// $ go run stdERR.go >/dev/null 2>&1
}
