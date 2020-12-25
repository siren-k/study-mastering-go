package main

import "fmt"

/*
 * 익명 함수를 리턴
 */
func funReturnFun() func() int {
	i := 0
	return func() int {
		i++
		return i * i
	}
}

func main() {
	/*
	 * i와 j가 모두 funReturnFun()이란 함수로부터 생성됐지만, 서로 완전히 독립적이며
	 * 공통된 부분이 하나도 없다. 따라서 리턴 값이 같은 코드에서 나올 뿐, 서로 완전히
	 * 별개로 실행된다.
	 */
	i := funReturnFun()
	j := funReturnFun()

	fmt.Println("i1:", i())
	fmt.Println("i2:", i())
	fmt.Println("j1:", j())
	fmt.Println("j2:", j())
	fmt.Println("i3:", i())
}
