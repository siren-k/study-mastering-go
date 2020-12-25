package main

import (
	"container/ring"
	"fmt"
)

// size 변수는 생성할 링의 크기를 표현한다.
var size int = 10

func main() {
	// ring.New() 함수를 통해 새로운 링을 생성한다.
	// 이 때, 링의 크기를 표현하는 매개변수 하나를 지정해야 한다.
	myRing := ring.New(size + 1)
	fmt.Println("Empty ring:", *myRing)

	for i := 0; i < myRing.Len()-1; i++ {
		myRing.Value = i
		myRing = myRing.Next()
	}

	// 링에 2라는 값을 추가한다. 하지만 이 값은 이미 for 루프 안에서 추가한 값이다.
	// 마지막으로 링에서 0에 해당하는 값은 nil이란 값을 가진 원소 하나만 있는 링이다.
	myRing.Value = 2

	sum := 0
	// ring.Do() 함수는 링의 각 원소에 대해 함수를 호출할 때 사용한다. 그런데 이 함수가
	// 링을 수정하면 ring.Do()의 동작을 알 수 없게 된다. x.(init)이라고 적은 부분을
	// 타입 어써션(Type Assertion)이라 부른다.
	myRing.Do(func(x interface{}) {
		t := x.(int)
		sum = sum + t
	})
	fmt.Println("Sum:", sum)

	/*
	 * 링이 가지는 유일한 단점은 ring.Next()가 끊임없이 호출될 수 있다는 점이다. 따라서
	 * 이 작업을 멈추는 수단을 마련해야 한다. 여기서는 ring.Len() 함수를 이용했다. 개인적으로는
	 * 링의 모든 원소에 대해 순환하는 작업을 ring.Do() 함수로 처리하는 방식을 선호한다. 이렇게
	 * 하면 코드가 훨씬 깔끔하다. 물론 for 루프도 그리 나쁘지 않다.
	 */
	for i := 0; i < myRing.Len()+2; i++ {
		myRing = myRing.Next()
		fmt.Print(myRing.Value, " ")
	}
	fmt.Println()

	// ❯ go run conRing.go
	// Empty ring: {0xc00000c080 0xc00000c1a0 <nil>}
	// Sum: 47
	//
	// 결과를 보면 링에 중복된 값이 들어 있는 것을 볼 수 있다. 다시 말해 ring.Len() 함수가 없다면
	// 링의 크기를 제대로 알아낼 방법이 없다는 것을 알 수 있다.
	// 0 1 2 3 4 5 6 7 8 9 2 0 1
}
