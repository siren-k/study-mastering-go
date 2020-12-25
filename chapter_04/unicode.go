package main

import (
	"fmt"
	"unicode"
)

func main() {
	const sL = "\x99\x00ab\x50\x00\x23\x50\x29\x9c"
	for i := 0; i < len(sL); i++ {
		// unicode.IsPrint() 함수는 주어진 룬을 출력할 수 있다면 true를 그렇지 않다면 false를 리턴한다.
		if unicode.IsPrint(rune(sL[i])) {
			fmt.Printf("%c\n", sL[i])
		} else {
			fmt.Println("Not Printable!")
		}
	}
}
