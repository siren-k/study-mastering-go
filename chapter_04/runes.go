package main

import (
	"fmt"
)

/*
 * 룬(Rune)이란 int32에 해당되는 Go 언어에서 정식으로 제공하는 타입으로
 * 유니코드 코드 포인트(Unicode Code Point)를 표현하는데 주로 사용한다.
 * 유니코드 코드 포인트란 하나의 유니코드 문자를 표현하는 숫자 값을 의미한다.
 * 노트: 스트링은 룬의 묶음(Collection)으로 볼 수 있다.
 *
 * 룬 리터럴(Rune Literal)은 작은 따옴표로 묶은 문자 하나를 의미한다. 룬 리터럴을 '룬 상수(Rune Constant)'라고
 * 볼 수도 있다. 내부 처리 방식을 보면, 룬 리터럴 하나는 유니코드 코드 포인트 하나에 대응된다.
 */
func main() {
	// 룬 리터럴 선언
	const r1 = '€'
	/*
	 * 이러한 바이트 슬라이스를 fmt.Println()으로 출력하면 원하는 형태로 리턴하지 않을 수 있다.
	 * 룬을 문자로 변환하려면 fmt.Printf()문에서 %c를 저장해야 한다.
	 * 그리고 바이트 슬라이스를 스트링으로 출력하려면 fmt.Printf()에서 %s를 지정해야 한다.
	 */
	fmt.Println("(int32) r1:", r1)
	fmt.Printf("(HEX) r1: %x\n", r1)
	fmt.Printf("(as a String) r1: %s\n", r1)
	fmt.Printf("(as a character) r1: %c\n", r1)
	fmt.Println("A string is a collection of runes:", []byte("Mihalis"))

	// 바이트 슬라이스는 일종의 룬 묶음이다.
	aString := []byte("Mihalis")
	for x, y := range aString {
		fmt.Println(x, y)
		fmt.Printf("Char: %c\n", aString[x])
	}
	fmt.Printf("%s\n", aString)
}
