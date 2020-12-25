package main

import (
	"container/heap"
	"fmt"
)

type heapFloat32 []float32

/*
 * Pop(), Push()는 heap.Interface 인터페이스에 맞추기 위해 작성
 */
func (n *heapFloat32) Pop() interface{} {
	old := *n
	x := old[len(old)-1]
	new := old[0 : len(old)-1]
	*n = new

	return x
}

func (n *heapFloat32) Push(x interface{}) {
	*n = append(*n, x.(float32))
}

/*
 * Len(), Less(), Swap()은 sort.Interface에 필요한 세 가지 함수를 구현한다.
 */
func (n heapFloat32) Len() int {
	return len(n)
}

func (n heapFloat32) Less(a, b int) bool {
	return n[a] < n[b]
}

func (n heapFloat32) Swap(a, b int) {
	n[a], n[b] = n[b], n[a]
}

/*
 * container/heap 사용법
 * 힙(Heap)이란 일종의 트리로, 각 노드의 값은 그 노드 아래(서브트리(Subtree))에 있는
 * 원소들 중에서 가장 작은 값으로 구성한다. 여기서 주목할 부분은, 최소값(Minimum)이라 표현하지
 * 않고 가장 작은(Smallest) 값이라 표현한 점이다. 힙은 숫자가 아닌 다른 종류의 값도
 * 포함한다는 것을 강조하기 위해서다.
 *
 * 하지만 Go 언어에서 힙을 구현하려면, 원소들끼리 비교해서 어느 값이 더 작다고 판단해야 할지
 * 결정해야 한다. Go 언어에서 제공하는 인터페이스를 이용하면 이러한 동작을 정의할 수 있다.
 *
 * 다시 말해 container/heap 패키지는 container에 속한 다른 두 패키지에 비해 휠씬 고급
 * 기능을 제공한다. 그리고 container/heap 패키지에서 제공하는 기능을 사용하기 전에 미리
 * 정의해야 할 사항이 좀 있다. 구제적으로 설명하면 container/heap 패키지를 사용하려면
 * container/heap.Interface를 구현해야 한다. 해당 인터페이스는 다음과 같이 정의돼 있다.
 *
 * type Interface interface {
 *     sort.Interface
 *     Push(x interface{})  // x를 Len()이 가리키는 원소로 추가한다.
 *     Pop() interface{}    // Len() - 1에 있는 원소를 삭제하고 리턴한다.
 * }
 */
func main() {
	myHeap := &heapFloat32{1.2, 2.1, 3.1, -100.1}
	heap.Init(myHeap)
	size := len(*myHeap)
	fmt.Printf("Heap size: %d\n", size)
	fmt.Printf("%v\n", myHeap)
	// Heap size: 4
	// &[-100.1 1.2 3.1 2.1]

	/*
	 * heap.Push()를 이용하여 myHeap에 두 개의 원소를 새로 추가한다.
	 * 하지만 원소를 추가한 뒤에도 제대로 정렬된 상태를 유지하려면 heap.Init()을 호출해야 한다.
	 */
	myHeap.Push(float32(-100.2))
	myHeap.Push(float32(0.2))
	fmt.Printf("Heap size: %d\n", len(*myHeap))
	fmt.Printf("%v\n", myHeap)
	// Heap size: 6
	// &[-100.1 1.2 3.1 2.1 -100.2 0.2]

	heap.Init(myHeap)
	fmt.Printf("%v\n", myHeap)
	// &[-100.2 -100.1 0.2 2.1 1.2 3.1]  ==> 힙은 배열이나 슬라이스와 같은 선형 구조체가 아닌 일종의 트리다.
}
