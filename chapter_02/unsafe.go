package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var value int64 = 5
	var p1 = &value
	var p2 = (*int32)(unsafe.Pointer(p1)) // int32 타입 포인터를 생성
	// value란 이름의 int64 타입의 변수를 가르킴
	// Go 언어에서는 모든 포인터를 unsafe.Pointer로 변환할 수 있음
	// unsafe.Pointer 타압의 포인터는 Go 언어의 타입 시스템을 무시할 수 있다.
	// 그러면 당연히 속도는 빠르지만 위험한 상황이 발생할 수 있다. 또한 이 기능을
	// 이용하면 개발자는 데이터에 대해 더 많은 제어권을 가질 수 있다.

	fmt.Println("*p1:", *p1)
	fmt.Println("*p2:", *p2)
	*p1 = 5434123412312431212
	fmt.Println(value)
	fmt.Println("*p2:", *p2) // 32비트 포인터로는 64비트 정수를 담을 수 없음
	*p1 = 54341234
	fmt.Println(value)
	fmt.Println("*p2:", *p2)
}
