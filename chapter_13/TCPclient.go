package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/*
 * TCP는 Transmission Control Protocol(전송 제어 프로토콜)의 약자로, 안정적인 전송을
 * 제공한다는 점이 특징인 프로토콜이다. TCP 패킷의 헤더마다 출발지 포트(Source Port)와
 * 목적지 포트(Destination Port) 필드가 있다. 이러한 두 필드와 더불어 출발지 IP 주소와
 * 목적지 IP 주소를 한데 묶어 하나의 TCP 연결에 대한 고유한 식별자로 사용된다.
 */
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	// net.Dial() 함수는 원격 서버에 연경하는데 사용된다. net.Dial() 함수의 첫 번째 매개변수는
	// 사용할 네트워크를 정의하는 반면, 두 번째 매개변수는 서버의 주소를 지정한다. 이 값에는 포트 번호도
	// 반드시 포함해야 한다. 첫 번째 매개변수로 지정할 수 있는 값으로는 tcp, tcp4(IPv4 전용),
	// tcp6(IPv6 전용), udp, udp4(IPv4 전용), udp6(IPv6 전용), ip, ipv4(IPv4 전용),
	// ipv6(IPv6 전용) Unix(유닉스 소켓), Unixgram, Unixpacket 등이 있다.
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// 입력값은 os.Stdin 파일로부터 읽는다.
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		// reader.ReadString()에서 반환한 error 값은 무시했는데, 바람직한 작성 방식은
		// 아니지만 여기서는 공간을 정략하기 위해 생략했다. 당연히 프로덕션용 소프트웨어를 작성할
		// 때는 이렇게 하면 안 된다.
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: ", message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}

	// nc -l 127.0.0.1 8001
	// ❯ go run TCPclient.go localhost:8001
}
