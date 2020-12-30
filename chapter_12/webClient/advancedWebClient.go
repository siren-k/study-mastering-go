package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		return
	}

	URL, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println("Error in parsing:", err)
		return
	}

	c := &http.Client{Timeout: 15 * time.Second}
	// http.NewRequest() 함수는 인수로 전달한 메소드와 URL과 옵션에 대해 http.Request 오브젝트를 반환한다.
	request, err := http.NewRequest("GET" /* http.MethodGet */, URL.String(), nil)
	if err != nil {
		fmt.Println("Get:", err)
		return
	}

	// http.Do() 함수는 http.Client를 이용해 HTTP 요청(Request)를 보내고 HTTP 응답(Response)를 받는다.
	// 따라서 http.Do()는 앞에서 본 http.Get()이 하는 일을 좀 더 포괄적으로 처리한다.
	httpData, err := c.Do(request)
	if err != nil {
		fmt.Println("Error in Do():", err)
		return
	}

	fmt.Println("Status code:", httpData.Status)
	header, _ := httputil.DumpResponse(httpData, false)
	fmt.Println(string(header))

	contentType := httpData.Header.Get("Content-Type")
	characterSet := strings.SplitAfter(contentType, "charset=")
	if len(characterSet) > 1 {
		fmt.Println("Character Set:", characterSet[1])
	}

	if httpData.ContentLength == -1 {
		fmt.Println("ContentLength is unknown!")
	} else {
		fmt.Println("ContentLength:", httpData.ContentLength)
	}

	length := 0
	var buffer [1024]byte
	r := httpData.Body
	for {
		n, err := r.Read(buffer[0:])
		if err != nil {
			fmt.Println(err)
			break
		}
		length = length + n
	}
	fmt.Println("Calculated response data length:", length)
}
