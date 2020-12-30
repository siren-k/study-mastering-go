package main

import (
	"fmt"
	"math"
	"net"
	"net/rpc"
	"os"
	"sharedRPC"
)

/*
 * MyInterace 인터페이스와 Power() 헬퍼 함수를 구현했다.
 */
type MyInterface struct{}

func Power(x, y float64) float64 {
	return math.Pow(x, y)
}

func (t *MyInterface) Multiply(arguments *sharedRPC.MyFloats, reply *float64) error {
	*reply = arguments.A1 * arguments.A2
	return nil
}

func (t *MyInterface) Power(arguments *sharedRPC.MyFloats, reply *float64) error {
	*reply = Power(arguments.A1, arguments.A2)
	return nil
}

/*
 * '원격 프로시저 호출(RPC, Remote Procedure Call)이란, TCP/IP를 이용해 클라이언트와 서버가
 * '프로세스간 통신(IPC, InterProcess Communication)'을 하기 위한 메커니즘이다.
 */
func main() {
	PORT := ":1234"
	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}

	myInterface := new(MyInterface)
	// rpc.Register() 함수를 사용함으로써 이 프로그램이 RPC 서버가 된다. 하지만 RPC 서버는
	// TCP를 사용하기 때문에 net.ResolveTCPAddr()과 net.ListenTCP() 함수도 여전히
	// 사용한다.
	rpc.Register(myInterface)
	t, err := net.ResolveTCPAddr("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := net.ListenTCP("tcp4", t)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		/*
		 * RemoteAddr() 함수는 RPC 클라이언트와 통신하는데 사용할 IP 주소와 포트 번호를
		 * 리턴한다. rpc.ServeConn() 함수는 RPC 클라이언트에게 서비스를 제공한다.
		 */
		fmt.Printf("%s\n", c.RemoteAddr())
		rpc.ServeConn(c)
	}
}
