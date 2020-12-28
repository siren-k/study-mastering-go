package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	// sync.Mutex 타입은 Go 언어에서 뮤텍스를 구현한 것이다.
	// 뮤텍스란 일종의 상호 배제용 Lock이다.
	// 뮤텍스가 0이면 잠기지 않은 뮤텍스란 뜻이다.
	// 한 번 사용한 뮤텍스는 복사하면 안 된다.
	// ==> sync.Mutex 타입의 정의에서 특이한 점은 없다. 중요한 작업은 sync.Mutex를
	//     잠그거나 해제하는 sync.Lock()과 sync.Unlock() 함수에서 처리한다. 뮤텍스를
	//     잠근다는 말은 sync.Unlock() 함수로 그 뮤텍스를 해제하기 전까지 아무도 잠글 수
	//     없다는 것을 의미한다.
	m  sync.Mutex
	v1 int
)

func change(i int) {
	/*
	 * sync.Mutex.Lock()과 sync.Mutex.Unlock() 사이가 크리티컬 섹션이다.
	 */
	m.Lock()
	time.Sleep(time.Second)
	v1 = v1 + 1
	if v1%10 == 0 {
		v1 = v1 - 10*i
	}
	m.Unlock()
}

func read() int {
	m.Lock()
	a := v1
	m.Unlock()
	return a
}

/*
 * '공유 메모리(Shared Memory)와 공유 변수(Shared Variable)'는 유닉스 스레드끼리 서로 통신하는데 가장
 * 흔한 사용하는 방식이다.(Mutual Exclusion, 상호배제의 줄임말인) 뮤텍스(Mutex) 변수는 주로 스레드를
 * 동기화하고 여러 스레드가 동시에 공유 데이터를 동시에 쓸 때 이를 보호하기 위한 목적으로 주로 사용된다. 뮤텍스는
 * 마치 크기가 1인 버퍼 채널처럼 작동한다. 그래서 공유 변수에 접근할 수 있는 Go 루틴의 수는 최대 한 개다. 다시
 * 말해 두 개 이상의 Go 루틴이 이 변수를 동시에 업데이트할 수 없다.
 * '크리티컬 섹션(Critical Section, 임계영역)'이란 동시성 프로그램 코드 중에서 여러 프로세스, 스레드, Go
 * 루틴 등에서 동시에 실행할 수 없는 영역이다. 이 영역은 뮤텍스로 보호해야 한다. 따라서 코드에서 어느 부분이
 * 크리티컬 섹션인지를 먼저 밝혀둬야 전반적인 프로그래밍 과정을 순조롭게 진행할 수 있기 때문에 각별히 신경써서
 * 처리해야 한다.
 */
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please give me an integer!")
		return
	}

	numGR, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	var waitGroup sync.WaitGroup

	fmt.Printf("%d ", read())
	for i := 0; i < numGR; i++ {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()
			change(i)
			fmt.Printf("-> %d", read())
		}(i)
	}

	waitGroup.Wait()
	fmt.Printf("-> %d\n", read())
}
