package documentMe

import "fmt"

func ExampleS1() {
	fmt.Println(S1("123456789"))
	fmt.Println(S1(""))
	// Output:
	// 9
	// 0
}

func ExampleF1() {
	fmt.Println(F1(10))
	fmt.Println(F1(2))
	// Output:
	// 1
	// 55
}

// ❯ cp -R documentMe ~/.go/src
// ❯ ls ~/.go/src/documentMe
// documentMe.go      documentMe_test.go
// ❯ go install documentMe
// ❯ ls -l ~/.go/pkg/darwin_amd64/documentMe.a
// -rw-r--r--  1 benjamin  staff  3196 Dec 29 17:54 /Users/benjamin/.go/pkg/darwin_amd64/documentMe.a
// ❯ godoc -http=":8080"
// using GOPATH mode

// ❯ go test -v documentMe/documentMe*
// === RUN   ExampleS1
// --- PASS: ExampleS1 (0.00s)
// === RUN   ExampleF1
// --- FAIL: ExampleF1 (0.00s)
// got:
// 20
// 4
// want:
// 1
// 55
// FAIL
// FAIL    command-line-arguments  0.473s
// FAIL
