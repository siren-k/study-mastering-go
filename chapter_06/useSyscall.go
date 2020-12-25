package main

import (
	"fmt"
	"os"
	"syscall"
)

/*
 * syscall 패키지는 저수준(Low-Level) OS 요소에 대한 함수와 타입을 방대하게 제공한다.
 * 또한, os, net, time과 같이 OS의 이식성에 관련된 인터페이스를 제공하는 다른 Go 패키지
 * 에서도 syscall 패키지를 널리 사용하고 있다. 따라서 Go 라이브러리 중에서 syscall 패키지만
 * 이식성이 뛰어난 것이 아니고, 이 패키지의 역할도 이식성에 있지 않다. 유닉스 시스템은 서로
 * 비슷한 점이 많지만 이와 동시에 차이점도 많다. 특히 시스템 내부를 들여다보면 그 차이가 더욱
 * 두드러진다. syscall 패키지의 주 목적은 이러한 차이점을 최대한 부드럽게 처리하는 것이다.
 * 이러한 특성과 함께 문서화도 잘 돼 있기 때문에 syscall은 굉장히 성공적인 패키지로 자리잡았다.
 *
 * 엄밀히 말하면 시스템 콜(System Call)은 애플리케이션이 OS의 커널에게 뭔가를 요청하기 위한
 * 인터페이스다. 따라서, 시스템 콜은 프로세스, 스토리지 디바이스, 데이터 출력, 네트워크 인터페이스,
 * 모든 종류의 파일 등과 같은 유닉스의 저수준 요소들을 접근하고 다루는 역할을 담당한다. 쉽게 말해
 * 시스템 콜 없이는 유닉스 시스템을 다룰 수 없다. strace(1)이나 dtrace(1)과 같은 유틸리티를
 * 이용하면, 유닉스 프로세스에 대한 시스템 콜을 살펴볼 수 있다.
 *
 * syscall 패키지를 사용하는 표준 Go 패키지를 알고 싶다면 유닉스 쉘에서 다음과 같이 커맨드를
 * 실행한다.
 * ❯ grep \"syscall\" `find /usr/local/Cellar/go/current/libexec/src -name "*.go"`
 * /usr/local/Cellar/go/current/libexec/src/cmd/go/testdata/testterminal18153/terminal_test.go:	"syscall"
 * /usr/local/Cellar/go/current/libexec/src/cmd/go/go_unix_test.go:	"syscall"
 * /usr/local/Cellar/go/current/libexec/src/cmd/go/internal/modload/stat_unix.go:	"syscall"
 *                                        .
 *                                        .
 *                                        .
 * /usr/local/Cellar/go/current/libexec/src/vendor/golang.org/x/net/nettest/nettest_unix.go:import "syscall"
 * /usr/local/Cellar/go/current/libexec/src/vendor/golang.org/x/sys/cpu/syscall_aix_ppc64_gc.go:	"syscall"
 * /usr/local/Cellar/go/current/libexec/src/debug/pe/file_test.go:import "syscall"
 */
func main() {
	/*
	 * syscall.Syscall() 호출을 통해 프로세스 ID와 사용자 ID를 알아내서 화면에 출력한다.
	 * syscall.Syscall()의 첫 번째 매개변수는 원하는 정보의 종류를 지정한다.
	 * syscall.Syscall()은 이식성을 제공하지 않는다.
	 */
	pid, _, _ := syscall.Syscall(39, 0, 0, 0)
	fmt.Println("My pid is", pid)
	uid, _, _ := syscall.Syscall(24, 0, 0, 0)
	fmt.Println("User ID is", uid)

	message := []byte{'H', 'e', 'l', 'l', 'o', '!', '\n'}
	fd := 1
	// syscall.Write()를 통해 화면에 메시지를 출력한다. 첫 번째 매개변수는 쓸 파일에 대한
	// 디스크립터를 지정하며, 두 번째 매개변수는 출력할 실제 메시지를 담은 바이트 슬라이스를
	// 지정한다. sysall.Write() 함수는 이식성을 제공한다.
	syscall.Write(fd, message)

	fmt.Println("Using syscall.Exec()")
	command := "/bin/ls"
	env := os.Environ()
	// syscall.Exec() 함수로 외부 커맨드를 실행하는 방법을 알 수 있다. 그런데 이 커맨드의
	// 실행 결과는 곧바로 화면에 출력되며 이 과정을 제어할 수는 없다.
	syscall.Exec(command, []string{"ls", "-a", "-x"}, env)
}
