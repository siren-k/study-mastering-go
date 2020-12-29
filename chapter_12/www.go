package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// 첫 번째 핸들러 함수 구현
// 핸들러 함수(Handler Function)는 설정에 따라 한 개 이상의 URL에 대해 서비스 제공
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

// 두 번째 핸들러 함수 구현
// 컨텐츠를 동적으로 생성
func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

func main() {
	// 기본 포트 8001
	PORT := ":8001"
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Using default port number: ", PORT)
	} else {
		PORT = ":" + arguments[1]
	}

	// http.HandleFunc() 함수는 URL과 핸들러 함수를 연결
	// 매칭되는 URL이 없을 경우 myHandler를 연결한다.
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/", myHandler)

	// http.ListenAndServe() 함수에 원하는 포트를 지정해서 호출하면 웹 서버가 구동된다.
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
