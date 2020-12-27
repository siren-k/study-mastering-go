package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// 프로그램의 모든 곳에서 minusD와 minusF에 접근할 수 있도록 전역 변수로 선언
var minusD bool = false
var minusF bool = false

/*
 * -d 옵션은 디렉토리 경로 앞에 별표 문자(*)를 출력한다.
 * -f 옵션은 디렉토리가 아닌 일반 파일에 대한 경로 앞에 덧셈 문자(+)를 추가한다.
 */
func walk(path string, info os.FileInfo, err error) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	mode := fileInfo.Mode()
	// mode.isRegular() 함수는 일반 파일인지 확인
	if mode.IsRegular() && minusF {
		fmt.Println("+", path)
		return nil
	}

	// mode.isDir() 함수는 디렉토리인지 확인
	if mode.IsDir() && minusD {
		fmt.Println("*", path)
		return nil
	}

	fmt.Println(path)
	return nil
}

func main() {
	/*
	 * 커맨드라인 인수를 효율적으로 처리하도록 flag 패키지를 이용하고 있다.
	 */
	starD := flag.Bool("d", false, "Signify directories")
	plusF := flag.Bool("f", false, "Signify regular files")
	flag.Parse()
	flags := flag.Args()

	Path := "."
	if len(flags) == 1 {
		Path = flags[0]
	}

	minusD = *starD
	minusF = *plusF

	err := filepath.Walk(Path, walk)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
