package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	myUrl string
	delay int = 5
	w     sync.WaitGroup
)

// 웹 서버의 응답을 담을 구조체
type myData struct {
	r   *http.Response
	err error
}

func connect(c context.Context) error {
	/*
	 * HTTP 연결
	 */
	defer w.Done()
	data := make(chan myData, 1)

	tr := &http.Transport{}
	httpClient := &http.Client{Transport: tr}

	req, _ := http.NewRequest("GET", myUrl, nil)

	go func() {
		response, err := httpClient.Do(req)
		if err != nil {
			fmt.Println(err)
			data <- myData{nil, err}
			return
		} else {
			pack := myData{response, err}
			data <- pack
		}
	}()

	select {
	case <-c.Done():
		tr.CancelRequest(req)
		<-data
		fmt.Println("The request was cancelled")
		return c.Err()
	case ok := <-data:
		err := ok.err
		resp := ok.r
		if err != nil {
			fmt.Println("Error select:", err)
			return err
		}
		defer resp.Body.Close()

		realHTTPData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error select:", err)
			return err
		}
		fmt.Printf("Server Response: %s\n", realHTTPData)
	}
	return nil
}

/*
 * useContext.go 프로그램은 두 개의 커맨드라인 인수를 받는다. 하나는 연결할 서버의 URL이고
 * 다른 하나는 이 서버의 응답을 기다릴 시간이다. 이 프로그램에 인수 하나만 지정하면 지연 시간을
 * 5초로 지정한다.
 */
func main() {
	if len(os.Args) == 1 {
		fmt.Println("Need a URL and a delay!")
		return
	}

	myUrl = os.Args[1]
	if len(os.Args) == 3 {
		t, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		delay = t
	}

	fmt.Println("Delay:", delay)
	c := context.Background()
	// 만료 주기는 context.WithTimeout() 함수로 정의했다.
	c, cancel := context.WithTimeout(c, time.Duration(delay)*time.Second)
	defer cancel()

	fmt.Printf("Connecting to %s\n", myUrl)
	w.Add(1)
	// connect() 함수는 Go 루틴으로 실행되며, 정상적으로 종료되거나 cancel() 함수가 호출되면서 종료된다.
	go connect(c)
	w.Wait()
	fmt.Println("Exiting...")
}
