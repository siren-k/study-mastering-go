package main

import (
	"fmt"
	"os"
	"testing"
)

var ERR error

func benchmark_Create(b *testing.B, buffer, filesize int) {
	var err error
	for i := 0; i < b.N; i++ {
		err = Create("random", buffer, filesize)
	}
	ERR = err
	err = os.Remove("random")
	if err != nil {
		fmt.Println(err)
	}
}

func Benchmark_1_Create(b *testing.B) {
	benchmark_Create(b, 1, 1000000)
}

func Benchmark_2_Create(b *testing.B) {
	benchmark_Create(b, 2, 1000000)
}

func Benchmark_4_Create(b *testing.B) {
	benchmark_Create(b, 4, 1000000)
}

func Benchmark_10_Create(b *testing.B) {
	benchmark_Create(b, 10, 1000000)
}

func Benchmark_1000_Create(b *testing.B) {
	benchmark_Create(b, 1000, 1000000)
}

// 쓰기 버퍼의 크기가 1바이트이면 프로그램 전체의 실행 속도를 떨어뜨릴 정도로 굉장히 비효율적이라는 사실은 명백히 알 수 있다.
// 또한, 버퍼 크기를 이렇게 작게 설정하면 메모리 연산이 너무나 많이 발생해서 프로그램 속도가 더욱 느려진다. 쓰기 버퍼의 크기를
// 2바이트로 지정하면 프로그램 속도가 이전보다 두 배나 빨라진다. 그래도 여진히 느린 편이다.
// 쓰기 버퍼의 크기를 4바이트로 지정할 때부터 속도가 전반적으로 높아진다. 그리고 쓰기 버퍼의 크기를 1000바이트로 지정하더라도
// 100바이트로 지정할 때에 비해 100배나 빨라지지 않는다. 따라서 최상의 성능을 발휘할 수 있는 버퍼 크기는 4바이트와 1000바이트
// 사이라는 점을 알 수 있다.

// ❯ go test -bench=. writeBU.go writeBU_test.go
// goos: darwin
// goarch: amd64
// Benchmark_1_Create-12                  1        2579924506 ns/op
// Benchmark_2_Create-12                  1        1280134758 ns/op
// Benchmark_4_Create-12             238256              4212 ns/op
// Benchmark_10_Create-12            498114              2139 ns/op
// Benchmark_1000_Create-12          718864              1667 ns/op
// PASS
// ok      command-line-arguments  26.303s

// ❯ go test -bench=. writeBU.go writeBU_test.go -benchmem
// goos: darwin
// goarch: amd64
// Benchmark_1_Create-12                  1        2510796987 ns/op        16002408 B/op    2000018 allocs/op
// Benchmark_2_Create-12                  1        1434416527 ns/op         8000520 B/op    1000012 allocs/op
// Benchmark_4_Create-12             263412              4083 ns/op             295 B/op          6 allocs/op
// Benchmark_10_Create-12            538075              2150 ns/op             289 B/op          5 allocs/op
// Benchmark_1000_Create-12          730364              1652 ns/op             284 B/op          5 allocs/op
// PASS
// ok      command-line-arguments  26.679s
