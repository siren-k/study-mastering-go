package main

import "fmt"

func main() {
	/*
	 * 2개의 슬라이스를 정의
	 */
	aSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(aSlice)
	integer := make([]int, 2)
	fmt.Println(integer)
	integer = nil
	// 이차원 슬라이스의 오브젝트들이 nil로 초기화돼 아무것도 출력되지 않는다.
	// 슬라이스 타입에서 0에 대한되는 값이 nil이기 때문이다.
	fmt.Println(integer)

	/*
	 * [:] 기호를 이용해 기존 배열을 참조하는 슬라이스를 새로 생성한다.
	 * 이 때 배열의 복사본이 생성되는 것이 아니라, 단지 배열을 참조하기만 한다.
	 */
	anArray := [5]int{-1, -2, -3, -4, -5}
	refAnArray := anArray[:]

	fmt.Println(anArray)
	fmt.Println(refAnArray)
	anArray[4] = -100
	fmt.Println(refAnArray)

	/*
	 * make() 함수를 이용해 일차원 슬라이스 한 개와 이차원 슬라이스 한 개를 생성한다.
	 */
	s := make([]byte, 5)
	fmt.Println(s)
	// 슬라이스는 Go 언어에서 자동으로 초기화해주기 때문에 앞에서 정의한 두 슬라이스에 담긴
	// 모든 원소는 슬라이스 타입에서 0에 해당되는 값(정수형 값은 0으로, 슬라이스에 대해서는 nil)으로
	// 초기화된다.
	twoD := make([][]int, 3)
	fmt.Println(twoD)
	fmt.Println()

	/*
	 * 이차원 슬라이스의 원소를 직접 초기화하는 방법
	 */
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < 2; j++ {
			// 기존 슬라이스를 더 크게 확장하려면, append() 함수를 사용해야 하며, 존재하지 않는 인덱스는 참조하면 안된다.
			// 없는 원소를 참조하면 panic: runtime error: index out of range라는 에러 메시지가 발생한다.
			twoD[i] = append(twoD[i], i*j)
		}
	}

	/*
	 * range 키워드를 이용하여 이차원 슬라이스의 모든 원소를 출력하도록 작성했다.
	 */
	for _, x := range twoD {
		for i, y := range x {
			fmt.Println("i:", i, "value", y)
		}
	}
}
