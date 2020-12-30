package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	// net.Listen() 함수는 클라이언트의 연결 요청을 기다린다.
	// 두 번째 매개변수에 IP 주소는 없고 포트 번호만 있다면, net.Listen() 함수는 현재
	// 로컬 시스템에서 사용할 수 있는 IP 주소로부터 클라이언트 접속을 기다린다.
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	// net.Listener.Accept() 함수는 다음 연결이 들어올 때까지 기다리다가 연결되면 제네릭
	// 변수인 Conn을 리턴한다. 여기서 구현한 TCP 서버에서 한 가지 문제가 있다면 첫 번째로 들어온
	// TCP 클라이언트에 대해서만 서비스를 제공한다는 것이다. Accept() 함수를 뒤에 나올 for 루프
	// 밖에서 호출했기 때문이다.
	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}
		fmt.Print("-> ", string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC1123) + "\n"
		c.Write([]byte(myTime))
	}
}
