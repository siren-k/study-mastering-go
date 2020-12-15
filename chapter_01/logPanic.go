package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	sysLog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, "Some program!")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}

	// 프로그램이 다시 실행될 수 없을 정도로 오류가 발생하는 순간, 이에 관련된 정보를 최대한 알고 싶을 때 사용
	// log.Fatal()과 마찬가지로 적절한 로그 파일에 항목을 추가한 뒤에 곧바로 프로그램을 종료
	log.Panic(sysLog)
	fmt.Println("Will you see this?")
}
