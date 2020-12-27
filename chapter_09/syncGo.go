package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	n := flag.Int("n", 10, "Number of goroutines")
	flag.Parse()
	count := *n
	fmt.Printf("Going to create %d goroutines.\n", count)

	// sync 패키지의 소스 중 waitGroup.go를 보면
	// sync.WaitGroup 타입은 세 개의 필드로 구성된 구조체라는 것을 알 수 있다.
	// type WaitGroup struct {
	//     noCopy noCopy
	//     state1 [12]byte
	//     sema unit32
	// }
	var waitGroup sync.WaitGroup
	fmt.Printf("%#v\n", waitGroup)

	for i := 0; i < count; i++ {
		// sync.WaitGroup에 속한 Go 루틴의 수는 sync.Add() 함수의 호출 횟수에
		// 따라 정의된다.
		// sync.Add() 함수를 호출할 때마다 sync.WaitGroup 변수의 카운터가 증가한다.
		// 여기서 주목할 부분은 go문이 나오기 전에 sync.Add(1)를 호출해야 한다는 점이다.
		// 그래야 '경쟁 상태(Race Condition)'가 발생하는 것을 피할 수 있다.
		waitGroup.Add(1)
		go func(x int) {
			// 각각의 Go 루틴이 작업을 마칠 때마다 sync.Done() 함수가 실행된다.
			// 이 함수를 통해 sync.WaitGroup 변수의 카운터 값을 감소시킨다.
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	fmt.Printf("%#v\n", waitGroup)
	// sync.Wait()을 호출하면 sync.WaitGroup 변수에 있는 카운터가 0이 될 때까지
	// 실행을 멈춘다. 따라서 Go 루틴이 마칠 때까지 충분히 기다릴 수 있다.
	waitGroup.Wait()
	fmt.Println("\nExiting...")

	// 여전히 실행 순서가 제각각인 것을 볼 수 있다. 특히 생성한 Go 루틴의 개수가 많을 수록 심하다. 대부분은 이렇게
	// 실행되어도 상관없지만, 간혹 이렇게 실행되면 안될 때가 있다. 또한 Go 루틴의 수가 30개를 넘어가면, 두 번째로
	// 나온 fmt.Printf("%#v\n", waitGroup)이 실핼되기도 전에 몇몇 Go 루틴의 실행이 끝나 버릴 수 있다.
	// 마지막으로 주의할 점은, sync.WaitGroup의 state1 필드에 있는 원소 중 하나는 sync.Add()와 sync.Done()을
	// 호출할 때마다 증가하거나 감소하는 카운터 값을 가지고 있다는 점이다.
	//
	// sync.Add() 함수와 sync.Done() 함수의 실행 횟수가 다른 경우 에러가 발생한다.
	// 실전에서 프로그램을 작성할 때는 sync.Add()와 sync.Done()을 호출한 횟수를 주의 깊게 관리해야 한다.
	//
	// ❯ go run syncGo.go -n 100
	// Going to create 100 goroutines.
	// sync.WaitGroup{noCopy:sync.noCopy{}, state1:[3]uint32{0x0, 0x0, 0x0}}
	// 5 0 1 2 3 4 7 6 11 10 12 8 28 14 15 16 17 18 19 20 21 22 23 13 26 27 25 9 45 30 39 32 37 41 43 33 24 40 35 38 29 83 34 75 61 64 76 62 63 60 99 66 65 68 72 84 48 42 50 67 77 91 49 97 85 95 78 88 82 81 44 47 46 69 36 51 73 52 56 74 70 92 93 54 57 94 96 71 58 80 59 79 98 sync.WaitGroup{noCopy:sync.noCopy{}, state1:[3]uint32{0x0, 0x33, 0x0}}
	// 86 55 53 89 90 87 31
	// Exiting...
	// ❯ go run syncGo.go -n 100
	// Going to create 100 goroutines.
	// sync.WaitGroup{noCopy:sync.noCopy{}, state1:[3]uint32{0x0, 0x0, 0x0}}
	// 1 0 8 2 3 4 5 6 7 17 15 16 10 9 21 18 19 20 24 22 23 31 25 26 27 28 29 30 12 35 32 33 34 40 36 37 38 39 47 41 42 43 44 45 46 13 52 48 49 50 51 14 55 64 61 58 11 53 67 63 70 69 68 65 59 73 71 56 75 62 60 76 84 74 72 66 54 82 80 91 95 79 83 99 97 96 92 77 93 88 89 81 90 sync.WaitGroup{noCopy:sync.noCopy{}, state1:[3]uint32{0x0, 0x17, 0x0}}
	// 94 85 78 98 57 86 87
	//Exiting...
}
