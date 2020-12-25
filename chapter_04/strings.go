package main

import "fmt"

func main() {
	// \xAB 형태의 기호는 각각 sLiteral을 구성하는 문자 하나를 표현한다.
	const sLiteral = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c"
	fmt.Println(sLiteral)
	fmt.Printf("x: %x\n", sLiteral)
	fmt.Printf("sLiteral length: %d\n", len(sLiteral))

	// 스트링 리터럴을 마치 슬라이스인 것처럼 접근하고 있다.
	for i := 0; i < len(sLiteral); i++ {
		fmt.Printf("%x ", sLiteral[i])
	}
	fmt.Println()

	// %q ==> Go 언어의 문법에 따라 해석하지 않은 이스케이프(Escape)한 원래 모양대로 출력
	fmt.Printf("q: %q\n", sLiteral)
	// %+q ==> Go 언어의 문법에 따라 해석하지 않은 이스케이프(Escape)한 원래 모양 중에서 아스키(ASCII) 문자만 출력
	fmt.Printf("+q: %+q\n", sLiteral)
	// % x ==> 바이트 사이에 공백을 추가하여 출력
	fmt.Printf(" x: % x\n", sLiteral)
	// %s ==> 스트링 리터럴을 스트링으로 출력
	fmt.Printf("s: As a string: %s\n", sLiteral)

	s2 := "€£³1"
	for x, y := range s2 {
		// %#U ==> 각각의 문자를  U+0058 포맷으로 출력
		// € ==> 3 bytes
		// £ ==> 2 bytes
		// ³ ==> 2 bytes
		// 1 ==> 1 byte
		fmt.Printf("%#U starts at byte position %d\n", y, x)
	}
	// s2 변수에 유니코드 문자가 포함되어 있어 바이트 단위로 표현한 s2의 길이는 8이다.
	fmt.Printf("s2 length: %d\n", len(s2))

	const s3 = "ab12AB"
	fmt.Println("s3:", s3)
	fmt.Printf("x: % x\n", s3)
	fmt.Printf("s3 length: %d\n", len(s3))
	for i := 0; i < len(s3); i++ {
		fmt.Printf("%x ", s3[i])
	}
	fmt.Println()
}
