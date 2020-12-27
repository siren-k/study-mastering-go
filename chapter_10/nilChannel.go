package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
 * add() 함수는 닐 채널을 이용하여 구현하고 있다. <- t.C 케이스는 time.NewTimer() 함수를
 * 호출할 때 지정한 시간만큼 t 타이머의 C 채널을 블록한다. 여기서 이 함수의 매개변수인 c 채널과
 * 헷갈리면 안 된다. t.C 채널은 타이머 t에 속해 있다. 지정된 시간이 만료되면 타이머는 t.C 채널로
 * 값을 보낸다. 그러면 select문에서 이와 관련된 브랜치가 실행되면서 c 채널에 nil 값을 할당한 뒤
 * sum 변수를 화면에 출력한다.
 */
func add(c chan int) {
	sum := 0
	t := time.NewTimer(time.Second)

	for {
		select {
		case input := <-c:
			sum = sum + input
		case <-t.C:
			c = nil
			fmt.Println(sum)
		}
	}
}

/*
 * send() 함수는 채널이 열려 있는 동안 지속적으로 난수를 생성해서 채널로 보낸다.
 */
func send(c chan int) {
	for {
		c <- rand.Intn(10)
	}
}

/*
 * 닐 채널(Nil Channel)은 항상 블록되기 때문에 특수한 종류의 채널로 분류한다.
 */
func main() {
	c := make(chan int)
	go add(c)
	go send(c)

	time.Sleep(3 * time.Second)
}
