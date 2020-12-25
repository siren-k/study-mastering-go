package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
 * 여러 값을 리턴하는 함수
 * Go 언어의 함수는 여러 개의 값을 동시에 리턴할 수 있다. 이렇게 하면 함수로부터 리턴되는 여러 값에 접근하기
 * 위해 별도로 구조체를 정의하지 않아도 돼서 편하다. 예를 들어 int 값 두 개와, float64 값 한 개,
 * string 값 한 개, 총 네 개의 값을 리턴하는 함수를 정의하는 방법을 살펴보자. 먼저 이 함수는 다음과 같이
 * 선언한다.
 *
 * func aFunction() (int, int, float64, string) {
 * }
 */
func doubleSquare(x int) (int, int) {
	return x * 2, x * x
}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("The program needs 1 argument!")
		return
	}
	// ❯ go run functions.go
	// The program needs 1 argument!

	y, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	/*
	 * 익명 함수(Anonymous Function)는 따로 이름을 붙이지 않고 코드 안에서 정의하며, 주로 짧은
	 * 코드로 구현하는 용도로 사용된다. Go 언어에서 함수는 익명 함수를 리턴할 수도, 매개변수로 받을
	 * 수도 있따. 또한, 익명 함수를 변수에 연결할 수도 있다.
	 * 함수형 프로그래밍에서는 이러한 익명 함수를 클로저(Closure)라고도 부른다.
	 * 익명 함수를 이용하면 작업을 굉장히 쉽고 편하게 처리할 수 있다. 단, 특별한 이유 없이 익명 함수를
	 * 남용하지 않도록 주의한다.
	 *
	 * 좀 더 일반적으로 람다(Lambda) 함수라 표현하며, 함수 안에 정의되어 있지 않는 자유 변수가 담긴
	 * 오픈 람다 함수를 그 변수의 값을 확정할 수 있도록 문맥/환경을 제공해서 자유 변수가 없는 닫힌(Closed)
	 * 람다 함수로 만든 것을 클러저라 부른다. 뒤에 나온 익명 함수 예제를 보면 다른 언어의 클로저와 비슷하다는
	 * 것을 알 수 있다.
	 *
	 * 익명 함수는 로컬 영역에서 간단한 작업을 수행할 때만 활용하는 것이 바람직하다. 익명 함수가 다루는
	 * 영역이 로컬을 벗어난다면 일반 함수로 정의하는 것이 좋다.
	 */
	square := func(s int) int {
		return s * s
	}
	fmt.Println("The square of", y, "is", square(y))

	double := func(s int) int {
		return s + s
	}
	fmt.Println("The double of", y, "is", double(y))

	fmt.Println(doubleSquare(y))
	d, s := doubleSquare(y)
	fmt.Println(d, s)

	// ❯ go run functions.go 10
	// The square of 10 is 100
	// The double of 10 is 20
	// 20 100
	// 20 100
}
