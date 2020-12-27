package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/*
 * 디스크에 파일을 쓸 때 주로 io.Writer 인터페이스에서 제공하는 기능을 이용한다.
 */
func main() {
	// 각 줄을 쓸 때마나 사용할 s 바이트 슬라이스 정의
	s := []byte("Data to write\n")

	f1, err := os.Create("f1.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f1.Close()
	// f1.txt 파일 쓰기
	fmt.Fprint(f1, string(s))

	f2, err := os.Create("f2.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f2.Close()
	// file.WriteString() 함수를 이용하여 파일 쓰기
	n, err := f2.WriteString(string(s))
	fmt.Printf("wrote %d bytes", n)

	f3, err := os.Create("f3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// bufio.NewWriter() 함수를 이용하여 쓸 파일을 연다.
	w := bufio.NewWriter(f3)
	// bufio.WriteString() 함수를 이용하여 s 바이트 슬라이스를 파일에 쓴다.
	n, err = w.WriteString(string(s))
	fmt.Printf("wrote %d bytes\n", n)
	w.Flush()

	f4 := "f4.txt"
	// ioutil.WriteFile() 함수는 파일의 생성 및 파일 쓰기를 동시에 수행한다.
	err = ioutil.WriteFile(f4, s, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	f5, err := os.Create("f5.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// io.WriteString을 이용하여 데이트를 파일에 쓸 수 있다.
	n, err = io.WriteString(f5, string(s))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)
}
