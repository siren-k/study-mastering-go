package main

import (
	/*
	 * 플래그(Flag)란 프로그램의 동작을 제어하기 위해 전달하는 옵션으로, 특정한 포맷의 스트링으로 표현한다.
	 * 여러 가지 플래그를 지원하도록 직접 구현하기란 간단하지 않다. flag 패키지가 없어도 유닉스 시스템용
	 * 커맨드라인 유틸리티를 충분히 개발할 수는 있지만, 이 패키지를 활용하면 상당히 편하다. flag 패키지는
	 * 커맨드라인 인수나 옵션의 순서를 따로 구분하지 않고, 커맨드라인 옵션을 처리하는 과정에서 에러가 발생하면
	 * 이에 대한 정보를 화면에 출력해주는 기능을 제공한다.
	 *
	 * flag 패키지의 가장 큰 장점은 표준 Go 라이브러리의 일부로 제공된다는 점이다. 따라서 충분한 테스트를
	 * 거치고 디버깅이 됐다고 볼 수 있다.
	 */
	"flag"
	"fmt"
)

func main() {
	// k라는 이름의 불리언 타입의 커맨드 라인 옵션을 정의한다. 기본(Default) 값은 true로 지정했다.
	// 이 문장의 마지막 매개변수는 이 프로그램을 사용하는 방법을 표현하기 위해 화면에 출력될 스트링이다.
	minusK := flag.Bool("k", true, "k")
	// 정수형 커맨드 라인 옵션을 정의한다.
	// flag.Int()로 정의한 플래그에 대해 입력된 값을 자동으로 정수형 값으로 변환해준다.
	// 따라서 직접 타입을 변환할 필요가 없다. 또한 flag 패키지는 주어진 정수값이 적절한 형태로 되어 있는지
	// 검사도 해준다.
	minusO := flag.Int("O", 1, "O")
	// 커맨드라인 옵션을 정의한 뒤에는 반드시 flag.Parse()를 호출해야 한다.
	flag.Parse()

	valueK := *minusK
	valueO := *minusO
	valueO++

	fmt.Println("-k:", valueK)
	fmt.Println("-O:", valueO)

	// ❯ go run simpleFlag.go -O 100
	// -k: true
	// -O: 101
	// ❯ go run simpleFlag.go -O=100
	// -k: true
	// -O: 101
	// ❯ go run simpleFlag.go -O=100 -k
	// -k: true
	// -O: 101
	// ❯ go run simpleFlag.go -O=100 -k false
	// -k: true
	// -O: 101
	// ❯ go run simpleFlag.go -O=100 -k=false
	// -k: false
	// -O: 101
	// ❯ go run simpleFlag.go -O=notAnInteger
	// invalid value "notAnInteger" for flag -O: parse error
	// Usage of /var/folders/nf/gb0bnlfn51b03fycsh6g9f1h0000gn/T/go-build541260395/b001/exe/simpleFlag:
	//   -O int
	//         O (default 1)
	//   -k    k (default true)
	// exit status 2
}
