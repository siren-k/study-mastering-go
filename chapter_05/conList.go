package main

import (
	"container/list"
	"fmt"
	"strconv"
)

/*
 * printList() 함수에 list.List 변수에 대한 포인터를 전달하면 그 리스트에
 * 담긴 내용을 화면에 출력한다. 코드를 보면 list.List에 담긴 변수를 첫 번째부터
 * 마지막까지 출력하고 다시 반대 방향으로 출력하고 있다. 일반적으로 두 가지 방식 중
 * 하나로 구현한다. Prev()와 Next() 함수를 사용하면 원소를 역방향과 정방향으로
 * 순환할 수 있다.
 */
func printList(l *list.List) {
	for t := l.Back(); t != nil; t = t.Prev() {
		fmt.Print(t.Value, " ")
	}
	fmt.Println()

	for t := l.Front(); t != nil; t = t.Next() {
		fmt.Print(t.Value, " ")
	}
	fmt.Println()
}

func main() {
	values := list.New()

	/*
	 * list.PushBack() 함수와 list.PushFront() 함수는 리스트에 추가한 값을 리턴한다.
	 */
	// list.PushBack() 함수는 연결 리스트의 마지막에 오브젝트를 추가한다.
	e1 := values.PushBack("One")
	e2 := values.PushBack("Two")
	// list.PushFront() 함수는 리스트의 첫 부분에 오브젝트를 추가한다.
	values.PushFront("Three")

	/*
	 * list.InsertAfter() 함수와 list.InsertBefore() 함수는
	 * 지정한 원소가 존재하지 않으면 추가하지 않는다.
	 */
	// list.InsertAfter() 함수는 새로운 원소를 특정한 원소 앞에 추가한다.
	values.InsertBefore("Four", e1)
	// list.InsertBefore() 함수는 새로운 원소를 특정한 원소 뒤에 추가한다.
	values.InsertAfter("Five", e2)

	// list.Remove() 함수는 리스트에서 특정한 원소를 삭제한다.
	values.Remove(e2)
	values.Remove(e2)
	values.InsertAfter("FiveFive", e2)

	// list.PushBackList() 함수는 다른 리스트의 앞에 기존 리스트의 복사본을 추가한다.
	values.PushBackList(values)
	printList(values)

	// list.Init() 함수는 기존 리스트를 비우거나 새 리스트를 초기화한다.
	values.Init()
	fmt.Printf("After Init(): %v\n", values)
	// for 루프를 통해 새로운 리스트를 생성한다.
	// 이 때, strconv.Itoa() 함수로 정수를 스트링으로 변환한다.
	for i := 0; i < 20; i++ {
		values.PushFront(strconv.Itoa(i))
	}
	printList(values)
}
