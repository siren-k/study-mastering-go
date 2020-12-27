package main

import (
	"fmt"
	"os"
)

/*
 * 유닉스 시스템 프로그램밍에서 인기 있는 주제 중 하나는 유닉스 파일 접근 권한에 대한 것이다.
 * 이 절에서는 주어진 파일에 대한 접근 권한을 화면에 출력하는 방법을 소개한다. 단, 이 때 해당
 * 파일에 대한 접근 권한을 갖고 있어야 한다.
 */
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Printf("usage: permissions filename\n")
		return
	}

	filename := arguments[1]
	// os.Stat()을 호출하면 데이터가 많이 들어 있는 거대한 구조체 하나를 리턴한다. 여기서는
	// 지정한 파일에 대한 접근 권한만 알면 되기 때문에 Mode() 함수를 호출해서 그 결과를 화면에
	// 출력했다. 실제로 mode.String()[1:10]과 같이 작성하여 결과의 일부분만 화면에
	// 출력했다. 이 부분에 우리가 원하는 정보가 있기 떄문이다.
	info, _ := os.Stat(filename)
	mode := info.Mode()
	fmt.Println(filename, "mode is", mode.String()[1:10])

	// ❯ go run permissions.go f1.txt
	// f1.txt mode is rw-r--r--
}
