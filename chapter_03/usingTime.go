package main

import (
	"fmt"
	"time"
)

func main() {
	// 유닉스 에포크 시간(Unix Epoch Time)을 반환
	// 1970년 1월 1일 UTC 시각으로 00:00:00부터 경과된 시간을 초단위로 표현한 숫자
	fmt.Println("Epoch time:", time.Now().Unix())
	t := time.Now()
	// time.Format() 함수를 사용하면 time 변수를 다른 포맷으로 변환할 수 있다.
	fmt.Println(t, t.Format(time.RFC3339))
	fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year())

	// 지연 시간을 흉내내기 위한 용도로 time.Sleep() 함수를 사용한다.
	// 10초의 경과 시간을 표현하고 싶다면 time.Second에 10을 곱하면 된다.
	// 비슷한 상수로 time.Nanosecond, time.Microsecond, time.Millisecond, time.Minute, time.Hour 등이 있다.
	time.Sleep(time.Second)
	t1 := time.Now()
	// time.Sub() 함수는 두 시간 사이의 차를 알려준다.
	fmt.Println("Time Difference:", t1.Sub(t))

	// time.Format()을 이용해 새로운 날짜 포맷을 정의한다.
	formatT := t.Format("01 January 2006")
	fmt.Println(formatT)
	loc, _ := time.LoadLocation("Europe/Paris")
	londonTime := t.In(loc)
	fmt.Println("Paris:", londonTime)
}
