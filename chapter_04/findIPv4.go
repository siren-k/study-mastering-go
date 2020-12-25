package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"regexp"
)

func findIP(input string) string {
	// 각각의 부분을 0 ~ 255 범위를 가지는 확인하기 위한 정규표현식을 정의
	partIp := "(25[0-5]|2[0-4]0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	// 찾으려는 숫자가 4 부분으로 구성된다고 표현
	grammer := partIp + "\\." + partIp + "\\." + partIp + "\\." + partIp
	matchMe := regexp.MustCompile(grammer)
	return matchMe.FindString(input)
}

func main() {
	/*
	 * 입력된 커맨드라인 인수의 수가 충분한지 확인
	 */
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("usage: %s logFile\n", filepath.Base(arguments[0]))
		os.Exit(1)
	}

	/*
	 * 입력된 파일에서 IPv4 주소를 찾아 화면에 출력
	 */
	for _, filename := range arguments[1:] {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("error opening file %s\n", err)
			os.Exit(-1)
		}
		defer f.Close()

		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println("error reading file %s", err)
				break
			}

			ip := findIP(line)
			trial := net.ParseIP(ip)
			if trial.To4() != nil {
				fmt.Println(ip)
			}
		}
	}
}
