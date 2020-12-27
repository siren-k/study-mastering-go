package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func gen(min, max int, createNumber chan int, end chan bool) {
	for {
		/*
		 * select 블록에서 실제로 어떤 일이 발생할까? 이 select문은 크게 세 가지 case문으로
		 * 구성된다. 여기서는 default 브랜치를 지정하지 않았다. 세 번째 case문이 실질적으로
		 * default 브랜치 역할을 한다. time.After() 함수는 일정한 시간이 지나면 리턴한다.
		 * 따라서 다른 채널이 블록됐다면 selelct문에 대한 블록도 해제한다.
		 *
		 * select문은 순차적으로 실행되지 않는다. 모든 채널을 동시에 확인한다. select문에 있는
		 * 채널 중 어떤 것도 사용할 수 있는 상태가 아니라면 어느 한 채널이 사용 가능한 상태가 될
		 * 때까지 블록된다. select문에서 사용 가능한 상태의 채널이 여러 개라면 Go 런타임은
		 * 그 중 하나를 임의로 선택한다. 이 때, 런타임은 최대한 공평하게 선택한다.
		 *
		 * select의 가장 큰 장점은 여러 채널에 연결할 수 있을 뿐만 아니라 이러한 채널을 관리하고
		 * 각 채널의 동작을 조율할 수 있다는 점이다. 채널은 Go 루틴에 연결돼 있기 때문에 select는
		 * 이렇게 Go 루틴에 연결된 채널을 연결한다. 따라서 select문은 Go 언어에서 제공하는 동시성
		 * 모델에서 상당히 중요한 기능에 속한다.
		 */
		select {
		case createNumber <- rand.Intn(max-min) + min:
		case <-end:
			close(end)
			return
		case <-time.After(4 * time.Second):
			fmt.Println("\ntime.After()!")
		}
	}
}

/*
 * select 키워드는 상당히 강력한 기능을 제공하며 이를 적용할 수 있는 상황도 굉장히 다양하다.
 * Go 언어에서 select문은 채널에 대한 switch문과 같다. 실제로 select를 이용해 Go 루틴이
 * 여러 개의 통신 연산을 기다리게 할 수 있다. 따라서 select의 가장 큰 장점은 하나의 select
 * 블록에서 여러 개의 채널을 다룰 수 있다는 것이다. 결론적으로 채널에 대해 논블로킹 방식으로
 * 연산을 수행할 수 있다.
 *
 * select 키워드로 여러 채널을 다룰 때 발생할 수 있는 가장 큰 문제는 데드락(Deadlock)이다.
 * 따라서 프로세스를 설계하고 개발할 때 데드락이 발생하지 않도록 각별히 신격써야 한다.
 */
func main() {
	rand.Seed(time.Now().Unix())
	createNumber := make(chan int)
	end := make(chan bool)

	if len(os.Args) != 2 {
		fmt.Println("Please give me an integer!")
		return
	}

	// strconv.Atoi() 함수에서 리턴한 error 값을 공간 절약을 위해 하지 않았다. 실전에서는
	// 이렇게 작성하면 안 된다.
	n, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Going to create %d random numnbers.\n", n)

	go gen(0, 2*n, createNumber, end)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", <-createNumber)
	}

	// gen() 함수의 time.After() 함수가 리턴할 때까지 시간을 충분히 주기 위해서 작성하였다.
	time.Sleep(5 * time.Second)
	fmt.Println("Exiting...")
	// gen() 함수 안에 있는 select문의 'case <- end' 브랜치를 실행시킨다.
	end <- true
}
