package main

import "fmt"

func main() {
	aMap := map[string]int{}
	aMap["test"] = 1
	fmt.Println(aMap)
	// map[test:1]

	// nil 맵에 데이터를 추가할 수 없다.
	// 그런데 nil 맵을 검색하거나 길이를 알아내거나 range 루프를 작성해도 실행이 되기는 한다.
	bMap := map[string]int{}
	bMap = nil
	fmt.Println(bMap)
	bMap["test"] = 1
	// map[]
	// panic: assignment to entry in nil map
}
