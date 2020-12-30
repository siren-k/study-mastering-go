package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
 * 피보나치 수열에 속한 자연수를 생성하는 f() 함수를 구현하고 있다. 여기에 적용된 알고리즘을
 * 처음 볼 때는 이해하기 쉽지 않지만, 굉장히 효율적이고 속도가 빠르다. 먼저 f() 함수는 fn란
 * 맵을 사용하는데 피보나치 수열을 계산하는 용도로는 다소 생소한 편이다. 두 번째로 f() 함수는
 * for 루프를 사용하는데 이 점 역시 좀 어색하다. 마지막으로 f() 함수는 재귀호출(Recursion)을
 * 사용하지 않는데 이 부분이 바로 알고리즘의 연산 속도에 가장 큰 영향을 미치는 요인이다.
 *
 * f()에 나온 알고리즘은 '동적 프로그래밍(Dynamic Programming)'이란 기법을 적용하고 있다.
 * 그래서 피보나치 수를 하나씩 계산할 때마다 같은 값을 다시 계산하지 않도록 fn 맵에 집어넣는다.
 * 간단한 개념이지만 이를 통해 시간을 크게 절약할 수 있다. 계산해야 할 피보나치 수가 클수록 그
 * 효과는 두드러진다. 동일한 피보나치 수를 여러 번 계산할 필요가 없기 떄문이다.
 */
func f(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	return fn[n]
}

/*
 * handleConnection() 함수는 TCP 서버에 접속하는 여러 클라이언트를 동시에 처리한다.
 */
func handleConnection(c net.Conn) {
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}

		fibo := "-1\n"
		n, err := strconv.Atoi(temp)
		if err == nil {
			fibo = strconv.Itoa(f(n)) + "\n"
		}
		c.Write([]byte(string(fibo)))
	}
	time.Sleep(5 * time.Second)
	c.Close()
}

/*
 * Go 루틴을 이용하여 동시성을 지원하도록 TCP 서버를 구현
 * TCP 서버에 연결이 들어올 때마다 각각의 요청을 처리하는 Go 루틴을 새로 구동한다.
 * 이렇게 하다보면 많은 요청을 받을 수 있어서 TCP 서버에서 여러 개의 클라이언트를
 * 동시에 처리할 수 있다.
 */
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		// 동시성을 지원하는 부분으로 TCP 클라이언트가 접속할 때마다
		// 여기서 Go 루틴을 새로 생성한다. 이렇게 생성된 여러 개의
		// Go 루틴은 동시에 실행된다. 따라서 서버는 여러 클라이언트를
		// 동시에 처리할 수 있다.
		go handleConnection(c)
	}
}
