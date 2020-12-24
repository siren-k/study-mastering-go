package main

import (
	"fmt"
	"sort"
)

/*
 * 구조체는 다양한 타입으로 표현한 여러 변수로 구성된 데이터 타입이다.
 */
type aStructure struct {
	person string
	height int
	weight int
}

func main() {
	/*
	 * mySlice란 이름의 슬라이스를 새로 정의했는데, 원소의 타입은 앞서 정의한 aStructure라는 구조체로 지정한다.
	 */
	mySlice := make([]aStructure, 0)
	mySlice = append(mySlice, aStructure{"Mihalis", 180, 90})
	mySlice = append(mySlice, aStructure{"Bill", 134, 45})
	mySlice = append(mySlice, aStructure{"Marietta", 155, 45})
	mySlice = append(mySlice, aStructure{"Epifanios", 144, 50})
	mySlice = append(mySlice, aStructure{"Athina", 134, 40})
	fmt.Println("0:", mySlice)

	// 구체적인 정렬 방법은 익명 함수로 정의했으며 aStructure 구조체의 height 필드를 사용
	// sort.Slice() 함수는 정렬 함수의 실행 결과에 따라 슬라이스에 담긴 원소의 순서를 바꾼다.
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height
	})
	fmt.Println("<:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height > mySlice[j].height
	})
	fmt.Println(">:", mySlice)
}
