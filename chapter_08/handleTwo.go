package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(signal os.Signal) {
	fmt.Println("handleSignal() Caught:", signal)
}

/*
 * 처리할 시그널은 SIGINFO와 SIGINT이다. Go 언어에서는 이 시그널을
 * syscall.SIGINFO와 os.Interrupt라고 표현한다.
 *
 * os 패키지에 대한 공식 문서를 보면 syscall.SIGKILL과 syscall.SIGINT라는 두 가지
 * 시그널만 모든 시스템에서 확실히 지원한다고 나와 있다. 이 시그널은 Go 언어에서 os.Kill과
 * os.Interrupt라고 정의돼 있다.
 */
func main() {
	// sigs란 이름의 채널(Channel)을 정의한다. 이 채널을 이용해 데이터를 주고 받는다.
	sigs := make(chan os.Signal, 1)
	// signal.Notify()를 호출하여 원하는 시그널을 보낸다.
	signal.Notify(sigs, os.Interrupt, syscall.SIGINFO)

	// Go 루틴으로 실행할 익명 함수(Anonymous Function)를 구현한다. 이 함수는 원하는
	// 시그널이 들어올 때 수행할 동작을 정의한다.
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				fmt.Println("Caught:", sig)
			case syscall.SIGINFO:
				handleSignal(sig)
				return
			}
		}
	}()

	// 프로그램이 끝나지 않도록 time.Sleep()을 호출했다. 예제에서는 실제로 하는 일이 없기
	// 때문이다. 따라서 go run handleTwo.go라고 실행하지 않고 handleTwo.go를 먼저
	// 컴파일한 뒤에 생성된 실행 파일을 따로 구동한다.
	for {
		fmt.Printf(".")
		time.Sleep(20 * time.Second)
	}

	// ❯ go build handleTwo.go
	// ❯ ls -l handleTwo
	// -rwxr-xr-x  1 benjamin  staff  2202584 Dec 27 11:06 handleTwo
	// ❯ ./handleTwo
	// .^CCaught: interrupt
	// ^CCaught: interrupt
	// ..[1]    13459 killed     ./handleTwo
}
