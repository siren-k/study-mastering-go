package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func printFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		io.WriteString(os.Stdout, scanner.Text())
		io.WriteString(os.Stdout, "\n")
	}
	return nil
}

/*
 * 유닉스 철학에 따르면 유닉스 커맨드라인 유틸리티는 반드시 한 가지 작업만 수행하고, 그것도 잘 해야 한다.
 * 그래서 여러 작업을 수행하는 거대한 유틸리티를 만들지 않고, 여러 개의 조그만 프로그램의 형태로 구현해야
 * 한다. 그리고 각각을 서로 결합함으로써 원하는 작업을 수행할 수 있도록 구성해야 한다. 이처럼 두 개 이상의
 * 유닉스 커맨드라인 유틸리티가 서로 통신하며 협업하는데 가장 많이 사용하는 방법은 파이프를 이용하는 것이다.
 * 유닉스 파이프(Unix Pipe)는 한 커맨드라인 유틸리티에서 출력한 결과를 다른 커맨드라인 유틸리티의 입력으로
 * 전달한다. 이 때, 두 개 이상의 프로그램을 연결할 수도 있다. 유닉스 파이프를 나타내는 기호는 '|'다.
 *
 * 파이프는 다음과 같은 두 가지 심각한 한계가 있다. 하나는 대부분 한 방향으로만 통신한다는 점이고, 다른
 * 하나는 조상이 같은 프로세스에 대해서만 적용할 수 있다는 점이다. 유닉스 파이프를 이렇게 구현한 배경에는
 * 처리할 파일이 없다면 표준 입력으로부터 받을 때까지 기다리도록 하기 위해서다. 마찬가지로, 결과를 파일에
 * 저장하라는 지시가 없다면 그 결과를 표준 출력으로 보내야 하기 때문이다. 그래야 사용자가 보거나 다른
 * 프로그램이 이를 처리할 수 있다.
 */
func main() {
	filename := ""
	arguments := os.Args
	// cat.go에 아무런 커맨드라인 인수를 주지 않고 실행하면 표준 입력을 표준 출력으로
	// 복사하기만 한다. 이 동작은 io.Copy(os.Stdout, os.Stdin)이란 문장으로
	// 표현했다. 하지만 커맨드라인 인수를 지정할 경우, 모든 인수를 들어온 순서대로 처리한다.
	if len(arguments) == 1 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	for i := 1; i < len(arguments); i++ {
		filename = arguments[i]
		err := printFile(filename)
		if err != nil {
			fmt.Println(err)
		}
	}
}
