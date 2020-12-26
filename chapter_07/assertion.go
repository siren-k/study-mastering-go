package main

import "fmt"

/*
 * 타입 어써선(Type Assertion)(타입 단언)이란 x.(T) 형식의 표기법으로, x는 인터페이스 타입을,
 * T는 구체적인 타입을 지정한다. 이 때, x에 실제로 저장되는 값의 타입은 T이며, T는 반드시 x의
 * 인터페이스 타입을 충족해야 한다. 타입 어써션의 정의 방식이 다소 생소하게 느껴질 수 있다.
 *
 * 타입 어써션으로 할 수 있는 일은 두 가지가 있다.
 * 첫 번째는 인터페이스 값이 특정한 타입을 따르는지 확인하는 것이다. 타입 어써션을 이렇게 활용할 때는
 * 내부값과 bool 타입의 값이라는 두 가지 값을 리턴한다. 내부 값은 실제로 사용하려는 값인 반면, 불리언
 * 값은 그 타입 어써션이 제대로 실행됐는지 여부를 알려준다.
 *
 * 타입 어써션으로 할 수 있는 두 번째 일은 인터페이스에 저장된 구체적인 값을 이용하거나 이 값을 새 변수에
 * 할당하는 것이다. 다시 말해 인터페이스에 int 변수가 있을 때 타입 어써션을 이용해 그 값을 가져올 수 있다.
 *
 * 그런데 타입 어써션에 오류가 발생했을 때 이를 제대로 처리하지 않으면 프로그램은 뻗어버린다.
 */
func main() {
	// myInt 변수를 선언하였고 이 변수는 int 타입의 123을 동적으로 지정했다.
	var myInt interface{} = 123

	// 타입 어써션으로 int에 대해 테스트
	k, ok := myInt.(int)
	if ok {
		fmt.Println("Success:", k)
	}

	// 타입 어써션으로 float64에 대해 테스트
	// myInt 변수에는 float64 타입의 값이 없기 때문에 myInt.(float64)라고 적은 타입 어써션에서
	// 발생하는 오류를 적절히 처리하지 않으면 프로그램이 멈춘다. 예제에서는 ok 변수를 적절히 활용하여
	// 프로그램 전체가 뻗지 않게 했다.
	v, ok := myInt.(float64)
	if ok {
		fmt.Print(v)
	} else {
		fmt.Println("Failed without panicking!")
	}

	// 변수 i의 타입이 int이고, 값은 myInt에 저장돼 있던 123이라는 것을 알 수 있다. 이처런 int는
	// myInt 인터페이스를 충족하는 타입이고, myInterface를 구현하는 함수가 필요 없기 때문에
	// myInt.(int)는 int 값이 된다.
	i := myInt.(int)
	fmt.Println("No checking:", i)

	// myInt.(bool)을 실행할 때 오류가 발생한다. myInt에 담긴 값은 불리언(Boolean) 값이 아니기
	// 때문이다.
	j := myInt.(bool)
	fmt.Println(j)

	// ❯ go run assertion.go
	// Success: 123
	// Failed without panicking!
	// No checking: 123
	// panic: interface conversion: interface {} is int, not bool
	//
	// goroutine 1 [running]:
	// main.main()
	//         /Users/benjamin/Lab/github/siren-k/study-mastering-go/chapter_07/assertion.go:23 +0x1b8
	// exit status 2
}
