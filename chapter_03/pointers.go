package main

import "fmt"

/*
 * 포인터를 매개변수로 전달하도록 작성하면, 이 함수에 변수를 전달하는 것만으로도
 * 변수를 업데이터할 수 있다. 이 함수를 호출한 코드로 결과를 리턴하지 않아도 된다.
 * 이렇게 할 수 있는 이유는 바로 매개변수가 전달한 포인터에 업데이트할 변수의 메모리 주소가 담겨 있기 때문다.
 */
func getPointer(n *int) {
	*n = *n * *n
}

/*
 * 정수형 매개변수를 받아서 정수값에 대한 포인터를 리턴한다. 이를 return &v라고 표현했다.
 * 얼핏보면 유용하지 않아 보이자만 Go 구조체에 대한 포인터에 대해 설명할 때와 이후 좀 더 복잡한
 * 구조체를 다룰 때 그 진가를 제대로 알 수 있다.
 */
func returnPointer(n int) *int {
	v := n * n
	return &v
}

func main() {
	/*
	 * i와 j는 모두 일반 정수형 변수이다.
	 */
	i := -10
	j := 25

	/*
	 * 이에 반해 pI와 pJ는 각각 i와 j를 가리키는 포인터이다.
	 */
	pI := &i
	pJ := &j

	/*
	 * pI는 포인터의 메모리 주소인 반면, *pI는 그 주소에 저장된 값이다.
	 */
	fmt.Println("pI memory:", pI)
	fmt.Println("pJ memory:", pJ)
	fmt.Println("pI value:", *pI)
	fmt.Println("pJ value:", *pJ)

	/*
	 * 변수 i를 가리키는 포인터인 pI를 통해 i 값을 변경
	 * 1) 새로운 값을 직접 할당
	 * 2) -- 연사자를 사용
	 */
	*pI = 123456
	*pI--
	fmt.Println("i:", i)

	/*
	 * getPointer() 함수 안에서 pJ 변수를 변경하면 변수 j의 값에 영향을 미친다.
	 * pJ 변수가 j 변수를 가리키고 있기 때문이다.
	 */
	getPointer(pJ)
	fmt.Println("j:", j)

	k := returnPointer(12)
	fmt.Println(*k)
	fmt.Println(k)

	// C 언어에서 스트링은 포인터 타입인데 반해, Go 언어에서는 값 타입(Value Type)이다
}
