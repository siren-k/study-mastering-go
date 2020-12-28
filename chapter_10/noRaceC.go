package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

// 데이터 경쟁 상태 제거를 위하여 sync.Mutex 변수 추가
var aMutex sync.Mutex

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
		go func(i int) {
			defer waitGroup.Done()
			aMutex.Lock()
			k[i] = i
			aMutex.Unlock()
		}(i)
	}

	aMutex.Lock()
	k[2] = 10
	aMutex.Unlock()
	waitGroup.Wait()
	fmt.Printf("k = %v\n", k)

	// ❯ go run -race noRaceC.go 10
	// k = map[0:0 1:1 2:10 3:3 4:4 5:5 6:6 7:7 8:8 9:9]
}
