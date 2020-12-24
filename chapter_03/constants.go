package main

import (
	"fmt"
)

/*
 * 2개의 타입(Digit과 Power2)과 4개의 상수(PI, C1, C2, C3)를 선언
 * type 키워드는 타입을 다른 이름으로 정의할 때(Named Type) 사용하며,
 * 실제 타입은 기존에 제공되는 내부 타입을 그대로 사용한다. 이렇게 하는 주된 이유는
 * 같은 종류의 데이터를 다른 타입으로 구분하기 위해서다
 */
type Digit int
type Power2 int

const PI = 3.1415926
const (
	C1 = "C1C1C1"
	C2 = "C2C2C2"
	C3 = "C3C3C3"
)

func main() {
	const s1 = 123
	var v1 float32 = s1 * 12
	fmt.Println(v1)
	fmt.Println(PI)

	const (
		Zero Digit = iota // Digit 타입의 상수를 연속으로 정의하도록 '상수 생성자(iota)'를 지정했다.
		One
		Two
		Three
		Four
	)
	fmt.Println(One)
	fmt.Println(Two)

	const (
		p2_0 Power2 = 1 << iota // 1 << 0 ==> 1
		_                       // i << 1 ==> 2 불필요한 값을 건너뛸 수 있다. 두 번째로 iota의 값은 항상 증가한다.
		p2_2                    // 1 << 2 ==> 4
		_                       // 1 << 3 ==> 8
		p2_4                    // 1 << 4 ==> 16
		_                       // 1 << 5 ==> 32
		p2_6                    // 1 << 6 ==> 64
	)
	fmt.Println("2^0", p2_0)
	fmt.Println("2^2", p2_2)
	fmt.Println("2^4", p2_4)
	fmt.Println("2^6", p2_6)
}
