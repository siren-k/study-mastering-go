package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		if i%20 == 0 {
			continue
		}

		if i == 95 {
			break
		}

		fmt.Print(i, " ")
	}
	fmt.Println()

	// Go 언어는 while 루프를 작성하기 위한 키워드를 직접 제공하지 않고
	// for 루프를 이용하여 while 루프를 작성할 수 있다.
	// for 루프를 빠져나오려면 break문을 명시적으로 작성해야 한다.
	i := 10
	for {
		if i < 0 {
			break
		}
		fmt.Print(i, " ")
		i--
	}
	fmt.Println()

	i = 0
	anExpression := true
	for ok := true; ok; ok = anExpression {
		if i > 10 {
			anExpression = false
		}

		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

	// range 키워드를 for 루프 안에서 사용하면 Go 데이터 타입에 대해
	// 반복문을 수행하는 코드를 이해하기 쉽게 표현할 수 있다.
	// range 키워드의 가장 큰 장점은 슬라이스나 맵의 크기(Cardinality)를
	// 몰라도 그 안에 들어 있는 원소를 하나씩 처리할 수 있다는 것이다.
	anArray := [5]int{0, 1, -1, 2, -2}
	// 배열에 대한 변수에 range 키워드를 붙이면 두 개의 값을 리턴한다.
	// 하나는 배열의 인덱스고 다른 하나는 그 인덱스가 가리키는 원소의 값이다.
	for i, value := range anArray {
		fmt.Println("index:", i, "value:", value)
	}
}
