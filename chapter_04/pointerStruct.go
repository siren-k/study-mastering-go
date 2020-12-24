package main

import "fmt"

type myStructure struct {
	Name    string
	Surname string
	Height  int32
}

/*
 * 새로운 구조체 변수를 생성할 때 원하는 값으로 초기화하는 과정에서 여러 가지 장점이 있다.
 * 가령, 입력한 값이 정확하고 유효한지 확인할 수 있다. 또한 코드가 휠씬 깔끔해진다.
 * 구조체 변수에 대해 초기화한 곳에서 처리하기 때문에 struct 변수에 문제가 발생할 때 어디서
 * 일어났는지, 무엇 때문에 발생했는지 명확히 알 수 있다. 참고로 createStruct() 대신
 * NewStruct()란 이름을 선호하는 사람도 있다.
 *
 * C나 C++에 익숙한 독자를 위해 한 가지 알려주면 Go 언어의 함수에서 로컬 변수에 대한 메모리
 * 주소를 리턴하는 기능도 정식으로 지원한다. 따라서 어느 누구도 아쉬울 것 없이 모두가 만족할 수 있다.
 */
func createStruct(n, s string, h int32) *myStructure {
	if h > 300 {
		h = 0
	}
	return &myStructure{n, s, h}
}

func retStructure(n, s string, h int32) myStructure {
	if h > 300 {
		h = 0
	}
	return myStructure{n, s, h}
}

func main() {
	s1 := createStruct("Mihalis", "Tsoukalos", 123)
	s2 := retStructure("Mihalis", "Tsoukalos", 123)
	fmt.Println((*s1).Name) // 구조체 포인터가 가리키는 오브젝트를 사용하려면 리턴된 포인터를 역참조해야 한다.
	// 이로 인해 코드가 좀 지저분해 보일 수 있다.
	fmt.Println(s2.Name)
	fmt.Println(s1)
	fmt.Println(s2)

	/*
	 * new와 make의 가장 큰 차이점은 make로 생성한 변수는 정상적으로 초기화된 반면, new로 생성된 변수는
	 * 할당된 메모리 공간에 단지 0만 채운다는 점이다. 또한 make는 맵, 채널, 슬라이스에만 적용할 수 있으며,
	 * 메모리 주소를 리턴하지 않는다. 다시 말해 make는 포인터를 리턴하지 않는다.
	 */
	pS := new(myStructure)
	fmt.Println(pS)
	pS.Name = "[Name]"
	pS.Surname = "[Surname ]"
	pS.Height = 100
	fmt.Println("pS:", pS)
	fmt.Println("pS.Name:", pS.Name)
	fmt.Println("(*pS).Name:", (*pS).Name)
	fmt.Println("&(pS):", &(pS))
	fmt.Println("&(pS).Name", &(pS).Name)
	fmt.Println("&(pS).Surname", &(pS).Surname)
	fmt.Println("&(pS).Height", &(pS).Height)
}
