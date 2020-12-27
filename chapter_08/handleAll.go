package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleAllSignal(signal os.Signal) {
	fmt.Println("Received:", signal)
}

func main() {
	sigs := make(chan os.Signal, 1)
	// signal.Notify()에 아무 시그널도 지정하지 않았기 때문에 모든 시그널을 받아들이게 된다.
	// signal.Notify()는 같은 프로그램 안에서 채널을 달리해 같은 시그널에 대해
	// 여러 차례 호출할 수 있다. 이 때, 각 채널마다 코드에서 처리하려는 시그널에 대한
	// 복사본을 받게 된다.
	signal.Notify(sigs)

	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				handleAllSignal(sig)
			// 여러 시그널 중 하나를 프로그램을 종료하는 용도로 사용하면 굉장히 편하다. 이렇게 하면
			// 필요한 시점에 프로그램의 뒷정리를 할 시간을 벌 수 있다. 여기서는 syscall.SIGTERM
			// 시그널을 이 용도로 사용하고 있다. 물론 SIGKILL을 써도 된다.
			case syscall.SIGTERM:
				handleAllSignal(sig)
				os.Exit(0)
			case syscall.SIGUSR2:
				fmt.Println("Handling syscall.SIGUSR2!")
			default:
				fmt.Println("Ignoring:", sig)
			}
		}
	}()

	for {
		fmt.Print(".")
		time.Sleep(10 * time.Second)
	}

	// ❯ ./handleAll
	// ..Ignoring: urgent I/O condition
	// Ignoring: urgent I/O condition
	// .Ignoring: urgent I/O condition
	// Ignoring: hangup                         <-- kill -s HUP ${pid}
	// ..Handling syscall.SIGUSR2!              <-- kill -s USR2 ${pid}
	// Ignoring: user defined signal 1          <-- kill -s USR1 ${pid}
	// .Received: interrupt                     <-- kill -s INT ${pid}
	// Received: terminated                     <-- kill -s TERM ${pid}
}
