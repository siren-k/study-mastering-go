package main

import (
	"fmt"
	"time"
)

/*
 * A() 함수는 매개변수 a에 저장된 채널에 의해 블록된다. main() 함수에서 이 채널의 블록
 * 상태가 풀리면 A() 함수가 실행되기 시작한다. 마지막으로 b 채널을 닫는데, 이렇게 하면
 * 다른 함수(여기서는 B() 함수)의 블록 상태가 풀린다.
 */
func A(a, b chan struct{}) {
	<-a
	fmt.Println("A()!")
	time.Sleep(time.Second)
	close(b)
}

/*
 * 함수 B()의 로직은 함수 A()와 같다. 이 함수는 채널 a가 닫힐 떄까지 블로된다. 이 채널이
 * 닫히면 작업을 수행한 뒤 b 채널을 닫는다. 여기서 a와 b 채널은 모두 이 함수의 매개변수
 * 이름을 가리킨다.
 */
func B(a, b chan struct{}) {
	<-a
	fmt.Println("B()!")
	close(b)
}

/*
 * C() 함수는 a 채널이 닫힐 떄까지 블록되어 있다가 패널이 닫히면 실행을 시작한다.
 */
func C(a chan struct{}) {
	<-a
	fmt.Println("C()!")
}

func main() {
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})

	go C(z)
	go A(x, y)
	go C(z)
	go B(y, z)
	go C(z)

	close(x)
	time.Sleep(3 * time.Second)

	// C() 함수를 Go 루틴처럼 여러 번 호출해도 정상적으로 실행된다. C() 함수에서 닫는 채널이 없기 때문이다.
	// ❯ go run defineOrder.go
	// A()!
	// B()!
	// C()!
	// C()!
	// C()!

	// 하지만 A()와 B()를 여러 번 호출하면 다음과 같이 에러 메시지가 발생한다.
	// ❯ go run defineOrder.go
	// A()!
	// A()!
	// B()!
	// C()!
	// C()!
	// C()!
	// panic: close of closed channel
	//
	// goroutine 8 [running]:
	// main.A(0xc000068060, 0xc0000680c0)
	//         /Users/benjamin/Lab/github/siren-k/study-mastering-go/chapter_10/defineOrder.go:17 +0xab
	// created by main.main
	//         /Users/benjamin/Lab/github/siren-k/study-mastering-go/chapter_10/defineOrder.go:46 +0xff
	// exit status 2
}
