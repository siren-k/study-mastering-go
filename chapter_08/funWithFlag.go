package main

import (
	"flag"
	"fmt"
	"strings"
)

type NamesFlag struct {
	Names []string
}

func (s *NamesFlag) GetNames() []string {
	return s.Names
}

func (s *NamesFlag) String() string {
	return fmt.Sprint(s.Names)
}

/*
 * Set() 함수는 필요한 커맨드라인 옵션이 모두 설정됐는지 확인한다. 그런 다음, 입력을 받아
 * strings.Split() 함수를 이용하여 인수들을 구분한다. 이렇게 추려낸 인수를 NamesFlag
 * 구조체의 Names 필드에 저장한다.
 */
func (s *NamesFlag) Set(v string) error {
	if len(s.Names) > 0 {
		return fmt.Errorf("Cannot use names flag more than once!")
	}

	names := strings.Split(v, ",")
	for _, item := range names {
		s.Names = append(s.Names, item)
	}
	return nil
}

func main() {
	var manyNames NamesFlag
	minusK := flag.Int("k", 0, "An int")
	minusO := flag.String("o", "Mihalis", "The name")
	flag.Var(&manyNames, "names", "Comma-separated list")

	flag.Parse()
	fmt.Println("-k:", *minusK)
	fmt.Println("-o:", *minusO)

	for i, item := range manyNames.GetNames() {
		fmt.Println(i, item)
	}

	fmt.Println("Remaining command-line arguments:")
	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)
	}

	// ❯ go run funWithFlag.go -names=Mihalis,Jim,Athina 1 two three
	// -k: 0
	// -o: Mihalis
	// 0 Mihalis,Jim,Athina
	// Remaining command-line arguments:
	// 0 : 1
	// 1 : two
	// 2 : three
	// ❯ go run funWithFlag.go -names=Mihalis,Jim,Athina 1 two three
	// -k: 0
	// -o: Mihalis
	// 0 Mihalis
	// 1 Jim
	// 2 Athina
	// Remaining command-line arguments:
	// 0 : 1
	// 1 : two
	// 2 : three
	// ❯ go run funWithFlag.go -Invalid=Marietta 1 two three
	// flag provided but not defined: -Invalid
	// Usage of /var/folders/nf/gb0bnlfn51b03fycsh6g9f1h0000gn/T/go-build228624204/b001/exe/funWithFlag:
	//   -k int
	//         An int
	//   -names value
	//         Comma-separated list
	//   -o string
	//         The name (default "Mihalis")
	// exit status 2
	// ❯ go run funWithFlag.go -names=Marietta -names=Mihalis 1 two three
	// invalid value "Mihalis" for flag -names: Cannot use names flag more than once!
	// Usage of /var/folders/nf/gb0bnlfn51b03fycsh6g9f1h0000gn/T/go-build160222415/b001/exe/funWithFlag:
	//   -k int
	//         An int
	//   -names value
	//         Comma-separated list
	//   -o string
	//         The name (default "Mihalis")
	// exit status 2
}
