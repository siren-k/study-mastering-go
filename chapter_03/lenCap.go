package main

import (
	"fmt"
)

func printSlice(x []int) {
	for _, number := range x {
		fmt.Print(number, " ")
	}
	fmt.Println()
}

func main() {
	// 슬라이스 리터럴(Slice Literal)을 정의하는 방법은 원소의 수를 지정하는 부분만 없을 뿐, 배열과 거의 같다.
	// 여기서 배열의 크기를 지정하면 슬라이스가 아닌 배열이 생성된다.
	//
	// make() 함수에 원하는 길이(Length)와 용량(Capacity)을 매개변수로 전달하면 빈 슬라이스를 생성할 수 있다.
	// 이 때, 용량에 대한 매개변수는 생략할 수 있다. 그러면 슬라이스의 용량은 길이와 같게 된다
	//
	// 20개를 저장할 수 있으면서 자동으로 확장되는 빈 슬라이스를 정의하려면 다음과 같이 작성한다.
	//   integer := make([]int, 20)
	// 초기화에 사용하는 값은 슬라이스에 저장하는 오브젝트의 타입에 따라 결정된다.
	//
	// 기존 슬라이스를 비우고 싶다면 슬라이스 변수의 값을 nil로 지정하면 된다.
	//
	// append() 함수를 이용해 슬라이스에 원소를 추가할 수도 있다. 그러면 슬라이스 크기가 자동으로 증가한다.
	//   integer = append(integer, -5000)
	//
	// integer의 첫 번째 원소는 integer[0]으로 접근할 수 있고,
	// integer 슬라이스의 마지막 원소는 integer[len(integer) - 1]로 접근할 수 있다.
	//
	// 또한 [:]을 사용해 슬라이스에 연속으로 나와 있는 원소들을 접근할 수 있다. 예를 들어 슬라이스에서
	// 두 번째 원소부터 세 번째 원소를 선택하려면 다음과 같이 작성한다.
	// integer[1:3]
	//
	// [:]는 기존 슬라이스나 배열을 이용하여 새로운 슬라이스를 생성할 때도 사용할 수 있다.
	// 이를 리슬라이싱(Re-Slicing)이라 부른다.
	// 리슬라이스는 다음과 같은 이슈를 가지고 있다.
	//   1) 리슬라이싱한 슬라이스의 원소를 변경하면 원본 슬라이스의 원소도 변경된다.
	//      ==> 두 슬라이스가 동일한 메모리 주소를 가리키고 있기 때문이다. 리슬라이싱은 원본 슬라이스에 대한 복사본을 만드는 것이 아니다.
	//   2) 어떤 슬라이스의 아주 작은 부분만 리슬라이싱하더라도 원본 슬라이스에 대한 내부 배열은 리슬라이싱한 부분이 존재하는한 메모리에
	//      계속 남아 있데 된다.
	//      ==> 원본 슬라이스가 작다면 문제가 되지 않지만, 대용량 파일을 슬라리스로 불러와서 그 중 일부분만 사용하고 싶을 때는 문제가 발생할 수 있다.
	//
	// 슬라이스는 두 가지 속성, 즉 용량(Capacity)와 길이(Length)를 가진다.
	// 슬라이스의 길이는 배열의 길이와 같으며, len() 함수로 알아낼 수 있다.
	// 반면 슬라이스 용량은 현재 이 슬라이스에 할당된 공간을 의마하며, cap() 함수로 알아낼 수 있다.
	// 슬라이스의 용량은 동적으로 변한다. 슬라이스에 공간이 부족하면 Go 언어는 현재 길이를 자동으로 2배로 확자해 더 많은 원소를 담을 공간을 확보한다.
	// ==> 엄청나게 거대한 슬라이스라면 원소 하나를 추가하더라도 예상보다 훨씬 많은 메모리 공간을 차지할 수 있다.
	aSlice := []int{-1, 0, 4}
	fmt.Printf("sSlice: ")
	printSlice(aSlice)
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))

	aSlice = append(aSlice, -100)
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))

	aSlice = append(aSlice, -2)
	aSlice = append(aSlice, -3)
	aSlice = append(aSlice, -4)
	printSlice(aSlice)
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))
}
