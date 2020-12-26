package main

import "fmt"

// 두 개의 필드를 가진 구조체를 twoInts란 이름으로 정의
type twoInts struct {
	X int64
	Y int64
}

/*
 * regularFunction()이란 이름의 함수를 새로 정의했다. 이 함수는 twoInts 타입의 매개변수
 * 두 개를 받아서 하나의 twoInts 값을 리턴한다.
 */
func regularFunction(a, b twoInts) twoInts {
	temp := twoInts{X: a.X + b.X, Y: a.Y + b.Y}
	return temp
}

/*
 * method() 함수는 regularFunction() 함수와 같다. 하지만 method() 함수는 타입 메소드이기
 * 때문에 호출하는 방법이 좀 다르다.
 */
func (a twoInts) method(b twoInts) twoInts {
	temp := twoInts{X: a.X + b.X, Y: a.Y + b.Y}
	return temp
}

/*
 * Go 언어에서 '타입 메소드(Type Method)'란 특수한 수신자(Receiver) 인수를 받는 함수다.
 * 타입 메소드는 일반 함수처럼 선언하되 함수 이름 앞에 특수한 매개변수가 더 나온다. 이 매개변수
 * 는 함수와 여기에 추가한 매개변수의 타입을 연결한다. 따라서 이러한 매개변수를 '메소드의 수신자
 * (Receiver of the method)'라 부른다.
 *
 * 타입 메소드는 인터페이스와도 관련이 있다.
 */
func main() {
	i := twoInts{X: 1, Y: 2}
	j := twoInts{X: -5, Y: -2}
	fmt.Println(regularFunction(i, j))
	fmt.Println(i.method(j))
}
