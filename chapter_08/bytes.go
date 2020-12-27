package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

/*
 * 표준 Go 패키지인 bytes는 '바이트 슬라이스(byte slice)를 다루는데 필요한 함수들을 제공한다.
 * strings 패키지에서 '스트링(string)'과 관련된 기능을 제공하는 방식과 같다.
 */
func main() {
	// bytes.Buffer 변수를 생성
	var buffer bytes.Buffer
	// buffer.Writer()와 fmt.Fprintf() 함수를 이용해 buffer 변수에 데이터를 저장한다.
	buffer.Write([]byte("This is"))
	fmt.Fprintf(&buffer, " a string!\n")
	// 첫 번째 buffer.WriteTo()는 buffer에 담긴 내용을 출력한다.
	// 하지만, 두 번째 buffer.WriteTo()는 첫 번째로 buffer.WriteTo()르 호출하고 나서
	// buffer 변수가 비어 있으므로 아무것도 출력하지 않는다.
	buffer.WriteTo(os.Stdout)
	buffer.WriteTo(os.Stdout)

	// buffer 변수를 초기화
	buffer.Reset()
	// Write() 함수는 buffer 변수에 데이터를 저장한다.
	buffer.Write([]byte("Mastering Go!"))
	// bytes.NewReader()로 새 리더(Reader)를 하나 생성한다.
	r := bytes.NewReader([]byte(buffer.String()))
	fmt.Println(buffer.String())
	for {
		b := make([]byte, 3)
		// io.Reader 인터페이스에서 제공하는 Read() 함수를 이용하여
		// buffer 변수에 있는 데이터를 읽는다.
		n, err := r.Read(b)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Read %s Bytes: %d", b, n)
	}
}
