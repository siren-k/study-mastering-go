package main

import (
	"fmt"
	"html/template"
	"os"
)

type TextTemplateEntry struct {
	Number int
	Square int
}

/*
 * 템플릿은 주로 외부 파일에 저장하기 때문에 여기서 소개하는 예제도 text.gotext란 템플릿
 * 파일을 사용하도록 구성했다.
 */
func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Need the template file!")
		return
	}

	tFile := arguments[1]
	// DATA 변수는 2차원 슬라이스로서 데이터에 대한 초기 버전을 담고 있다.
	DATA := [][]int{{-1, 1}, {-2, 4}, {-3, 9}, {-4, 16}}

	var Entires []TextTemplateEntry
	for _, i := range DATA {
		if len(i) == 2 {
			temp := TextTemplateEntry{Number: i[0], Square: i[1]}
			Entires = append(Entires, temp)
		}
	}

	/*
	 * template.Must() 함수는 반드시 필요한 초기화 작업을 수행한다. 이 함수는 Template 타입의
	 * 값을 리턴하는데, 이는 파싱한 템플릿을 구조체다. template.ParseGlob() 함수는 외부 템플릿
	 * 파일을 읽는다. 또한 데이터를 처리하고 그 결과를 os.Stdout을 이용하여 파일에 출력하는 것과
	 * 같은 모든 작업은 template.Execute() 함수에서 처리한다.
	 *
	 * 텍스트 템플릿 파일에서 공백 라인은 나름 의미가 있으며, 최종 결과에 빈 줄로 출력된다.
	 */
	t := template.Must(template.ParseGlob(tFile))
	t.Execute(os.Stdout, Entires)

	// ❯ go run textT.go text.gotext
	// Calculating the squares of some integers
	// The square of -1 is 1
	// The square of -2 is 4
	// The square of -3 is 9
	// The square of -4 is 16
}
