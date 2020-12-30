package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is:"
	fmt.Fprintf(w, `<h1 align="center">%s</h1>`, Body)
	fmt.Fprintf(w, `<h2 align="center">%s</h2>`, t)
	fmt.Fprintf(w, `Serving: %s\n`, r.URL.Path)
	fmt.Printf("Served time for %s\n", r.Host)
}

func main() {
	PORT := ":8081"
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Printf("Listening on http://0.0.0.0%s\n", PORT)
	} else {
		PORT = ":" + arguments[1]
		fmt.Printf("Listening on http://0.0.0.0%s\n", PORT)
	}

	m := http.NewServeMux()
	// http.Server 구조체는 ReadTimeout과 WriteTimeout을 제공한다.
	// ReadTimeout 필드의 값은 바디를 포함한 요청 전체를 읽는데 걸리는 최대 시간을 설정한다.
	// WriteTimeout 필드는 응답을 쓰는데 허용된 최대 시간을 설정한다. 쉽게 말해 요청 헤더를
	// 읽는 작업이 끝난 시점부터 응답을 쓰는 작업이 끝날 때까지의 시간이다.
	srv := &http.Server{
		Addr:         PORT,
		Handler:      m,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
	m.HandleFunc("/time", timeHandler)
	m.HandleFunc("/", myHandler)

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}

	// time nc localhost 8081
	// ==> 어떤 커맨드도 실행하지 않았기 때문에 HTTP 서버는 연결을 종료한다.
	//     time(1) 유틸리티의 결과를 통해 서버가 연결을 끊을 떄까지 걸린 시간을 확인할 수 있다.
}
