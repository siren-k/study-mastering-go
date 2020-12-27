package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	/*
	 * strings.NewReader() 함수는 인수로 전달한 스트링에 대한 읽기 전용 리더(Reader)를
	 * 생성한다. strings.Reader 오브젝트는 io.Reader, io.ReaderAt, io.Seeker,
	 * io.WriterTo, io.ByteScanner, io.RuneScanner 인터페이스를 구현한다.
	 *
	 * strings.Reader의 io.Reader 인터페이스에서 제공하는 Read() 함수로 바이트 단위로
	 * 스트링을 읽는다.
	 */
	r := strings.NewReader("test")
	fmt.Println("r length:", r.Len())

	/*
	 * 생성된 용량이 1인 바이트 슬라이스를 선언하여 strings.NewReader() 함수에서 반환된
	 * r에서 1 바이트씩 읽어 화면에 출력한다.
	 */
	b := make([]byte, 1)
	for {
		n, err := r.Read(b)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Read %s Bytes: %d\n", b, n)
	}

	/*
	 * io.Reader 인터페이스를 이용하여 데이터를 읽어서 io.Writer 인터페이스를 통해 데이터를 쓰는 기능 구현
	 * ==> strings 패키지에서 제공하는 기능을 활용해 표준 에러에 쓰는 방법을 보여주고 있다.
	 *
	 * strings.NewReader() 함수는 인수로 전달한 스트링에 대한 읽기 전용 리더(Reader)를
	 * 생성한다. strings.Reader 오브젝트는 io.Reader, io.ReaderAt, io.Seeker,
	 * io.WriterTo, io.ByteScanner, io.RuneScanner 인터페이스를 구현한다.
	 */
	s := strings.NewReader("This is an error!\n")
	fmt.Println("r length:", s.Len())
	n, err := s.WriteTo(os.Stderr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Wrote %d bytes to os.Stderr\n", n)

	// ❯ go run str.go
	// r length: 4
	// Read t Bytes: 1
	// Read e Bytes: 1
	// Read s Bytes: 1
	// Read t Bytes: 1
	// r length: 18
	// This is an error!
	// Wrote 18 bytes to os.Stderr
	// ❯ go run str.go 2>/dev/null
	// r length: 4
	// Read t Bytes: 1
	// Read e Bytes: 1
	// Read s Bytes: 1
	// Read t Bytes: 1
	// r length: 18
	// Wrote 18 bytes to os.Stderr
}
