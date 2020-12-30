package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide prot number")
		return
	}

	SERVER := "localhost" + ":" + arguments[1]
	s, err := net.ResolveTCPAddr("tcp", SERVER)
	if err != nil {
		fmt.Println(err)
		return
	}

	// net.ListenTCP() 함수는 net.Listen()의 TCP 네트워크 버전이다.
	l, err := net.ListenTCP("tcp", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	buffer := make([]byte, 1024)
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting TCP server!")
			conn.Close()
			return
		}

		fmt.Print("> ", string(buffer[0:n]))
		_, err = conn.Write(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
