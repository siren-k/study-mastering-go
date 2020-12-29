package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----")
}

/*
 * go tool trace 유틸리티는 굉장히 간편하고 강력한 도구지만 성능에 관련된
 * 모든 문제를 이것만으로 해결할 수 없다. 프로그램이 특정한 함수에서 유난히
 * 많은 시간을 잡아 먹을 때에는 go tool pprof가 더 적합할 것이다.
 */
func main() {
	/*
	 * 트레이싱 파일 생성
	 */
	f, err := os.Create("traceFile.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	/*
	 * 트레이스 시작 및 종료
	 */
	err = trace.Start(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer trace.Stop()

	/*
	 * 프로그램 기능 구현
	 */
	var mem runtime.MemStats
	printStats(mem)
	for i := 0; i < 3; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
	}
	printStats(mem)
	for i := 0; i < 5; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
		time.Sleep(time.Millisecond)
	}
	printStats(mem)

	// go tool trace aTraceFile
}
