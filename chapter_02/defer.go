package main

import (
	"fmt"
)

func d1() {
	for i := 9; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
}

func d2() {
	for i := 6; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
}

func d3() {
	for i := 3; i > 0; i-- {
		// defer문을 작성하는 가장 바람직한 방법으로
		// 원하는 변수를 명시적으로 익명 함수에 전달해 이해하기 쉽기 때문이다.
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
}

func main() {
	// 어떤 함수(A)를 호출하는 문장 앞에 defer 키워드를 붙이면,
	// 이런 defer문을 담고 있는 함수가 리턴될 떄까지 그 함수(A)의 실행을 미룬다
	// defer문은 파일 입력 및 출력 연산을 수행할 때 흔히 사용하는데,
	// 이렇게 하면 연 파일을 언제 닫을지 신경쓸 필요가 없기 때문이다.
	// 즉, defer 키워드를 사용하먼 연 파일을 닫는 함수를 그 파일을 여는 함수 가까이에 둘 수 있다.
	//
	// defer문을 담고 있는 함수가 리턴된 후에 defer 키워드를 이용해 실행이 미뤄진 함수(deferred function)가
	// 호출되는 순서는 LIFO(Last In First Out) 방식을 따른다.
	// 어떤 함수에서 defer로 호출할 함수를 f1(), f2(), f3()의 순서로 지정했다면
	// 나중에 그 함수가 리턴되는 시점에 defer로 지정했던 함수가 f3(), f2(), f1()의 순서로 실행된다는 뜻이다.
	d1()
	fmt.Println()
	d2()
	fmt.Println()
	d3()
}
