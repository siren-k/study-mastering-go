package main

import (
	"fmt"
	"net/rpc"
	"os"
	"sharedRPC"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a host:port string!")
		return
	}

	CONNECT := arguments[1]
	// RPC 서버는 TCP를 사용하지만 서버에 연결할 때 net.Dial() 함수 대신
	// rpc.Dial() 함수를 사용한다.
	c, err := rpc.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	args := sharedRPC.MyFloats{16, -0.5}
	var reply float64

	// RPC 클라이언트와 RPC 서버는 Call() 함수를 이용하여 함수 이름과 그 함수의 인수, 그리고
	// 함수 호출의 결과를 주고 받는다. RPC 클라이언트는 그 함수의 실제 구현 내용에 대해서는
	// 전혀 모른다.
	err = c.Call("MyInterface.Multiply", args, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply (Multiply): %f\n", reply)

	err = c.Call("MyInterface.Power", args, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply (Power): %f\n", reply)
}
