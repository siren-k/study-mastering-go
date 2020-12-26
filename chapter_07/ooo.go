package main

import "fmt"

type aa struct {
	XX int
	YY int
}

type bb struct {
	AA string
	XX string
}

/*
 * 합성(Composition) 기능을 이용하면 여러 개의 struct 타입을 하나의 구조체로 만들 수 있다.
 * 여기서는 cc라는 데이터 타입으로 aa 타입 변수와 bb 타입 변수를 하나로 묶고 있다.
 */
type cc struct {
	A aa
	B bb
}

/*
 * 두 메소드는 이름이 같이만(A()), 함수 헤더는 다르다.
 * 첫 번째 메소드는 aa 타입 변수를 사용하고, 두 번째 메소드는 bb 타입 변수를 사용한다.
 * 여기에 나온 기법을 사용하면 다른 타입에 속한 함수의 이름을 똑같이 표현할 수 있다.
 */
func (A aa) A() {
	fmt.Println("Function A() for A")
}

func (B bb) A() {
	fmt.Println("Function A() for B")
}

/*
 * Go 언어는 상속을 지원하지 않는다. 대신 합성(Composition)을 지원하고, 인터페이스로
 * 다형성(Polymorphism)을 지원한다. 따라서 Go 언어는 OOP 언어는 아니지만, 어느 정도
 * OOP를 흉내낼 수 있는 기능은 몇 가지 갖추고 있다.
 *
 * 객체지향 방법론에 따라 애플리케이션을 개발하고 싶다면 Go 언어 보다는 다른 언어를 선책하는 편이 좋다.
 * 개인적으로 자바(Java)를 그다지 좋아하지 않기 때문에 C++이나 파이썬(Python)을 추천한다. Go 언어는
 * 다루기도 힘들고 유지보수도 어려운 복잡하고 계층이 깊은 타입을 사용하는 것을 금지하고 있다.
 *
 * 추상 클래스(Abstract Class)나 상속(Inheritance)을 지원하는 다른 OOP 언어로
 * 작성한 코드에 비해 단순한 편이다. 하지만 구조체를 가진 요소나 타입을 생성하거나,
 * 이름이 같은 메소드로 여러 가지 데이터 타입을 다루기에는 이 방식이 더 적합하다.
 */
func main() {
	var i cc
	i.A.A()
	i.B.A()
}
