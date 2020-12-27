package main

import (
	"fmt"
	"time"
)

func writeToChannel2(c chan int, x int) {
	fmt.Println("1", x)
	c <- x
	close(c)
	fmt.Println("2", x)
}

func main() {
	c := make(chan int)
	go writeToChannel2(c, 10)
	time.Sleep(1 * time.Second)

	// '<- c'라는 표현을 통해 채널 c에서 데이터를 읽도록 지정했다.
	// 읽은 값을 출력하지 않고 변수 k에 저장하고 싶다면 'k := <- c'와 같이 적으면 된다.
	fmt.Println("Read:", <-c)
	// 채널에서 읽을 시간을 지정했다.
	time.Sleep(1 * time.Second)

	/*
	 * 채널 c가 열렸는지를 확인하는 방법을 제시한다.
	 * 채널이 닫혀 있어도 코드 실행에는 문제가 없다. 하지만 채널이 열려 있다면 채널에서 읽은 값을 버린다.
	 * '_, ok := <- c'에서 '_'라고 적은 부분 때문이다. '_' 대신 적절한 변수를 지정했다면 채널이
	 * 열렸을 때 읽은 값도 볼 수 있다.
	 */
	_, ok := <-c
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}
}
