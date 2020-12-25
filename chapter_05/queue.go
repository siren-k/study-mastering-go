package main

import "fmt"

type QueueNode struct {
	Value int
	Next  *QueueNode
}

var queueSize = 0
var queue = new(QueueNode)

func Push(t *QueueNode, v int) bool {
	if queue == nil {
		queue = &QueueNode{v, nil}
		queueSize++
		return true
	}

	t = &QueueNode{v, nil}
	t.Next = queue
	queue = t

	queueSize++

	return true
}

func Pop(t *QueueNode) (int, bool) {
	if queueSize == 0 {
		return 0, false
	}

	if queueSize == 1 {
		queue = nil
		queueSize--
		return t.Value, true
	}

	temp := t
	for t.Next != nil {
		temp = t
		t = t.Next
	}

	v := temp.Next.Value
	temp.Next = nil

	queueSize--

	return v, true
}

func traverseQueue(t *QueueNode) {
	if queueSize == 0 {
		fmt.Println("Empty Queue!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func main() {
	queue = nil
	Push(queue, 10)
	fmt.Println("Size:", queueSize)
	traverseQueue(queue)

	v, b := Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", queueSize)

	for i := 0; i < 5; i++ {
		Push(queue, i)
	}
	traverseQueue(queue)
	fmt.Println("Size:", queueSize)

	v, b = Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", queueSize)

	v, b = Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", queueSize)
	traverseQueue(queue)
}
