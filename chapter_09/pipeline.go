package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var CLOSEA = false
var DATA = make(map[int]bool)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

/*
 * 생성된 난수를 out 채널에 쓰기
 */
func first(min, max int, out chan<- int) {
	for {
		// second() 함수가 first() 함수에게 첫 번째 채널을 닫으라고 알려주기
		// 위한 수단으로, CLOSEA란 전역 변수를 선언했다. first() 함수는 CLOSEA
		// 변수를 읽을 수만 있고 값을 변경하는 것은 second() 함수만 할 수 있다.
		if CLOSEA {
			close(out)
			return
		}
		out <- random(min, max)
	}
}

/*
 * in 채널로부터 데이터를 받아서 out 채널로 전달
 */
func second(out chan<- int, in <-chan int) {
	for x := range in {
		fmt.Print(x, " ")
		// 이미 맵에 존해하는 데이터를 발견하면 전역 변수 CLOSEA의 값을
		// true로 변경하고 out 채널을 닫는다.
		_, ok := DATA[x]
		if ok {
			CLOSEA = true
		} else {
			DATA[x] = true
			out <- x
		}
	}
	fmt.Println()
	close(out)
}

/*
 * in 채널에서 값을 읽어 화면에 출력
 */
func third(in <-chan int) {
	var sum int
	sum = 0
	/*
	 * second() 함수에 의해 in 채널이 닫히면 sum 변수의 값을 출력한다.
	 */
	for x2 := range in {
		sum = sum + x2
	}
	fmt.Printf("The sum of the random numbers is %d\n", sum)
}

/*
 * 파이프라인(Pipeline)이란 Go 루틴과 채널을 연결하는 기법으로 채널로 데이터를 전송하는 방식으로
 * 한쪽 Go 루틴의 출력을 다른 Go 루틴의 입력으로 연결할 수 있다.
 * 파이프라인을 사용함으로써 얻을 수 있는 장점
 * 1) 하나는 데이터 흐름을 일정하게 구현할 수 있다는 점이다. Go 루틴이나 채널은 다른 작업이 끝날
 *    때까지 기다릴 필요 없이 실행을 시작할 수 있기 때문이다.
 * 2) 주고 받는 값을 일일이 변수에 저장할 필요가 없기 때문에 변수 사용 횟수를 줄일 수 있고, 결과적으로
 *    메모리 공간도 절약할 수 있다.
 * 3) 파이프라인을 사용하면 프로그램 설계를 간결하게 구성하여 유지보수성이 높아진다는 장점도 있다.
 */
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Need two integer parameters!")
		os.Exit(1)
	}

	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])

	if n1 > n2 {
		fmt.Printf("%d should be smaller than %d\n", n1, n2)
		return
	}

	rand.Seed(time.Now().UnixNano())
	/*
	 * 채널 A는 first() 함수로부터 난수를 가져오고 이를 second() 함수로 보내는 용도이다.
	 * 채널 B는 second() 함수에서 적절한 난수를 third() 함수로 보내는 용도이다.
	 */
	A := make(chan int)
	B := make(chan int)

	go first(n1, n2, A)
	go second(B, A)
	third(B)
}
