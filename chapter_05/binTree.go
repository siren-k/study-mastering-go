package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

/*
 * 이진 트리를 구성하는 모든 노드를 재귀 호출(Recursion) 방식으로 방문
 */
func traverse(t *Tree) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Print(t.Value, " ")
	traverse(t.Right)
}

/*
 * 이진 트리를 정수형 난수로 채움
 */
func create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++ {
		temp := rand.Intn(n * 2)
		t = insert(t, temp)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	/*
	 * 트리가 비었는지 검사한다.
	 * 트리가 비어 있다면 새로 생성한 노드가 트리의 루트가 되며 &Tree{nil, v, nil}로 생성한다.
	 */
	if t == nil {
		return &Tree{nil, v, nil}
	}

	/*
	 * 트리에 추가하려는 값이 이미 트리에 들어 있는지 검사한다.
	 * 이미 존재한다면, 아무 일도 하지 않고 리턴한다.
	 */
	if v == t.Value {
		return t
	}

	/*
	 * 현재 노드의 상태에 따라 추가할 값을 노드의 왼쪽에 넣을지, 아니면 오른쪽에 넣을지 판단한다.
	 */
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}

	t.Right = insert(t.Right, v)

	return t
}

/*
 * 이진 트리의 장점
 * 계층형(Hierarchical) 데이터를 표현하는데 트리만큼 좋은 것은 없다. 그래서 프로그래밍 언어의 컴파일러에서
 * 프로그램을 파싱할 때 트리를 굉장히 많이 사용한다.
 *
 * 또한, 트리는 기본적으로 순서를 가진다. 다시 말해 값을 정렬하기 위한 작업을 따로 수행할 필요가 없다. 원소를
 * 적절한 위치에 추가하는 과정이 곧 정렬하는 과정이다. 하지만 원소를 삭제하는 것은 트리의 구성 방식에 따라
 * 쉽지 않을 수 있다.
 *
 * 이진 트리의 균형이 잡여 있다면, 검색, 추가, 삭제 연산의 속도는 대략 log(n)이다. 여기서 n은 트리에 있는
 * 노드의 수를 가리킨다. 또한 균형 이진 트리의 높이는 대략 log2(n)이다. 다시 말해 원소의 수가 10,000인
 * 균형 트리의 높이는 대략 14다. 이는 상당히 작은 편에 속한다. 마찬가지로 원소의 주가 1,000,000인 균형 트리의
 * 높이는 대략 20이다. 다시 말새 균형 이진 트리에 엄청난 수의 원소를 추가하더라도 속도는 크게 변하지 않는다.
 * 다르게 표현하면, 1,000,000 개의 노드로 구성된 트리에서 특정한 노드에 도달하는데 걸리는 시간이 많아야 20이란
 * 뜻이다.
 *
 * 이진 트리의 가장 큰 단점은 트리의 모양이 원소를 추가하는 순서에 따라 달라진다는 것이다. 트리에서 키 값이 길고
 * 복잡하면 비교할 대상이 많아지기 때문에, 원소를 추가하거나 탐색하는 시간이 길어진다. 마지막으로 트리의 균형이 잡혀
 * 있지 않으면 성능을 예측할 수 없다.
 *
 * 이진 트리보다 연결 리스트나 배열이 훨씬 빠르지만, 검색 연산에 대한 유연성은 이러한 성능 및 관리의 오버헤드를
 * 상쇄하고도 남는다. 이진 트리에서 특정한 원소를 탐색할 때, 그 원소의 값이 현재 노드보다 크거나 작은지 검사할 수
 * 있고, 이러한 결과를 이용해 트리를 탐색할 때 어느 방향으로 가야 하는지 결정할 수 있다. 이러한 점은 전반적인
 * 탐색 시간을 크게 절약해 준다.
 */
func main() {
	tree := create(10)
	fmt.Println("The value of the root of the tree is", tree.Value)

	traverse(tree)
	fmt.Println()

	tree = insert(tree, -10)
	tree = insert(tree, 2)
	traverse(tree)
	fmt.Println()

	fmt.Println("The value of the root of the tree is", tree.Value)
}
