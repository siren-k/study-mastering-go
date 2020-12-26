package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

/*
 * io.Read()의 첫 번째 리턴 값을 이용해 그 크기만큼 가진 바이트 슬라이스를 리턴한다.
 */
func readSize(f *os.File, size int) []byte {
	buffer := make([]byte, size)

	// io.Reader.Read() 메소드는 두 개의 매개변수를 리턴한다.
	// 하나는 읽은 바이트 수고, 다른 하나는 error 변수다.
	// 사소한 차이라고 볼 수도 있지만, 파일의 끝 부분에 도달할 때 출려될 내용이 입력된 파알과
	// 같은지, 그리고 쓸데 없는 문자가 들어 있지는 않은지 확인하는데 유용하다.
	n, err := f.Read(buffer)
	// io.EOF는 파일의 끝 부분에 도달했다는 것을 알려주는 에러이다.
	// 이러한 에러가 발생하면 함수는 리턴한다.
	if err == io.EOF {
		return nil
	}

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return buffer[0:n]
}

func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Println("<buffer size> <filename>")
		return
	}

	bufferSize, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	file := arguments[2]
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	for {
		readData := readSize(f, bufferSize)
		if readData != nil {
			fmt.Print(string(readData))
		} else {
			break
		}
	}
}
