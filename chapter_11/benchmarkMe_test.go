package main

import "testing"

var result int

func benchmark_fibo_1(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibo_1(n)
	}
	result = r
}

func benchmark_fibo_2(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibo_2(n)
	}
	result = r
}

func benchmark_fibo_3(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibo_3(n)
	}
	result = r
}

func Benchmark30_fibo_1(b *testing.B) {
	benchmark_fibo_1(b, 30)
}

func Benchmark30_fibo_2(b *testing.B) {
	benchmark_fibo_2(b, 30)
}

func Benchmark30_fibo_3(b *testing.B) {
	benchmark_fibo_3(b, 30)
}

func Benchmark50_fibo_1(b *testing.B) {
	benchmark_fibo_1(b, 50)
}

func Benchmark50_fibo_2(b *testing.B) {
	benchmark_fibo_2(b, 50)
}

func Benchmark50_fibo_3(b *testing.B) {
	benchmark_fibo_3(b, 50)
}

// 여기서 두 가지 사실이 있다. 하나는 실행할 벤치마크 함수를 -bench 매개변수의 값으로 지정했다는
// 것이다. 여기서 '지정한'이란 일종의 정규표현식으로 모든 정상적인 벤치마크 함수를 뜻한다. 다른
// 하나는 -bench 매개변수를 생략하면 아무런 벱치마크 함수도 실핼되지 않는다.
//
// Benchmark30_fibo_1-12                284           4257632 ns/op
// Benchmark30_fibo_1-12                284           4251834 ns/op               0 B/op          0 allocs/op
//                   -12 ==> 실행되는 동안 사용된 Go 루틴의 개수
//                                      284 ==> 함수가 실행된 횟수로 실행 속도가 빠른 함수의 실행 횟수가 느린 함수보다 많다.
//                                                    4251834 ns/op ==> 평균 실행 시간
//                                                                                0 B/op ==> 각 벤치마크 함수를 실행할 때마다 할당된 평균 메모리 양
//                                                                                                0 allocs/op ==> 네 번째 열의 메모리 값을 할당한 횟수
//
// ❯ go test -bench=. benchmarkMe.go benchmarkMe_test.go
// goos: darwin
// goarch: amd64
// Benchmark30_fibo_1-12                284           4257632 ns/op
// Benchmark30_fibo_2-12                255           4512539 ns/op
// Benchmark30_fibo_3-12             371082              2883 ns/op
// Benchmark50_fibo_1-12                  1        63542650785 ns/op
// Benchmark50_fibo_2-12                  1        66453408935 ns/op
// Benchmark50_fibo_3-12             293539              3820 ns/op
// PASS
// ok      command-line-arguments  136.004s
//
// fibo_1()과 fibo_2() 함수는 원래 사용하려던 것 외에 다른 특별한 메모리를 사용하지 않아도 된다.
// 둘 다 자료 구조를 사용하지 않기 때문이다. 반면, fibo_3() 함수는 맵 변수를 사용하기 때문에
// Benchmark30_fibo_3-12의 실행 결과에서 네 번째와 다섯 번째 열의 값이 모두 0보다 크게 나왔다.
//
// ❯ go test -benchmem -bench=. benchmarkMe.go benchmarkMe_test.go
// goos: darwin
// goarch: amd64
// Benchmark30_fibo_1-12                284           4251834 ns/op               0 B/op          0 allocs/op
// Benchmark30_fibo_2-12                274           4441386 ns/op               0 B/op          0 allocs/op
// Benchmark30_fibo_3-12             365750              2953 ns/op            2236 B/op          6 allocs/op
// Benchmark50_fibo_1-12                  1        63915435334 ns/op              0 B/op          0 allocs/op
// Benchmark50_fibo_2-12                  1        67242305044 ns/op              0 B/op          0 allocs/op
// Benchmark50_fibo_3-12             267098              4248 ns/op            2481 B/op         10 allocs/op
// PASS
// ok      command-line-arguments  139.830s
