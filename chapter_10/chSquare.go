package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var times int

func f1(cc chan chan int, f chan bool) {
	/*
	 * 일반적인 형태인 int 채널을 선언한 뒤에, 이를 채널에 대한 채널 변수로 전달한다.
	 * 그러고 나서 일반 int 채널로부터 데이터를 읽거나 시그널 채널인 f를 사용해 함수를
	 * 종료하는 select문을 작성한다.
	 */
	c := make(chan int)
	cc <- c
	defer close(c)

	/*
	 * c 채널로부터 값 하나를 읽었다면 for 루프로 들어가 0부터 읽은 변수 값까지의 모든
	 * 값을 더한다. 그런 다음 계산한 값을 c int 채널로 보내고 리턴한다.
	 */
	sum := 0
	select {
	case x := <-c:
		for i := 0; i <= x; i++ {
			sum = sum + i
		}
		c <- sum
	case <-f:
		return
	}
}

/*
 * '채널에 대한 채널(Channel of Channel)'은 특수한 종류의 채널 변수로서 다른 변수 타입이 아닌
 * 채널을 다룬다. 그렇긴 하지만 채널에 대한 채널이라도 데이터 타입은 여전히 선언해야 한다. 채널에 대한
 * 채널은 다음과 같이 chan 키워드를 한 줄에 두 번 지정하는 방식으로 정의한다.
 * c1 := make(chan chan int)
 *
 * 채널에 대한 채널에 비해 이 장에서 소개하는 다른 타입의 채널이 훨씬 자주 사용하고 쓸모도 많다.
 */
func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Need just one integer argument!")
		return
	}

	times, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// 채널에 대한 채널인 cc라는 변수를 선언했다. 이 채널은 에제에서 가장 핵심적인 부분으로서
	// 모든 작업을 이 변수를 통해 처리한다. cc 변수가 f1() 함수에 전달되면 다음에 나오는
	// for 루프 안에서 이를 사용한다.
	cc := make(chan chan int)
	for i := 1; i < times+1; i++ {
		// f 채널은 '시그널 채널(Signal Channel)'로 실제 작업이 끝날 때 Go 루틴이
		// 종료하는데 사용한다. 시그널 채널의 타입은 마음대로 정할 수 있다. 에제에서는
		// bool로 지정했고, 다음 절에서는 struct{}로 지정한 것을 볼 수 있다. 시그널
		// 채널을 struct{} 타입으로 지정하면 아무런 데이터를 보내지 않아도 되서, 버그가
		// 발생하거나 실수하는 일을 줄일 수 있다.
		f := make(chan bool)
		go f1(cc, f)
		// 채널에 대한 채널 변수로부터 일반 채널을 받아온다.
		ch := <-cc
		// 일반 채널인 ch 채널에 int 값을 보낸다.
		ch <- i
		// for 루프에서 채널에 들어온 값을 읽는다. f1() 함수는 값 하나만 되돌려 받도록
		// 작성했지만 여러 값을 읽을 수도 있다. 이 때, 각 i는 다양한 Go 루틴으로부터 값을
		// 받는다.
		for sum := range ch {
			fmt.Print("Sum(", i, ")=", sum)
		}
		fmt.Println()
		time.Sleep(time.Second)
		close(f)
	}
}
