package main

import "fmt"

func function1(i int) int {
	return i + i
}

func function2(i int) int {
	return i * i
}

/*
 * funFun() 함수는 두 개의 매개변수를 받는데, 하나는 함수 매개변수인 f고, 다른 하나는
 * int 값인 v를 받는다. f 매개변수는 한 개의 int 인수를 받는 함수로서, int 값을 리턴한다.
 */
func funFun(f func(int) int, v int) int {
	return f(v)
}

func main() {
	// 첫 번째로 호출한 fmt.Println()은 funFun()의 결과를 출력하는데, 여기서 funFun의 매개변수로 전달하는
	// function1에 소괄호를 적지 않았다.
	fmt.Println("function1:", funFun(function1, 123))
	// 두 번째로 호출한 fmt.Println()에서는 funFun()의 매개변수로 function2를 지정했다.
	fmt.Println("function2:", funFun(function2, 123))
	// 마지막 세 번째로 호출한 fmt.Println()을 보면, 특이하게도 funFun()을 호출할 때 전달할 함수 매개변수를
	// 여기서 정의했다. 함수 매개변수가 작고 간단하다면 이렇게 작성해도 되지만, 여러 줄로 구성된 함수를
	// 매개변수로 전달할 때는 그리 좋은 방법은 아니다.
	fmt.Println("Inline:", funFun(func(i int) int { return i * i * i }, 123))
}
