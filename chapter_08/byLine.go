package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func lineByLine(file string) error {
	var err error

	// 읽을 파일을 열기
	f, err := os.Open(file)
	if err != nil {
		return nil
	}
	defer f.Close()

	// 읽을 파일을 열었다면 bufio.NewReader()로 리더를 새로 생성한다.
	r := bufio.NewReader(f)
	for {
		// 그런 다음, 이 리더를 이용해 bufio.ReadString()를 호출해서 입력 파일을
		// 한 줄씩 읽는다. 여기서 bufio.ReadString()의 매개변수로 줄바꿈 문자을
		// 지정했는데, 그럼나 이 문자가 나타날 때까지 계속 읽는다. 줄바꿈 문자가 나타날
		// 때마다 bufio.ReadString()을 계속 호출함으로써 파일을 한 줄씩 읽을 수 있다.
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
		// 여기서 읽은 줄을 화면에 출력할 때 fmt.Println()이 아닌 fmt.Print()를 사용했다.
		// 읽은 줄 안에 줄바꿈 문자가 이미 포함되어 있기 때문이다.
		fmt.Print(line)
	}
	return nil
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: byLine <file1> [<file2>. ..]\n")
		return
	}

	for _, file := range flag.Args() {
		err := lineByLine(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
