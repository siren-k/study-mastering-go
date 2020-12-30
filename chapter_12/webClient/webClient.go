package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

/*
 * webClient.go가 가진 가장 큰 문제는 전반적인 과정을 제어할 수 없다는 것이다.
 * HTML 문서 전체를 받어나 아루런 결과도 받지 못할 뿐이다.
 */
func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		return
	}

	// 원하는 URL을 커맨드라인 인수로 받는다.
	URL := os.Args[1]
	// 모든 작업은 http.Get()으로 처리한다. 매개변수나 옵션에 신경쓰지 않다록 되므로 상당히 편리하다.
	// 하지만 이렇게 하면 전반적인 과정에 대해 유연하게 대처하기 힘들다.
	// http.Get() 함수는 http.Response를 리턴한다.
	data, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		defer data.Body.Close()
		// http.Response 구조체의 Body 필드에 담긴 내용을 표준 출력으로 복사했다.
		_, err := io.Copy(os.Stdout, data.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
