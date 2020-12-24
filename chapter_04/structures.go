package main

import "fmt"

func main() {
	/*
	 * 다양한 변수를 묶어서 새로운 타입으로 정의하고 싶다면 구조체를 사용해야 한다.
	 * 이렇게 구조체를 구성하는 요소를 '구조체의 필드(Fields of the structure), 또는 간단히 '필드(Field)'라고 부른다.
	 *
	 * 구조체의 타입 동일성(Type Identity)을 판단할 때는 그 구조체 타입을 정의할 때 나열한 필드의 순서가 중요하다.
	 * 쉽게 말해 서로 같은 필드로 구성된 구조체가 두 개 있을 때, 그 안에 필드가 나열된 순서가 정확히 일치하지 않으면
	 * Go 언어는 두 개가 서로 다르다고 판단한다.
	 */
	// 구조체는 Go 패키지 전체에서 사용할 수 있도록, 스코프를 글로벌(전역 범위)로 지정정하기 위해
	// main() 함수 밖에서 정의할 때가 많다. 물론 특별히 현재 스코프에서만 사용하고 다른 곳에서는
	// 쓸 일이 없을 때는 로컬(지역 범위)로 정의해도 된다.
	type XYZ struct {
		X int
		Y int
		Z int
	}

	var s1 XYZ
	fmt.Println(s1.Y, s1.Z)

	/*
	 * 구조체 리터럴(Structure Literal) p1과 p2를 정의
	 */
	p1 := XYZ{23, 12, -2}
	p2 := XYZ{Z: 12, Y: 13}
	fmt.Println(p1)
	fmt.Println(p2)

	/*
	 * pSlice 구조체 배열을 생성했다.
	 * 구조체 배열에 구조체를 할당하면 그 내용이 배열에 복사되기 때문에 구조체 배열에 담긴
	 * 오브젝트가 변하더라도 원래 구조체의 값은 아무런 영향을 받지 않고 그대로 남아 있다.
	 */
	pSlice := [4]XYZ{}
	pSlice[2] = p1
	pSlice[0] = p2
	fmt.Println(pSlice)
	p2 = XYZ{1, 2, 3}
	fmt.Println(pSlice)
	pSlice[0] = p2
	fmt.Println(pSlice)
}
