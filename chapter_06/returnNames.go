package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
 * 리턴값에 이름을 지정했다. 하지만 좀 까다로운 부분이 있다. 바로 namedMinMax() 함수의
 * return문을 보면 리턴할 변수를 명시적으로 지정하지 않은 점이다. 하지만 이 함수는 시그니처에서
 * 리턴값에 이름을 지정했기 때문에 함수를 정의할 때 min과 max가 순서대로 값이 리턴된다.
 */
func namedMinMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return
}

/*
 * return문에 리턴할 변수와 순서를 명시적으로 지정했다.
 */
func minMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return min, max
}

func main() {
	arguments := os.Args
	if len(arguments) < 3 {
		fmt.Println("The program needs at least 2 arguments!")
		return
	}

	a1, _ := strconv.Atoi(arguments[1])
	a2, _ := strconv.Atoi(arguments[2])

	fmt.Println(minMax(a1, a2))
	min, max := minMax(a1, a2)
	fmt.Println(min, max)

	fmt.Println(namedMinMax(a1, a2))
	min, max = namedMinMax(a1, a2)
	fmt.Println(min, max)
}
