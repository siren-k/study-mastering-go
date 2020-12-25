package main

import (
	"fmt"
	s "strings" // strings에 대한 앨리어스(Alias)를 생성
	"unicode"
)

var f = fmt.Printf // f 변수에 fmt.Printf()를 할당

func main() {
	upper := s.ToUpper("Hello there!")
	f("To Upper: %s\n", upper)
	f("To Lower: %s\n", s.ToLower("Hello THERE"))
	f("%s\n", s.Title("tHis wiLL be A title!"))

	// strings.EqualFold() 함수를 사용하면 서로 다른 문자로 구성된 스트링이 서로 같은지 알아낼 수 있다.
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHALis"))
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHALi"))

	// strings.HasPrefix() 함수는 첫 번째 매개변수로 지정한 스트링이 두 번째 매개변수로
	// 지정한 스트링으로 시작하면 true를 리턴하고 그렇지 않으면 false를 리턴한다.
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "Mi"))
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "mi"))

	// strings.HasSuffix() 함수는 첫 번째 매개변수로 지정한 스트링이 두 번째 매개변수로
	// 지정한 스트링으로 끝나면 true를 그렇지 않으면 false를 리턴한다.
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "is"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "IS"))

	f("Index: %v\n", s.Index("Mihalis", "ha"))
	f("Index: %v\n", s.Index("Mihalis", "Ha"))

	// strings.Count() 함수는 첫 번째 매개변수로 전달된 스트링에서 두 번쨰 매개변수가
	// 나타난 횟수를 중첩되지 않게 센다.
	f("Count: %v\n", s.Count("Mihalis", "i"))
	f("Count: %v\n", s.Count("Mihalis", "I"))

	f("Repeat: %s\n", s.Repeat("ab", 5))

	f("TrimSpace: %s\n", s.TrimSpace(" \tThis is a line. \n"))
	f("TrimLeft: %s\n", s.TrimLeft(" \tThis is a\t line. \n", "\n\t "))
	f("TrimRight: %s\n", s.TrimRight(" \tThis is a\t line. \n", "\n\t "))

	// strings.Compare() 함수는 두 개의 스트링을 사전 항목 나열 순으로(Lexicographically) 비교한다.
	// 따라서 두 값이 같으면 0을 아니면 -1이나 +1을 리턴한다.
	f("Compare: %v\n", s.Compare("Mihalis", "MIHALIS"))
	f("Compare: %v\n", s.Compare("Mihalis", "Mihalis"))
	f("Compare: %v\n", s.Compare("MIHALIS", "MiHalis"))

	// strings.Fields() 함수는 스트링 매개변수를 공백 문자를 기준으로 쪼갠다.
	// 이 함수는 unicode.IsSpace() 함수에서 정의하는 공백 문자를 사용한다.
	f("Fields: %v\n", s.Fields("This is a string!"))
	f("Fields: %v\n", s.Fields("Thisis\na\tstring!"))

	// strings.Split() 함수를 이용하면 주어진 스트링을 매개변수로 지정한 구분자(Separator) 스트링을 기준을 쪼갠다.
	// strings.Split() 함수는 스트링 슬라이스를 리턴한다.
	// strings.Split() 함수의 두 번째 매개변수를 ""으로 지정하면 스트링을 문자 단위로 처리할 수 있다.
	f("%s\n", s.Split("abcd efg", ""))

	// strings.Replace() 함수는 네 개의 매개변수를 받는다. 첫 번째 매개변수는 처리하려는 원본 스트링을 지정한다.
	// 두 번째 매개변수는 기존 스트링에서 검색해서 세 번째 매개변수로 대체할 스트링을 지정한다. 마지막 매개변수는 교체할
	// 최대 횟수를 지정한다. 이 값을 음수로 지정하면 무제한으로 교체할 수 있다.
	f("%s\n", s.Replace("abcd efg", "", "_", -1))
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	f("%s\n", s.Replace("abcd efg", "", "_", 2))

	lines := []string{"Line 1", "Line 2", "Line 3"}
	f("Join: %s\n", s.Join(lines, "+++"))

	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "++"))

	// strings.SplitAfter() 함수는 첫 번째 매개변수로 지정한 스트링을 여러 서브스트링으로 나누며
	// 이 때, 두 번째 매개변수로 지정한 스트링이 나온 바로 뒤에서 자른다.
	trimFuncion := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t .", trimFuncion))
}
