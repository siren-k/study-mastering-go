package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * 한 줄의 텍스트에서 특정한 열(Column)을 선택하여 화면에 출력하는 유틸리티 프로그램
 */
func main() {
	/*
	 * 최소한 2개의 커맨드라인 인수(Command-Line Arguments)를 받는다.
	 * 첫 번째 인수는 원하는 열 번호고, 두 번째 인수는 처리할 텍스트 파일에 대한 경로이다.
	 */
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Printf("usage: selectColumn colun <file1> [<file2>...<fileN>]\n")
		os.Exit(1)
	}

	/*
	 * 입력된 열 번호가 숫자값인지 확인
	 */
	temp, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("Column value is not an integer:", temp)
		return
	}

	/*
	 * 입력된 열 번호가 0보다 큰지 확인
	 */
	column := temp
	if column < 0 {
		fmt.Println("Invalid Column number!")
		os.Exit(1)
	}

	for _, filename := range arguments[2:] {
		/*
		 * 입력된 테스트 파일이 실제로 있는지 검사
		 */
		fmt.Println("\t\t", filename)
		f, err := os.Open(filename) // 파일을 여는 작업
		if err != nil {
			fmt.Printf("error opening file %s\n", err)
			continue
		}
		defer f.Close()

		/*
		 * 각 줄을 공백 문자를 기준으로 나누는 가장 간단한 형태의 패턴 매칭
		 */
		r := bufio.NewReader(f)
		for {
			/*
			 * 매개변수로 지정한 내용을 처음 발견할 때까지 파일을 읽는다.
			 * 유닉스의 줄바꿈 문자인 \n을 발견할 때까지 파일을 한 줄 한줄 읽는다.
			 */
			line, err := r.ReadString('\n') // 바이트 슬라이스를 리턴
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file %s", err)
			}

			/*
			 * 줄을 열 단위로 분할
			 */
			data := strings.Fields(line)
			if len(data) >= column {
				fmt.Println("column", column, (data[column-1]))
			}
		}
	}
}
