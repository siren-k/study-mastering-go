package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var timeout = time.Duration(time.Second)

// Timeout() 함수는 http.Transport 변수의 Dial 필드에서 사용한다.
func Timeout(network, host string) (net.Conn, error) {
	conn, err := net.DialTimeout(network, host, timeout)
	if err != nil {
		return nil, err
	}
	// SetDeadline() 함수는 주어진 네트워크 연결에서 읽기와 쓰기를 수행하는데 적용한 데드라인을 설정한다.
	// SetDeadline() 함수의 작동 방식의 특성상 읽기나 쓰기 연산을 수행하기 전에 SetDeadline()을
	// 호출한다. 참고로 Go 언어에서는 타임아웃을 구현할 때 데드라인을 이용한다. 따라서 애플리케이션이 데이터를
	// 보내거나 받을 때마다 이 값을 리셋할 필요는 없다.
	conn.SetDeadline(time.Now().Add(timeout))
	return conn, nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s URL TIMEOUT\n", filepath.Base(os.Args[0]))
		return
	}

	if len(os.Args) == 3 {
		temp, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Using Default Timeout!")
		} else {
			timeout = time.Duration(time.Duration(temp) * time.Second)
		}
	}

	URL := os.Args[1]
	t := http.Transport{Dial: Timeout}
	client := http.Client{Transport: &t}

	data, err := client.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		defer data.Body.Close()
		_, err := io.Copy(os.Stdout, data.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
