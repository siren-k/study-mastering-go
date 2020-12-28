package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

/*
 * '데이터 경쟁 상태(data Race Condition)'란 두 개 이상의 스레드나 Go 루틴과 같은 요소들이 공유
 * 리소스나 프로그램의 변수를 서로 제어하거나 수정하려고 경쟁하는 상황을 말한다. 구체적으로 데이터 경쟁은
 * 두 개 이상의 인스트럭션(명령)이 동일한 메모리 주소를 접근하는 상환에서 어느 하나가 쓰기 연상을 수행할
 * 때 발생한다.
 *
 * Go 소스 파일을 빌드하거나 실행할 때 -race 플래그를 지정하면 Go 언어에서 제공하는 '레이스 디텍터
 * (Race Detector, 경쟁 상태 감지기)'가 작동하면서 컴파일러는 기존 실행 파일과는 다른 형태로 실행
 * 파일을 생성한다. 이렇게 수정된 버전은 공유 메모리에 대한 접근하는 모든 경우와 sync.Mutex나 sync.WaitGroup
 * 에 대한 호출을 비롯한 모든 동기화 이벤트를 기록한다. 레이스 디텍터는 이러한 이벤트를 분석한 뒤에 그
 * 결과를 출력한다. 이를 통해 프로그래머는 어디서 문제가 발생했는지 쉽게 찾아서 수정할 수 있다.
 */
func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Give me a natural number")
		return
	}

	numGR, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	var waitGroup sync.WaitGroup
	var i int

	k := make(map[int]int)
	k[1] = 12

	for i = 0; i < numGR; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			k[i] = i
		}()
	}

	k[2] = 10
	waitGroup.Wait()
	fmt.Printf("k = %v\n", k)

	// ❯ go run -race raceC.go 10
	// ==================
	// ==> 첫 번째 데이터 경쟁 상태는 main.main.func1() 안에서 발생한다. 이 함수는 Go 루틴에서 호출하는 for 루프에서
	//     호출한다. Previous write 메시지를 보면 문제에 대한 정보를 좀 더 볼 수 있다. 여기서 가리키는 코드를 붆석해보면,
	//     실제 원인은 익명 함수에서 매개변수를 받지 않기 때문이다. 다시 말해 for 루프에서 사용하는 i 값을 일정하게 받아올 수
	//     없다. for 루프 안에서 쓰기 연산을 수행해서 값이 계속 바뀌기 때문이다.
	// WARNING: DATA RACE
	// Read at 0x00c0000140d0 by goroutine 7:
	//   main.main.func1()
	//       /Users/benjamin/Lab/github/siren-k/study-mastering-go/chapter_10/raceC.go:45 +0x8c
	//
	// Previous write at 0x00c0000140d0 by main goroutine:
	//   main.main()
	//       /Users/benjamin/Lab/github/siren-k/study-mastering-go/chapter_10/raceC.go:41 +0x22e
	//
	// Goroutine 7 (running) created at:
	//   main.main()
	//       /Users/benjamin/Lab/github/siren-k/study-mastering-go/chapter_10/raceC.go:43 +0x204
	// ==================
	// ==================
	// ==> 두 번째 데이터 경쟁 상태에 대한 메시지를 보면 'Write at 0x00c00007a180 by goroutine 8:'라고 나와 있다.
	//      그 아래 나온 메시지를 좀 더 읽어보면 이 경쟁 상태는 쓰기 연산에 관련된 것으로, 하나의 Go 맵에 최소 두 개의 Go 루틴이
	//      쓴다는 것을 알 수 있다. 이러한 Go 루틴은 둘 다 이름이 같아서(main.main.func1()) 동일한 Go 루틴을 가리킨다는
	//      것을 알 수 있다. 이처럼 두 Go 루틴이 동일한 변수에 쓰는 것이 바로 데이터 경쟁 상태다.
	// WARNING: DATA RACE
	// Write at 0x00c00007a180 by goroutine 8:
	//   runtime.mapassign_fast64()
	//       /usr/local/Cellar/go/current/libexec/src/runtime/map_fast64.go:92 +0x0
	//   main.main.func1()
	//       /Users/benjamin/Lab/github/siren-k/study-mastering-go/chapter_10/raceC.go:45 +0xb9
	//
	// Previous write at 0x00c00007a180 by goroutine 7:
	//   runtime.mapassign_fast64()
	//       /usr/local/Cellar/go/current/libexec/src/runtime/map_fast64.go:92 +0x0
	//   main.main.func1()
	//       /Users/benjamin/Lab/github/siren-k/study-mastering-go/chapter_10/raceC.go:45 +0xb9
	//
	// Goroutine 8 (running) created at:
	//   main.main()
	//       /Users/benjamin/Lab/github/siren-k/study-mastering-go/chapter_10/raceC.go:43 +0x204
	//
	// Goroutine 7 (finished) created at:
	//   main.main()
	//       /Users/benjamin/Lab/github/siren-k/study-mastering-go/chapter_10/raceC.go:43 +0x204
	// ==================
	// k = map[1:12 2:10 3:3 5:6 6:6 7:8 8:8 10:10]
	// Found 4 data race(s)
	// exit status 66
}
