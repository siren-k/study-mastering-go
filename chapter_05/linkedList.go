package main

import "fmt"

type LinkedListNode struct {
	Value int
	Next  *LinkedListNode
}

// root를 전역 변수로 정의했기 때문에 코드의 어디서나 접근할 수 있다.
var linkedListRoot = new(LinkedListNode)

// 연결 리스트는 원래 중복된 항목을 저장하지 않는다. 따라서 코드도 그렇게 구현했다. 또한
// 연결 리스트가 정렬돼 있지 않을 때는 새 노드를 리스트의 마지막에 추가한다.
func addLinkedListNode(t *LinkedListNode, v int) int {
	/*
	 * 연결 리스트가 비어 있는지를 검사한다.
	 */
	if linkedListRoot == nil {
		t = &LinkedListNode{v, nil}
		linkedListRoot = t
		return 0
	}

	/*
	 * 추가하려는 값이 이미 리스트에 있는지 확인한다.
	 */
	if v == t.Value {
		fmt.Println("LinkedListNode already exists:", v)
		return -1
	}

	/*
	 * 연결 리스트의 끝에 도달했는지 확인한다. 이 상태에서 노드를 새로 추가하면
	 * t.Next = &LinkedListNode{v, nil}이란 문장을 통해 리스트의 끝에 추가된다.
	 */
	if t.Next == nil {
		t.Next = &LinkedListNode{v, nil}
		return -2
	}

	/*
	 * 세 가지 조건 중에서 어느 하나도 걸리지 않으면 return addLinkedListNode(t.Next, v)가 실행되면서
	 * 연결 리스트의 다음 노드에 대해 방금 수행한 addLinkedListNode() 함수의 동작을 다시 반복한다.
	 */
	return addLinkedListNode(t.Next, v)
}

func traverseLinkList(t *LinkedListNode) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

/*
 * 주어진 원소가 연결 리스트에 존재하는지를 확인
 * 연결 리스트에서 원하는 값을 찾기 위해 모든 원소에 접근한다. 원하는 값을 찾지 못한 채
 * 연결 리스트의 꼬리에 도달하면 연결 리스트에 그 값이 없다고 볼 수 있다.
 */
func lookupLinkedListNode(t *LinkedListNode, v int) bool {
	if linkedListRoot == nil {
		t = &LinkedListNode{v, nil}
		linkedListRoot = t
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupLinkedListNode(t.Next, v)
}

/*
 * 연결 리스트의 크기를 리턴
 */
func sizeLinkedList(t *LinkedListNode) int {
	if t == nil {
		fmt.Println("-> Empty list!")
		return 0
	}

	i := 0
	for t != nil {
		i++
		t = t.Next
	}

	return i
}

/*
 * 연결 리스트의 가장 큰 장점은 이해하고 구현하기 쉽고 활용 범위가 상당히 넓다는 것이다. 따라서
 * 굉장히 다양한 종류의 데이터에 대한 모델을 만드는 데 사용할 수 있다. 또한 연결 리스트에서
 * 포인터를 이용하면 순차 탐색(Sequential Searching)을 굉장히 빠르게 처리할 수 있다.
 *
 * 연결 리스트는 데이터를 정렬만 할 때는 큰 장점이 없지만, 원소를 추가하거나 삭제한 후에도 정렬
 * 상태를 유지하는 데는 뛰어나다. 정렬된 상태의 연결 리스트에서 노드를 삭제하는 방법은 정렬되지 않은
 * 연결 리스트에 대해 수행할 때와 같다. 하지만 정렬된 상태의 연결 리스트에 새로운 노드를 추가하는
 * 방법은 다르다. 그 이유는 연결 리스트의 정렬 상태를 유지하려면 새로운 노드를 적절한 위치에 집어
 * 넣어야 하기 때문이다. 실제로 데이터가 많이 있으면서 항상 데이터를 삭제해야 할 때는 해시 테이블이나
 * 이진 트리보다 연결 리스트를 이용하는 편이 낫다.
 *
 * 마지막으로 정렬된 연결 리스트(Sorted Linked List)를 사용하면 노드를 탐색하거나 추가하는 과정에서
 * 다양한 최적화 기법을 적용할 수 있다. 가장 흔히 사용하는 방법은 정렬된 연결 리스트의 중심 노드에
 * 포인터를 저장하고 있다가 탐색을 수행할 때 그 지점에서 시작한다. 이렇게 간단히 최적화하는 것만으로도
 * 탐색 연산의 수행 시간을 절반으로 단축시킬 수 있다.
 */
func main() {
	fmt.Println(linkedListRoot)
	linkedListRoot = nil
	traverseLinkList(linkedListRoot)

	addLinkedListNode(linkedListRoot, 1)
	addLinkedListNode(linkedListRoot, -1)
	traverseLinkList(linkedListRoot)

	addLinkedListNode(linkedListRoot, 10)
	addLinkedListNode(linkedListRoot, 5)
	addLinkedListNode(linkedListRoot, 45)
	addLinkedListNode(linkedListRoot, 5)
	addLinkedListNode(linkedListRoot, 5)
	traverseLinkList(linkedListRoot)

	addLinkedListNode(linkedListRoot, 100)
	traverseLinkList(linkedListRoot)

	if lookupLinkedListNode(linkedListRoot, 100) {
		fmt.Println("Node exists!")
	} else {
		fmt.Println("Node does not exist!")
	}

	if lookupLinkedListNode(linkedListRoot, -100) {
		fmt.Println("Node exists!")
	} else {
		fmt.Println("Node does not exist!")
	}
}
