package main

import "fmt"

func func1() int {
	fmt.Println("Entering func1()")
	return -10
	fmt.Println("Exiting func1()")
	return -1
}

func func2() int {
	if true {
		return 10
	}
	fmt.Println("Exiting func2()")
	return 0
}

/*
 * ❯ go vet cannotReach.go
 * command-line-arguments
 * ./cannotReach.go:8:5: unreachable code ==> 8번째 줄에 실행되지 않는 코드가 있다고 출력된다.
 */
func main() {
	fmt.Println(func1())
	fmt.Println("Exiting program...")
}
