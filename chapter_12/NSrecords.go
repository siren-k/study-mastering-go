package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need a domain name!")
		return
	}

	domain := arguments[1]
	// 모든 작업은 net.LookupNS() 함수로 처리한다. 이 함수는 주어진 도메인에
	// 대한 NS 레코드를 net.NS 타입의 슬라이스 변수로 리턴한다. 그래서 슬라이스에
	// 담긴 모든 net.NS 원소에 대해 루프를 돌며 Host 필드를 화면에 출력한다.
	NSs, err := net.LookupNS(domain)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, NS := range NSs {
		fmt.Println(NS.Host)
	}
}
