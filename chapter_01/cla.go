package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 커맨드라인 인수가 들어왔는지 확인
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}

	arguments := os.Args                           // os.Args ==> string 값을 가지는 Go 슬라이스
	min, _ := strconv.ParseFloat(arguments[1], 64) // 에러는 '_'[빈 식별자:Blank Identifier]를 이용하여 무시하도록 작성
	max, _ := strconv.ParseFloat(arguments[1], 64)

	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
