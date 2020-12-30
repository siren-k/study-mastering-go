package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptrace"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: URL\n")
		return
	}

	URL := os.Args[1]
	// http.Client 오브젝트는 서버로 요청을 보내고 응답을 받는 수단을 제공한다. HTTP에 대한
	// 세부 사항을 디폴트값이 아닌 다른 값으로 설정하고 싶으면, 이 오브젝트에 있는 Transport
	// 필드를 통해 설정한다.여기서 주의할 점은 프로덕션 환경에서는 http.Client 오브젝트의
	// 디폴트값을 절대로 사용해선 안 된다. 요청에 대한 타임아웃값이 설정되어 있지 않기 때문에
	// 프로그램과 Go 루틴의 성능에 악영향을 미칠 수 있다.또한, http.Client 오브젝트는 동시성
	// 프로그램에서 안전하게 사용하도록 설계된 것이다.
	client := http.Client{}

	req, _ := http.NewRequest("GET", URL, nil)
	// httptrace.ClientTrace 오브젝트는 관심있는 이벤트를 정의한다. 이러한 이벤트 중에
	// 하나가 발생하면 이를 처리하는 코드가 실행된다. 여기서 지원하는 이벤트의 종류와 목적에 대한
	// 자세한 사항은 net/http/httptrace 패키지에 대한 공식 문서를 참고한다.
	trace := &httptrace.ClientTrace{
		GotFirstResponseByte: func() {
			fmt.Println("First response byte!")
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
		ConnectStart: func(network, addr string) {
			fmt.Println("Dial start")
		},
		ConnectDone: func(network, addr string, err error) {
			fmt.Println("Dial done")
		},
		WroteHeaders: func() {
			fmt.Println("Wrote headers")
		},
	}

	// httptrace.WithClientTrace() 함수는 현재 지정된 부모 컨텍스트를 기반으로 새로운 컨텍스트를
	// 반환한다. 여기서 httptrace.DefaultClTransport.RoundTrip() 함수는 현재 요청을 추적하도록
	// http.TransPort.RoundTrip을 감싸고 있다. 여기서 주목할 부분은, Go 언어에서 제공하는 HTTP
	// 트레이싱 기능은 하나의 http.Transport.RoundTrip에 대한 이벤트만 추적할 수 있도록 설계됐다는
	// 점이다. 하지만 하나의 HTTP 요청에 대해 서비스를 제공할 때 여러 개의 URL로 리다이렉션할 수 있기
	// 때문에 현재 들어온 요청을 식별할 수 있어야 한다.
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	fmt.Println("Requesting data from server!")
	_, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	// http.Client.Do() 함수를 이용하여 실제로 웹 서버로 요청을 보낸다.
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	// HTTP 데이터를 받아서 화면에 출력한다.
	io.Copy(os.Stdout, response.Body)
}
