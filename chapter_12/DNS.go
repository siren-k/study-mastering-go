package main

import (
	"fmt"
	"net"
	"os"
)

// lookIP() 함수는 IP 주소를 입력 받아서 net.LookupAddr() 함수를 이용하여
// 주소와 매칭되는 도메인 네임 목록을 리턴한다.
func lookIP(address string) ([]string, error) {
	hosts, err := net.LookupAddr(address)
	if err != nil {
		return nil, err
	}
	return hosts, nil
}

// lookHostname() 함수는 호스트 네임을 입력 받아서 net.LookupHost() 함수를
// 이용해 그 네임에 해당되는 IP 주소 목록을 리턴한다.
func lookHostname(hostname string) ([]string, error) {
	IPs, err := net.LookupHost(hostname)
	if err != nil {
		return nil, err
	}
	return IPs, nil
}

/*
 * DNS란 Domain Name System의 줄임말로 IP 주소를 acornpub.co.kr과 같은 주소로
 * 변환하거나 반대로 호스트네임을 IP 주소로 변환한다. 커맨드라인 인수로 주어진 값이 올바른
 * 형식의 IP 주소라면 이를 그대로 사용한다. 그렇지 않다면 그 값이 호트네임이라 간주하고
 * IP 주소로 변환한다.
 */
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		return
	}

	input := arguments[1]
	// net.ParseIP() 함수는 인수로 전달된 스트링을 IPv4나 IPv6 주소로 파싱한다.
	// IP 주소가 올바른 값이 아니라면 net.ParseIP()는 nil을 리턴한다.
	IPaddress := net.ParseIP(input)
	if IPaddress == nil {
		IPs, err := lookHostname(input)
		if err == nil {
			for _, singleIP := range IPs {
				fmt.Println(singleIP)
			}
		}
	} else {
		hosts, err := lookIP(input)
		if err == nil {
			for _, hostname := range hosts {
				fmt.Println(hostname)
			}
		}
	}
}
