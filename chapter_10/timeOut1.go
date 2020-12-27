package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	go func() {
		// time.Sleep()을 호출한 이유는 함수가 작업을 끝낼 때까지 통상적으로
		// 걸리는 시간을 표현하기 위해서다. 여기서는 Go 루틴을 실행하는 익명 함수가
		// 대략 3초(time.Second * 3) 이내에 c1 채널에 메시지를 쓴다는 것을
		// 표현했다.
		time.Sleep(time.Second * 3)
		c1 <- "c1 OK"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	// time.After() 함수를 호출한 이유는 지정한 시간만큼 기다리기 위해서다.
	// 여기서는 time.After()에서 리턴하는 실제 값에는 관심 없고, time.After()
	// 함수가 끝났자는 시실, 다시 말해 그만큼 시간이 지났다는 사실만 중요하다. 이 때,
	// time.After() 함수에 지정한 값이 앞에 나온 코드에서 Go 루틴으로 실행할
	// time.Sleep() 호출에 지정한 값보다 작을수록 타임아웃 메시지를 받을 확률이
	// 높아진다.
	case <-time.After(time.Second * 1):
		fmt.Println("timeout c1")
	}

	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 3)
		c2 <- "c2 OK"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	// time.Sleep() 호출로 인해 대략 3초가 걸리는 Go 루틴을 실행하는 동시에,
	// time.After(time.Second * 4)라는 문장을 통해 만료 시간이 4초라고 정의하고
	// 있다. select 블록의 첫 번째 case에서 c2 채널로부터 값을 받는 부분이
	// time.After(time.Second * 4)가 리턴되기 전에 실행되면 타임아웃이 발생하지
	// 않느다. 반면 그렇지 않으면 타임아웃이 발생한다. 그런데 여기서는 time.After()을
	// 호출할 때 time.Sleep()이 리턴할 때까지 시간보다 길게 지정하면 여기서 타임아웃
	// 메시지를 받을 확률이 낮아진다.
	case <-time.After(time.Second * 4):
		fmt.Println("timeout c2")
	}

	// 첫 번째 Go 루틴이 정상적으로 끝나지 않은 반면, 두 번째 Go 루틴은 시간이 충분하여
	// 정상적으로 종료했다.
	// ❯ go run timeOut1.go
	// timeout c1
	// c2 OK
}
