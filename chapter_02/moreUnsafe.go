package main

import (
	"fmt"
	"unsafe"
)

func main() {
	array := [...]int{0, 1, -2, 3, 4}
	// 정수형 배열의 첫 번째 원소인 array[0]에 대한 메모리 주소를 가리킴
	pointer := &array[0]
	// pointer를 역참조할 때는 *pointer와 같이 표기함
	// 그러면 그 주소에 저장되어 있던 정수 값을 리턴함
	fmt.Print(*pointer, " ")
	// pointer 변수를 unsafe.Pointer() 함수로 변환하고 다시 uintptr로 변환하여 결과를 memoryAddress에 저장함
	// unsafe.Sizeof(array[0])의 값을 더하는 방식으로 이 배열의 그 다음 원소를 구할 수 있음
	// 배열의 각 원소가 메모리에서 차지하는 공간은 모두 같기 때문임
	memoryAddress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])

	for i := 0; i < len(array)-1; i++ {
		pointer = (*int)(unsafe.Pointer(memoryAddress))
		fmt.Print(*pointer, " ")
		memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	}
	fmt.Println()
	pointer = (*int)(unsafe.Pointer(memoryAddress))

	// 지정한 포인터와 메모리 주소에는 존재하지 않는 배열의 원소에 접근하고 있지만
	// unsafe 패키지를 사용하고 있기 때문에 Go 컴퍼알러에서 이런 논리적 에러를 찾아 주지 않음
	fmt.Println("One More: ", *pointer, " ")
	memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	fmt.Println()
}
