package main

import (
	"fmt"
)

func main() {
	a6 := []int{-10, 1, 2, 3, 4, 5}
	a4 := []int{-1, -2, -3, -4}
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	// copy() 함수를 이용하면 기존 배열의 원소로부터 슬라이스를 생성할 수도 있고,
	// 기존 슬라이에서 다른 슬라이스로 복사할 수도 있다. 기본 제공되는 copy(dst, src) 함수는
	// len(dst)와 len(src)의 최소값에 해당되는 수만큼의 원소만 복사한다.
	copy(a6, a4)
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	fmt.Println()
	// a6: [-10 1 2 3 4 5]
	// a4: [-1 -2 -3 -4]
	// a6: [-1 -2 -3 -4 4 5] ==> a6의 원소의 수가 a4보다 더 많기 때문에 a4의 모든 원소가 a6로 복사되지만 a6의 마지막 두 원소는 그래도 남아 있다.
	// a4: [-1 -2 -3 -4]

	b6 := []int{-10, 1, 2, 3, 4, 5}
	b4 := []int{-1, -2, -3, -4}
	fmt.Println("b6:", b6)
	fmt.Println("b4:", b4)
	copy(b4, b6)
	fmt.Println("b6:", b6)
	fmt.Println("b4:", b4)
	fmt.Println()
	// b6: [-10 1 2 3 4 5]
	// b4: [-1 -2 -3 -4]
	// b6: [-10 1 2 3 4 5]
	// b4: [-10 1 2 3] ==> b4의 원소가 4개이기 때문에 b6에서 첫 네 원소만 복사되었다.

	array4 := [4]int{4, -4, 4, -4}
	s6 := []int{1, 1, -1, -1, 5, -5}
	// array4[0:] ==> [:] 기호를 통해 배열이 슬라이스로 변환된다.
	// copy()는 슬라이스에 대한 인수만 받기 때문에 배열을 슬라이스로 변환해야 한다.
	copy(s6, array4[0:])
	fmt.Println("array4:", array4[0:])
	fmt.Println("s6:", s6)
	fmt.Println()

	array5 := [5]int{5, -5, 5, -5, 5}
	s7 := []int{7, 7, -7, -7, 7, -7, 7}
	copy(array5[0:], s7)
	fmt.Println("array5:", array5)
	fmt.Println("s7:", s7)
}
