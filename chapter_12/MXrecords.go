package main

import (
	"fmt"
	"net"
	"os"
)

/*
 * DNS 요청으로 흔히 들어오는 또 다른 정보로 도메인에 대한 MX 레코드가 있다. MX 레코드는
 * 도메인의 메일 서버를 지정한다.
 */
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need a domain name!")
		return
	}

	domain := arguments[1]
	MXs, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, MX := range MXs {
		fmt.Println(MX.Host)
	}
}
