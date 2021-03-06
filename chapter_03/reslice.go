package main

import (
	"fmt"
)

func main() {
	// [:]는 기존 슬라이스나 배열을 이용하여 새로운 슬라이스를 생성할 때도 사용할 수 있다.
	// 이를 리슬라이싱(Re-Slicing)이라 부른다.
	// 리슬라이스는 다음과 같은 이슈를 가지고 있다.
	//   1) 리슬라이싱한 슬라이스의 원소를 변경하면 원본 슬라이스의 원소도 변경된다.
	//      ==> 두 슬라이스가 동일한 메모리 주소를 가리키고 있기 때문이다. 리슬라이싱은 원본 슬라이스에 대한 복사본을 만드는 것이 아니다.
	//   2) 어떤 슬라이스의 아주 작은 부분만 리슬라이싱하더라도 원본 슬라이스에 대한 내부 배열은 리슬라이싱한 부분이 존재하는한 메모리에
	//      계속 남아 있데 된다.
	//      ==> 원본 슬라이스가 작다면 문제가 되지 않지만, 대용량 파일을 슬라리스로 불러와서 그 중 일부분만 사용하고 싶을 때는 문제가 발생할 수 있다.
	s1 := make([]int, 5)
	reSlice := s1[1:3]
	fmt.Println(s1)
	fmt.Println(reSlice)

	reSlice[0] = -100
	reSlice[1] = 123456
	fmt.Println(s1)
	fmt.Println(reSlice)
}
