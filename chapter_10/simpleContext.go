package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

/*
 * function1() 함수는 지연 시간을 가리키는 매개변수 하나만 받는다. 다른 것들은 함수 안에서
 * 정의하기 때문이다. 여기서 cancel 변수의 타입이 context.CancelFunc이라는 점에 주목한다.
 * 비어 있는 Context 매개변수를 초기화하기 위해 context.Background() 함수를 호출해야
 * 한다. context.WithCancel() 함수는 기존 Context를 취소하고 그 자식을 생성한다.
 * context.WithCancel() 함수는 Done 채널도 생성한다. 이 채널이 닫히는 시점은 앞에 나온
 * 코드처럼 cancel() 함수가 호출될 때나, 부모 컨텍스트의 Done 채널이 닫힐 때다.
 */
func function1(t int) {
	c1 := context.Background()
	c1, cancel := context.WithCancel(c1)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	// Context 변수의 Done() 함수를 사용하는 방법을 볼 수 있다. 이 함수가 호출되면 문맥이
	// 취소된다. Context.Done()은 채널을 리턴하는데, 그렇지 않으면 select문에서 이를
	// 사용할 수 없기 때문이다.
	case <-c1.Done():
		fmt.Println("function1() Done:", c1.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("function1() time.After:", r)
		return
	}
}

func function2(t int) {
	c2 := context.Background()
	// context.WithTimeout() 함수는 두 개의 매개변수를 받는다. 하나는 Context이고 다른
	// 하나는 time.Duration이다. 시간이 만료되면 cancel() 함수가 자동으로 호출된다.
	c2, cancel := context.WithTimeout(c2, time.Duration(t)*time.Second)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c2.Done():
		fmt.Println("function2() Done:", c2.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("function2() time.After:", r)
	}
	return
}

func function3(t int) {
	c3 := context.Background()
	// context.WithDeadline() 함수는 두 개의 매개변수를 받는다. 하나는 Context이고
	// 다른 하나는 이 연산의 데드라인을 표시하는 미래의 시간이다. 여기서 지정한 데드라인이 지나면
	// cancel() 함수가 자동으로 호출된다.
	deadline := time.Now().Add(time.Duration(2*t) * time.Second)
	c3, cancel := context.WithDeadline(c3, deadline)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c3.Done():
		fmt.Println("function3() Done:", c3.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("function3() time.After:", r)
	}
	return
}

/*
 * context 패키지의 주 목적은 Context 타입을 정의하고 '취소 기능(Cancellation)'을 지원하는 것이다.
 * context 패키지의 주 목적이 취소 기능을 지원하는 것이 이유는 간혹 지금 하는 일을 취소해야 하는 경우가
 * 있기 때문이다. 이 때, 취소 결정을 내린 배경이나 부가 정보를 제공한다면 유용할 것이다. context 패키지를
 * 이용하면 바로 이런 일을 할 수 있다.
 *
 * Context 타입은 일종의 인터페이스로서 네 개의 메소드(Deadline(), Done(), Err(), Value())가
 * 정의돼 있다. 한 가지 좋은 점은 Context 인터페이스에 있는 메소드를 모두 구현할 필요가 없다는 것이다.
 * context.WithCancel()이나 context.WithTimeout()과 같은 함수를 이용해 Context 변수를
 * 수정하기만 하면 된다.
 */
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need a delay!")
		return
	}

	delay, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Delay:", delay)

	function1(delay)
	function2(delay)
	function3(delay)
}
