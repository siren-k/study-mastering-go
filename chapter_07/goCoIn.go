package main

import "fmt"

/*
 * 타입에 함수를 추가해서 메소드를 정의하는 것이다.
 * 다시 말해 이렇게 작성한 함수와 타입은 일종의 오브젝트인 셈이다.
 */
type first struct{}

func (a first) F() {
	a.shared()
}

func (a first) shared() {
	fmt.Println("This is shared() from first")
}

/*
 * 새로 정의한 구조체 타입에 다른 타입을 집어 넣어서 일종의 계층을 형성하는 것이다.
 */
type second struct {
	first
}

func (a second) shared() {
	fmt.Println("This is shared() from second")
}

func main() {
	/*
	 * first{}.F()와 second{}.shared{}를 호출하면 각각 예상했던 결과를 출력한다.
	 *
	 * This is shared() from first
	 * This is shared() from second
	 */
	first{}.F()
	second{}.shared()

	/*
	 * j.F()를 호출하면 second.shared()가 아닌 여전히 first.shared()를 호출한다.
	 * second란 타입에서 shared()에 대한 구현을 변경했는데도 말이다. 이렇게 구현을 변경하는
	 * 기법을 OOP 용어로 '메소드 오버라이딩(Method Overriding)'이라고 부른다.
	 *
	 * This is shared() from first
	 */
	i := second{}
	j := i.first
	j.F()

	// ❯ go run goCoIn.go
	// This is shared() from first
	// This is shared() from second
	// This is shared() from first
}
