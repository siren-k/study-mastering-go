package main

import (
	"fmt"
)

func a() {
	fmt.Println("Inside a()")

	// panic()이 호출될 때마다 실행될 익명 함수 정의
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recover inside a()!")
		}
	}()

	fmt.Println("About to call b()")
	b()
	fmt.Println("b() exited!")
	fmt.Println("Exiting a()")
}

func b() {
	fmt.Println("Inside b()")
	panic("Panic in b()!")
	fmt.Println("Exiting b()")
}

func main() {
	a()
	fmt.Println("main() ended!")

	// ❯ go run panicRecover.go
	// Inside a()
	// About to call b()
	// Inside b()
	// Recover inside a()!
	// main() ended!

	// 원래 의도한 대로 패닉 상태에 빠지지 않고 종료했다.
	// defer문으로 정의한 익명 함수가 다시 제어궈능ㄹ 찾아왔기 때문이다.
	// 여기에서 주목할 점은, a() 함수는 b() 함수의 패닉을 처리하도록 작성됐음에도 불구하고,
	// b() 함수는 a() 함수에 대해 전혀 모르고 있다는 것이다.
}
