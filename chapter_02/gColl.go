package main

import (
	"fmt"
	"runtime"
	"time"
)

func printStatus(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem) // 가비지 컬렉션 통계에 대한 최신 정보를 조회하려면
	// 매번 runtime.ReadMemStats() 함수를 호출해야 한다
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----")
}

func main() {
	var mem runtime.MemStats
	printStatus(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
	}
	printStatus(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
		time.Sleep(5 * time.Second)
	}
	printStatus(mem)

	// ❯ go run gColl.go
	// mem.Alloc: 167240
	// mem.TotalAlloc: 167240
	// mem.HeapAlloc: 167240
	// mem.NumGC: 0
	// -----
	// mem.Alloc: 50166360
	// mem.TotalAlloc: 500218008
	// mem.HeapAlloc: 50166360
	// mem.NumGC: 9
	// -----
	// mem.Alloc: 162600
	// mem.TotalAlloc: 1500299440
	// mem.HeapAlloc: 162600
	// mem.NumGC: 20
	// -----

	// ❯ GODEBUG=gctrace=1 go run gColl.go
	//   ==> go run 커맨드 앞에 GODEBUG=1을 붙이면 가비지 컬렉터의 작동 과정에 대한 분석 데이터가 출력된다
	// gc 1 @0.028s 0%: 0.012+0.28+0.016 ms clock, 0.15+0.11/0.39/0.18+0.20 ms cpu, 4->4->0 MB, 5 MB goal, 12 P
	//                                                                              4->4->0
	//                                                                              4 ==> 가비지 컬렉터가 실행될 시점의 힙 크기
	//                                                                                 4 ==> 가비지 컬렉터가 실핻을 마칠 시점의 힙 크기
	//                                                                                    0 ==> 현재 힙 크기
	// gc 2 @0.048s 0%: 0.070+0.29+0.016 ms clock, 0.84+0.12/0.39/0.23+0.19 ms cpu, 4->4->0 MB, 5 MB goal, 12 P
	// gc 3 @0.065s 2%: 1.1+1.7+0.049 ms clock, 13+0.74/0.96/0.006+0.59 ms cpu, 4->4->0 MB, 5 MB goal, 12 P
	// gc 4 @0.088s 2%: 0.23+0.40+0.017 ms clock, 2.7+0.23/0.60/0.57+0.20 ms cpu, 4->4->0 MB, 5 MB goal, 12 P
	// gc 5 @0.093s 2%: 0.097+0.42+0.002 ms clock, 1.1+0.17/0.44/0.60+0.033 ms cpu, 4->4->0 MB, 5 MB goal, 12 P
	// gc 6 @0.098s 2%: 0.052+0.38+0.048 ms clock, 0.63+0.21/0.60/0.41+0.58 ms cpu, 4->4->0 MB, 5 MB goal, 12 P
	// gc 7 @0.101s 2%: 0.046+0.38+0.011 ms clock, 0.55+0.049/0.55/0.43+0.14 ms cpu, 4->4->0 MB, 5 MB goal, 12 P
	// # command-line-arguments
	// gc 1 @0.003s 4%: 0.010+1.0+0.015 ms clock, 0.12+0.076/2.2/0.99+0.18 ms cpu, 4->4->3 MB, 5 MB goal, 12 P
	// gc 2 @0.006s 6%: 0.047+1.0+0.017 ms clock, 0.56+0.22/2.3/1.4+0.21 ms cpu, 5->7->4 MB, 6 MB goal, 12 P
	// # command-line-arguments
	// gc 1 @0.001s 6%: 0.007+1.7+0.015 ms clock, 0.094+0.16/2.1/1.6+0.18 ms cpu, 4->6->5 MB, 5 MB goal, 12 P
	// gc 2 @0.006s 5%: 0.006+1.0+0.004 ms clock, 0.083+0/2.2/0.59+0.048 ms cpu, 9->9->6 MB, 10 MB goal, 12 P
	// gc 3 @0.014s 3%: 0.013+2.0+0.014 ms clock, 0.16+0/2.7/0.28+0.17 ms cpu, 12->12->8 MB, 13 MB goal, 12 P
	// gc 4 @0.036s 2%: 0.017+1.8+0.014 ms clock, 0.20+0.065/2.5/0.047+0.17 ms cpu, 15->15->9 MB, 17 MB goal, 12 P
	// mem.Alloc: 167704
	// mem.TotalAlloc: 167704
	// mem.HeapAlloc: 167704
	// mem.NumGC: 0
	// -----
	// gc 1 @0.001s 1%: 0.010+0.14+0.004 ms clock, 0.12+0.055/0.078/0.095+0.057 ms cpu, 47->47->0 MB, 48 MB goal, 12 P
	// gc 2 @0.024s 0%: 0.048+0.14+0.003 ms clock, 0.58+0.11/0.073/0.007+0.040 ms cpu, 47->47->0 MB, 48 MB goal, 12 P
	// gc 3 @0.026s 0%: 0.021+0.11+0.002 ms clock, 0.25+0.093/0.045/0.002+0.033 ms cpu, 47->47->0 MB, 48 MB goal, 12 P
	// gc 4 @0.027s 0%: 0.044+0.15+0.002 ms clock, 0.53+0.12/0.066/0.014+0.031 ms cpu, 47->47->0 MB, 48 MB goal, 12 P
	// gc 5 @0.029s 0%: 0.021+0.10+0.002 ms clock, 0.25+0.090/0.056/0.016+0.033 ms cpu, 47->47->0 MB, 48 MB goal, 12 P
	// gc 6 @0.031s 0%: 0.021+0.10+0.003 ms clock, 0.26+0.079/0.041/0+0.037 ms cpu, 47->47->0 MB, 48 MB goal, 12 P
	// gc 7 @0.033s 0%: 0.022+0.11+0.002 ms clock, 0.26+0.085/0.048/0.003+0.030 ms cpu, 47->47->0 MB, 48 MB goal, 12 P
	// gc 8 @0.034s 0%: 0.023+0.11+0.003 ms clock, 0.28+0.087/0.051/0+0.038 ms cpu, 47->47->0 MB, 48 MB goal, 12 P
	// gc 9 @0.036s 1%: 0.023+0.15+0.003 ms clock, 0.27+0.11/0.10/0.004+0.036 ms cpu, 47->47->0 MB, 48 MB goal, 12 P
	// mem.Alloc: 50166280
	// mem.TotalAlloc: 500218368
	// mem.HeapAlloc: 50166280
	// mem.NumGC: 9
	// -----
	// gc 10 @0.038s 1%: 0.045+1.4+0.010 ms clock, 0.54+0.060/0.12/0.093+0.13 ms cpu, 47->143->95 MB, 48 MB goal, 12 P
	// gc 11 @5.045s 0%: 0.037+0.16+0.006 ms clock, 0.45+0/0.23/0.032+0.073 ms cpu, 190->190->0 MB, 191 MB goal, 12 P
	// gc 12 @10.101s 0%: 0.049+0.14+0.003 ms clock, 0.59+0/0.15/0.011+0.047 ms cpu, 95->95->0 MB, 96 MB goal, 12 P
	// gc 13 @15.106s 0%: 0.032+0.16+0.004 ms clock, 0.39+0/0.24/0.011+0.059 ms cpu, 95->95->0 MB, 96 MB goal, 12 P
	// gc 14 @20.111s 0%: 0.031+0.14+0.005 ms clock, 0.37+0/0.20/0.031+0.062 ms cpu, 95->95->0 MB, 96 MB goal, 12 P
	// gc 15 @25.118s 0%: 0.033+0.14+0.005 ms clock, 0.39+0/0.18/0.028+0.063 ms cpu, 95->95->0 MB, 96 MB goal, 12 P
	// gc 16 @30.125s 0%: 0.054+0.17+0.005 ms clock, 0.65+0/0.22/0.025+0.061 ms cpu, 95->95->0 MB, 96 MB goal, 12 P
	// gc 17 @35.131s 0%: 0.031+0.13+0.005 ms clock, 0.38+0/0.19/0.019+0.061 ms cpu, 95->95->0 MB, 96 MB goal, 12 P
	// gc 18 @40.138s 0%: 0.055+0.19+0.004 ms clock, 0.66+0/0.27/0.048+0.055 ms cpu, 95->95->0 MB, 96 MB goal, 12 P
	// gc 19 @45.145s 0%: 0.055+0.18+0.005 ms clock, 0.66+0/0.21/0.046+0.062 ms cpu, 95->95->0 MB, 96 MB goal, 12 P
	// mem.Alloc: 164368
	// mem.TotalAlloc: 1500301648
	// mem.HeapAlloc: 164368
	// mem.NumGC: 19
	// -----
}
