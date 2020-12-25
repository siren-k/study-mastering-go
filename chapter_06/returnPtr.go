package main

import "fmt"

/*
 * y 변수에 대한 메모리 주소를 리턴할 수 있도록 return문에 &y라고 표기하였음
 */
func returnPtr(x int) *int {
	y := x * x
	return &y
}

func main() {
	sq := returnPtr(10)
	// * 문자는 '포인터 변수를 역참조' 연산을 의미한다.
	// 따라서 메모리 주소가 아닌 메모리가 주소가 가리키는 실제 값을 리턴한다.
	fmt.Println("sq:", *sq)
	fmt.Println("sq:", sq) // 메모리 주소는 실행될 때마다 달라질 수 있다.
}
