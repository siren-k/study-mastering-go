package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func charByChar(file string) error {
	var err error
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			return err
		}

		// 파일을 한 줄씩 받아서 range를 이용하여 쪼개고 있다. range는 두 값을 리턴한다.
		// line 변수에서 현재 문자가 있는 위치를 나타내는 첫 번째 값을 버리고 두 번째만 가져온다.
		// 그런데 이 값은 문자가 아니다. 그래서 string() 함수를 이용해 문자로 변환했다.
		for _, x := range line {
			fmt.Println(string(x))
		}
	}
	return nil
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: byChar <file1> [<file2> ...]\n")
		return
	}

	for _, file := range flag.Args() {
		err := charByChar(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
