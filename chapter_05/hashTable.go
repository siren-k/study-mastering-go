package main

import "fmt"

// 해시 테이블의 버킷 수
const SIZE = 15

type HashTableNode struct {
	Value int
	Next  *HashTableNode
}

/*
 * 해시 테이블의 실제 내용은 두 개의 필드로 구성된 HashTable 구조체에 저장된다. 두 번째 필드는
 * 해시 테이블의 크기를 가리키고, 첫 번째 필드는 정수에 대한 연결 리스트(*HashTableNode)에 대한 맵이다.
 * 따라서 여기서 만든 해시 테이블에서 연결 리스트는 버킷의 수만큼 있다. 이 말은 해시 테이블의 각
 * 버킷에 있는 노드는 연결 리스에 저장된다는 것을 의미한다.
 */
type HashTable struct {
	Table map[int]*HashTableNode
	Size  int
}

/*
 * 모듈로 연산자(Modulo Operator)를 사용한다.
 * 이렇게 모듈로 연상으로 해시 함수를 구현한 가장 큰 이유는 여기서 만들 해시 테이블이
 * 정수형 값을 다루기 때문이다. 스트링이나 부동 소수점 수를 다룰 때는 다른 방법을
 * 사용해야 한다.
 */
func hashFunction(i, size int) int {
	return i % size
}

/*
 * 해시 테이블에 원소를 추가할 때 호출된다.
 * 여기서는 추가할 값이 기존에 있는지 아닌지는 검사하지 않았다.
 */
func insertHash(hash *HashTable, value int) int {
	index := hashFunction(value, hash.Size)
	element := HashTableNode{Value: value, Next: hash.Table[index]}
	hash.Table[index] = &element
	return index
}

/*
 * 해시 테이블에 담긴 모든 값을 화면에 출력한다. 이 힘수는 해시 테이블에 있는
 * 연결 리스트에 담긴 모든 원소를 방문해서 그 값을 화면에 출력한다. 이 작업은
 * 연결 리스트 단위로 처리한다.
 */
func traverseHash(hash *HashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%d -> ", t.Value)
				t = t.Next
			}
			fmt.Println()
		}
	}
}

func lookupHash(hash *HashTable, value int) bool {
	index := hashFunction(value, hash.Size)
	if hash.Table[index] != nil {
		t := hash.Table[index]
		for t != nil {
			if t.Value == value {
				return true
			}
			t = t.Next
		}
	}
	return false
}

func main() {
	// 해시 테이블의 버킷을 담은 맵에 대한 변수인 table을 이용해 hash란 이름의 해시 테이블을
	// 새로 생성했다. 해시 테이블은 슬롯(버킷)은 연결 리스트로 구현했다. 해시 테이블의 연결 리스트를
	// 슬라이스나 배열이 아닌 맵으로 저장한 가장 큰 이유는 슬라이스나 배열에서는 키 값이 반드시 양의
	// 정수여야 하지만, 맵에서는 어떠한 타입의 값도 키로 사용할 수 있기 때문이다.
	table := make(map[int]*HashTableNode, SIZE)
	hash := &HashTable{Table: table, Size: SIZE}
	fmt.Println("Number of spaces:", hash.Size)

	for i := 0; i < 120; i++ {
		insertHash(hash, i)
	}
	traverseHash(hash)
}
