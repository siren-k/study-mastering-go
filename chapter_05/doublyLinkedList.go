package main

import "fmt"

type DoublyLinkedListNode struct {
	Value    int
	Previous *DoublyLinkedListNode
	Next     *DoublyLinkedListNode
}

var doublyLinkedListRoot = new(DoublyLinkedListNode)

func addDoublyLinkListNode(t *DoublyLinkedListNode, v int) int {
	if doublyLinkedListRoot == nil {
		t = &DoublyLinkedListNode{v, nil, nil}
		doublyLinkedListRoot = t
		return 0
	}

	if v == t.Value {
		fmt.Println("Node already exists:", v)
		return -1
	}

	if t.Next == nil {
		temp := t
		t.Next = &DoublyLinkedListNode{v, temp, nil}
		return -2
	}

	return addDoublyLinkListNode(t.Next, v)
}

func traverseDoublyLinkedList(t *DoublyLinkedListNode) {
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

func reverseDoublyLinkedList(t *DoublyLinkedListNode) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	temp := t
	for t != nil {
		temp = t
		t = t.Next
	}

	for temp.Previous != nil {
		fmt.Printf("%d -> ", temp.Value)
		temp = temp.Previous
	}
	fmt.Printf("%d -> ", temp.Value)
	fmt.Println()
}

func sizeDoublyLinkedList(t *DoublyLinkedListNode) int {
	if t == nil {
		fmt.Println("-> Empty list!")
		return 0
	}

	n := 0
	for t != nil {
		n++
		t = t.Next
	}
	return n
}

func lookupDoublyLinkedList(t *DoublyLinkedListNode, v int) bool {
	if doublyLinkedListRoot == nil {
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupDoublyLinkedList(t.Next, v)
}

/*
 * 이중 연결 리스트는 단일 연결 리스트보다 활용 범위가 넓다. 리스트를 탐색할 때 양방향 모두 이동할
 * 수 있을 뿐만 아니라, 원소를 더 쉽게 추가하거나 삭제할 수 있기 때문이다. 또한 이중 연결 리스트의
 * 헤드에 대한 포인터를 잃어버리더라도 리스트의 헤드 노드를 찾아낼 수 있다. 물론, 이렇게 활용도기
 * 높은 대신 처리야 할 대가가 있다. 바로 각 노드마다 두 개의 포인터를 관리해야 한다는 점이다. 이로
 * 인해 추가되는 오버헤드를 충분히 감수할 만한지 여부는 전적으로 개발자의 판단에 달려 있다.
 */
func main() {
	fmt.Println(doublyLinkedListRoot)
	doublyLinkedListRoot = nil
	traverseDoublyLinkedList(doublyLinkedListRoot)

	addDoublyLinkListNode(doublyLinkedListRoot, 1)
	addDoublyLinkListNode(doublyLinkedListRoot, 1)
	traverseDoublyLinkedList(doublyLinkedListRoot)

	addDoublyLinkListNode(doublyLinkedListRoot, 10)
	addDoublyLinkListNode(doublyLinkedListRoot, 5)
	addDoublyLinkListNode(doublyLinkedListRoot, 0)
	addDoublyLinkListNode(doublyLinkedListRoot, 0)
	traverseDoublyLinkedList(doublyLinkedListRoot)

	addDoublyLinkListNode(doublyLinkedListRoot, 100)
	fmt.Println("Size:", sizeDoublyLinkedList(doublyLinkedListRoot))

	traverseDoublyLinkedList(doublyLinkedListRoot)
	reverseDoublyLinkedList(doublyLinkedListRoot)
}
