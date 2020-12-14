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

	log.Fatal(sysLog) // 정말 나쁜 일이 발생해서 상황을 알려주자마자 프로그램을 종료하고 싶을 때 사용
	fmt.Println("Will you see this?")
}
