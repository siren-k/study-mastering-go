package main

import (
	"fmt"
	"net"
)

/*
 * 네트워크 설정에서 핵심적인 요소는 네 가지로, 인터페이스의 IP 주소, 인터페이스의 넷마스크(Netmask),
 * 머신의 DNS 서버, 머신의 디폴트 게이트웨이 또는 디폴트 라우터가 있다.
 */
func main() {
	// 현재 머신에 존재하는 모든 인터페이스를 net.Interface 타입의 슬라이스에 담아 리턴한다.
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	// net.Interface 타입의 슬라이스에 담긴 각 원소를 방문하여 IP 주소를 출력한다.
	for _, i := range interfaces {
		fmt.Printf("Interface: %v\n", i.Name)
		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			fmt.Println(err)
		}
		addresses, err := byName.Addrs()
		for k, v := range addresses {
			fmt.Printf("Interface Address #%v: %v\n", k, v.String())
		}
		fmt.Println()
	}
}
