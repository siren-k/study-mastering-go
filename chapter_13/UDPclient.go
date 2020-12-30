package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a host:port string")
		return
	}
	CONNECT := arguments[1]

	// net.ResolveUDPAddr() 함수는 두 번째 매개변수로 정의한 UDP 엔드포인트의
	// 주소를 리턴한다. 첫 번째 매개변수(udp4)는 이 프로그램에서 IPv4 프로토콜만
	// 지원하도록 지정한다.
	s, err := net.ResolveUDPAddr("udp4", CONNECT)
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
	defer c.Close()

	for {
		// 사용자가 입력한 텍스트는 표준 입력으로부터 읽는다.(bufio.NewReader(os.Stdin))
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		data := []byte(text + "\n")
		// 사용자가 입력한 텍스트를 받아서 UDP 서버로 보낸다.
		_, err = c.Write(data)
		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("Exiting UDP client!")
			return
		}
		if err != nil {
			fmt.Println(err)
			return
		}

		// 클라이언트 데이터가 전달된 후에는 UDP 서바가 데이터를 보낼 떄까지 기다려야 한다.
		buffer := make([]byte, 1024)
		n, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Reply: %s\n", string(buffer[0:n]))
	}

	// ❯ nc -v -u -l 127.0.0.1 8001
}
