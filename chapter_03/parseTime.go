package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var myTime string
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s string\n", filepath.Base(os.Args[0]))
		return
	}

	myTime = os.Args[1]
	// 시와 분을 나타내는 스트링을 파싱하려면, "15:04"와 같이 지정해야 한다.
	// 파싱 작업의 성공 여부는 err 변수의 값을 통해 확인할 수 있다.
	d, err := time.Parse("15:04", myTime)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Hour(), d.Minute())
	} else {
		fmt.Println(err)
	}
}
