package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

/*
 * /dev/random 시스템 디바이스의 목적은 랜덤 데이터를 생성하는 것이다.
 * 프로그램을 테스트하거나 난수 생성기의 씨드를 제공하는데 활용될 수 있다.
 */
func main() {
	// /dev/random은 일반 파일과 같은 방식으로 연다.
	// 유닉스에서는 모든 것을 파일로 취급하기 때문이다.
	f, err := os.Open("/dev/random")
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	var seed int64
	// /dev/random 시스템 디바이스로부터 데이터를 읽기 위해 binary.Read() 함수를 호출했다.
	// 이 함수는 세 개의 매개변수를 받는다. 두 번째 매개변수(binary.LittleEndian)의 값은
	// 바이트 순서로 '리틀 엔디안'을 적용하도록 지정한다. 또 다른 옵션으로 binary.BigEndian도 있다.
	binary.Read(f, binary.LittleEndian, &seed)
	fmt.Println("Seed:", seed)
}
