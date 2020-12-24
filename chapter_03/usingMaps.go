package main

import "fmt"

/*
 * 맵은 언제 사용해야 할까?
 * ==> 맵은 슬라이스나 배열보다 활용 범위가 훨씬 넓다. 하지만 이러한 유연함에는 치러야 할 대가가 있다.
 *     Go 언어에서 맵을 구현하는데 더 많은 프로세싱 파워가 필요하다는 것이다. 그럼에도 불구하고
 *     Go 언어에서 기본으로 제공하는 구조체들은 상당히 빠르다. 따라서 필요하다면 언제든지 주저하지
 *     말고 맵을 사용해도 된다.
 *     반드시 기억할 사항은, Go 언어의 맵은 굉장히 편리하고 다양한 종류의 데이터를 저장할 수 있는 동시에,
 *     이해하기도 쉽고 속도도 빠르다는 점이다.
 */

func main() {
	iMap := make(map[string]int)
	iMap["k1"] = 12
	iMap["k2"] = 13
	fmt.Println("iMap:", iMap)

	anotherMap := map[string]int{
		"k1": 12,
		"k2": 13,
	}
	fmt.Println("anotherMap:", anotherMap)

	// delete() 함수를 여러 차례 호출하더라도 아무런 경고 메시지 없이 똑같은 결과가 나온다.
	delete(anotherMap, "k1")
	delete(anotherMap, "k1")
	delete(anotherMap, "k1")
	fmt.Println("anotherMap:", anotherMap)

	// 지정한 키가 맵에 있는지 확인할 수 있다. 이 기법은 굉장히 중요하다. 이렇게 하지 않으면
	// 맵에 필요한 정보가 있는지 알아낼 수 없기 때문이다.
	//
	// 맵에 존재하지 않는 키 값을 구하면 0만 얻게 된다. 따라서 요청한 키가 없어서 0이 나왔는지,
	// 아니면 그 키 값이 가리키는 원소의 값이 실제로 0이어서 그런지 구분할 수 없다.
	_, ok := iMap["doesIfExist"]
	if ok {
		fmt.Println("Exists!")
	} else {
		fmt.Println("Does NOT exist")
	}

	// 코드가 휠씬 세련되고 간견해진다.
	// 맵의 원소를 출력할 때는 무작위의 순서로 나오기 때문에 정확한 순서를 알 수도 없고
	// 특정한 순서로 나온다고 가정해도 안 된다.
	for key, value := range iMap {
		fmt.Println(key, value)
	}
}
