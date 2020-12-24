package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	logs := []string{
		"127.0.0.1 - - [16/Nov/2017:10:49:46 +0200] 325504",
		"127.0.0.1 - - [16/Nov/2018:10:16:41 +0200] \"GET /CVEN HTTP/1.1\" 200 12531 \"=\" \"Mozilla/5.0 AppleWebKit/537.36",
		"127.0.0.1 200 9412 - - [12/Nov/2017:06:26:05 +0200] \"GET \"http://www.mtsoukalous.eu/taxonomy/term/47\" 1507",
		"[12/Nov/2017:16:27:21 +0300]",
		"[12/Nov/2017:20:88:21 +0200]",
		"[12/Nov/2017:20:21 +0200]",
	}

	for _, logEntry := range logs {
		// 문장에 담긴 날짜와 시간에 대한 스트링을 찾아서 처리하기 편하기 때문에 정규표현식을 사용하였음
		// regexp.MustCompile() 안에 들어갈 값은 역이용(`, 그레이브 액센트(grave accent)) 부호로 묶는다.
		r := regexp.MustCompile(`.*\[(\d\d/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\].*`)
		if r.MatchString(logEntry) {
			match := r.FindStringSubmatch(logEntry)
			dt, err := time.Parse("02/Jan/2006:15:04:05 -0700", match[1])
			if err == nil {
				// 정규표현식과 일치하는 스트링을 찾았다면, time.Parse()로 파싱해서
				// 올바른 날짜/시간 스트링인지 확인한 후 RFC850 포맷으로 출력한다.
				newFormat := dt.Format(time.RFC850)
				fmt.Println(newFormat)
			} else {
				fmt.Println("Not a valid date time format!")
			}
		} else {
			fmt.Println("Not a match!")
		}
	}
}
