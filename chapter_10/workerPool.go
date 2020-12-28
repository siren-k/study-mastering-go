package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

// Client 구조체로 처리할 요청마다 고유 ID를 할당하는 기법을 볼 수 있다.
type Client struct {
	id      int
	integer int
}

// Data 구조체는 Client의 데이터에 대해 이 프로그램에서 실제로 생성한 결과를
// 그룹으로 묶는데 사용한다.
type Data struct {
	job    Client
	square int
}

// 세 개의 전역 변수를 선
var (
	size = 10
	// 버퍼 채널인 clients와 data는 클라이언트로부터 새로 요청을 받을 때
	// 그리고 이에 대한 결과를 쓸 때 사용한다. 프로그램을 좀 더 빠르게 만들고
	// 싶다면 size 매개변수릐 값을 높인다.
	clients = make(chan Client, size)
	data    = make(chan Data, size)
)

/*
 * worker() 함수는 처리할 요청을 clients 채널에서 읽는다. 요청에 대한 처리가
 * 다 끝나면 그 결과를 data 채널에 쓴다. 여기서 time.Sleep(time.Second)로
 * 설정한 지연 시간은 없어도 되지만 생성된 결과가 출력되는 방식을 이해하는데 도움이 된다.
 */
func worker(w *sync.WaitGroup) {
	for c := range clients {
		square := c.integer * c.integer
		output := Data{c, square}
		data <- output
		time.Sleep(time.Second)
	}
	w.Done()
}

/*
 * '워커 풀(Worker Pool)'이란 할당된 작업을 처리하려는 스레드의 집합이다. 아파치 웹 서버도 이 방식으로
 * 처리한다. 메인 프로세스는 들어온 요청을 모두 받아서 이를 실제로 처리할 워크 프로세스로 전달한다. 워커
 * 프로세스가 작업을 마치면 새로운 클라이언트를 받아들일 준비를 한다. 하지만 여기서 소개할 워커 풀은 스레드가
 * 아닌 Go 루틴을 사용한다는 점에서 차이가 있다. 또한 스레드는 일반적으로 특정한 요청을 처리한 뒤에 곧바로
 * 죽지 않는다. 스레드를 종료했다가 새로 생성하는데 드는 비용은 상당히 높은 편이다. 반면 Go 루틴은 작업을
 * 끌내자마자 죽는다.
 * Go 언어에서 제공하는 워커 풀은 버퍼 채널(Buffered Channel)로 구현한 것이다. 이렇게 하면 동시에
 * 실행되는 Go 루틴의 수를 제한할 수 있기 때문이다.
 *
 * 이 기법은 고급 테크닉에 속하는 것으로 Go 루틴을 이용해 여러 클라이언트로부터 요청을 받아서 서비스를
 * 제공하는 서버 프로세스를 생성하는데 적용할 수 있다.
 */
/*
 * makeWP() 함수의 목적은 모든 요청을 처리하는데 사용할 worker() Go 루틴의 개수를 알아내는 것이다.
 * makeWP()에서 w.Add(1) 함수를 호출하지만 워커가 주어진 작업을 마치면 worker() 함수 안에서
 * w.Done()이 호출된다.
 */
func makeWP(n int) {
	fmt.Println("make worker pool:", n)
	var w sync.WaitGroup
	for i := 0; i < n; i++ {
		w.Add(1)
		go worker(&w)
	}
	w.Wait()
	fmt.Println("data channel will be closed!")
	close(data)
}

/*
 * create() 함수의 목적은 모든 요청을 Client 타입으로 생성한 뒤에 이를 처리하도록
 * clients 채널로 쓰는 것이다. 여기서 clients 채널은 worker() 함수에서 읽는다는
 * 점을 주목한다.
 */
func create(n int) {
	for i := 0; i < n; i++ {
		c := Client{i, i}
		clients <- c
	}
	close(clients)
}

func main() {
	fmt.Println("Capacity of clients:", cap(clients))
	fmt.Println("Capacity of data:", cap(data))

	if len(os.Args) != 3 {
		fmt.Println("Need #jobs and #workers!")
		os.Exit(1)
	}

	nJobs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("job size:", nJobs)

	nWorkers, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("worker size:", nWorkers)

	// 서버에서 처리할 클라이언트 요청을 흉내내도록 create() 함수를 호출한다.
	go create(nJobs)
	// finished 채널은 익명 Go 루틴이 data 채널을 다 읽을 때까지 프로그램 실행을
	// 블록시키기 위한 용도로 사용한다.
	finished := make(chan interface{})
	// data 채널에서 값을 읽는 부분은 익명 Go 루틴으로 정의했다.
	// 읽은 결과는 화면에 출력한다.
	go func() {
		for d := range data {
			fmt.Printf("Client ID: %d\tint: ", d.job.id)
			fmt.Printf("%d\tsquare: %d\n", d.job.integer, d.square)
		}
		fmt.Println("here")
		//finished <- true
	}()

	// makeWP() 함수를 호출해서 요청을 실제로 처리한다.
	makeWP(nWorkers)
	// fmt.Printf() 블록에 '<- finished'라고 적은 부분은 main() 함수에서 정의한 익명 Go 루틴이
	// finished 채널에 뭔가를 쓰지 전까지 프로그램을 종료하지 않는다는 것을 의미한다.
	fmt.Printf(": %v\n", <-finished)
}
