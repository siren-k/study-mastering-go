package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

// readValueChannel 채널은 난수를 읽는데 사용
var readValueChannel = make(chan int)

// writeValueChannel 채널은 새로운 난수를 가져오는데 사용
var writeValueChannel = make(chan int)

/*
 * setValue() 함수는 공유 변수의 값을 설정한다.
 */
func setValue(newValue int) {
	writeValueChannel <- newValue
}

/*
 * readValue() 함수는 저장된 변수에서 값을 읽는다.
 */
func readValue() int {
	return <-readValueChannel
}

/*
 * 이 프로그램의 핵심 로직은 모두 monitorValue() 함수에서 구현한다. 그 중에서도
 * 특히 select문을 통해 전반적인 연산을 제어한다. 읽기 요청에 대해서는 read() 함수를
 * 호출해 readValueChannel 채널로부터 값을 읽는다. 이 채널은 monitorValue()
 * 함수에서 제어한다. 그러면 value 변수에 현재 담겨 있는 값을 리턴한다. 반면 기존에
 * 저장된 값을 변경하려면 setValue() 함수를 호출한다. 이 함수는 writeValueChannel
 * 채널에 값을 쓴다. 이 함수 역시 select문을 통해 제어된다. 따라서 공유 변수 value는
 * monitorValue() 함수를 제외한 어느 누구도 직접 건드리지 않는다.
 */
func monitorValue() {
	var value int
	for {
		select {
		case newValue := <-writeValueChannel:
			value = newValue
			fmt.Printf("%d ", value)
		case readValueChannel <- value:
		}
	}
}

/*
 * 공유 메모리에 대한 마지막 절인 이 절에서는 특별히 지정한 Go 루팅으로 데이터를 공유하는
 * 방법을 소개한다. 공유 메모리는 스레드까리 통신하기 위한 수단 중에서도 에전 방식에 해당하지만
 * Go 언어에서는 하나의 Go 루틴이 데이터의 일정 영역을 소유할 수 있도록 동기화 기능을 제공한다.
 * 다시 말해, 어떤 Go 루틴이 가진 공유 데이터에 다른 Go 루틴이 접근하려면 반드시 데이터를
 * 보유한 Go 루틴에게 먼저 메시지를 보내야 한다. 이런 식으로 데이터가 손상되지 않도록 보호한다.
 * 이렇게 공유 데이터를 가진 Go 루팅을 '모니터 Go 루틴(Monitor Goroutine)이라 부른다.
 * Go 언어에서는 이를 '통신으로 공유한다(Sharing by communicating)'고 표현한다.
 *
 * 개인적으로 기존에 사용하던 공유 메모리 기법보다 모니터 Go 루틴을 선호한다. 그 이유는
 * 모니터 Go 루틴으로 구현하는 것이 좀 더 안정할 뿐만 아니라 Go 철학에 가깝기 때문이다.
 */
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please give an integer!")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Going to create %d random numbers.\n", n)
	rand.Seed(time.Now().Unix())
	go monitorValue()

	var w sync.WaitGroup

	for r := 0; r < n; r++ {
		w.Add(1)
		go func() {
			defer w.Done()
			setValue(rand.Intn(10 * n))
		}()
	}
	w.Wait()
	fmt.Printf("\nLast value: %d\n", readValue())
}
