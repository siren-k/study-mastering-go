package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func timeout(w *sync.WaitGroup, t time.Duration) bool {
	temp := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		defer close(temp)
		w.Wait()
	}()

	select {
	case <-temp:
		return false
	// time.After() 함수를 호출할 때 지정할 시간을 timeout() 함수의 매개변수로
	// 지정했다. 따라서 구체적인 값이 달라질 수 있다. 여기서도 마찬가지로 select 블록으로
	// 타임아웃 로직을 구현했다. 또한 sync.WaitGroup.Wait() 함수를 호출하면 매칭되는
	// sync.WaitGroup.Done() 함수가 호출될 때까지 timeout() 함수는 무한정 기다린다.
	// sync.WaitGroup.Wait() 함수가 리턴하면 select문의 첫 번째 브랜치가 실행된다.
	case <-time.After(t):
		return true
	}
}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Need a time duration")
		return
	}

	var w sync.WaitGroup
	w.Add(1)

	t, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Print(err)
		return
	}

	/*
	 * time.Duration() 함수는 정수값을 time.Duration 변수로 변환한다.
	 * 이 값은 나중에 사용한다.
	 */
	duration := time.Duration(int32(t)) * time.Millisecond
	fmt.Printf("Timeout period is %s\n", duration)

	if timeout(&w, duration) {
		fmt.Println("Timed out!")
	} else {
		fmt.Println("OK!")
	}

	// sync.WaitGroup.Done() 함수가 실행된 뒤에는 앞에서 나온 timeout() 함수가
	// 리턴한다. 하지만 두 번째로 호출한 timeout()에 대해서는 기다릴
	// sync.WaitGroup.Done() 함수가 없다.
	w.Done()
	if timeout(&w, duration) {
		fmt.Println("Timed out!")
	} else {
		fmt.Println("OK!")
	}

	// ❯ go run timeOut2.go 10000
	// Timeout period is 10s
	// Timed out!
	// OK!
	// ==> 타임아웃 주기를 익명 Go 루틴에 대한 time.Sleep(5 * time.Second) 함수
	//     호출보다 길지 않지 지정했다. 하지만 이에 대해 sync.WaitGroup.Done()을
	//     호출하지 않았기 때문에 익명 Go 루틴은 리턴할 수 없으며, 따라서 time.After(t)가
	//     먼저 끝난다. 그래서 첫 번쨰 if문의 timeout() 함수는 true를 리턴한다.
	//     두 번째 if문에서 익명 함수는 더 이상 기다릴 것이 없기 때문에 timeout() 함수는
	//     false를 리턴한다. time.Sleep(5 * time.Second)가 time.After(t)보다
	//     먼저 끝나기 때문이다.
}
