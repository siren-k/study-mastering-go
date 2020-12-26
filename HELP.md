# _2장 Go 언어의 내부 살펴보기_

---
## _요약_

---
> 2장에서는 Go 언어와 관련된 여러 가지 흥미로운 주제들을 살펴봤다. Go 언어의 가비지 컬렉에
> 대해 이론적인 관점과 실용적인 관점에서 살펴봤고, Go 프로그램에서 C 코드를 호출하는 방법도
> 배웠다. 간편하지만 다소 까다로운 defer 키워드에 대해서도 알아봤고, panic()과 recover() 함수에
> 대해서도 소개했다. 유닉스 도구인 strace(1), dtrace(1), dtruss(1)도 써봤고, 표준 Go 패키지인
> unsafe를 활용하는 방법도 설명했다.
> 
> 2장에서 반드시 기억할 점은, unsafe 패키지와 같은 도구나 Go 코드에서 C 코드를 호울하는 기능은
> 크세 세 가지 용도롤 사용한다는 것이다.
>> 첫째는 Go 코드의 안전성을 조금 희생하는 대신 성능을 최대한 끌어내고 싶을 때다.
> 
>> 둘째는 다름 프로그램밍 언어와 대화하고 싶을 때다
> 
>> 셋째는 Go 언어만으로 구현하기 힘든 것을 만들고 싶을 때다

## _Go 코딩에 참고한 만한 조언_

___
Go 코드를 잘 작성하는데 도움되는 조언을 몇 가지 제시하면 다음과 같다.

>* GO 함수에서 발생한 에러를 로그에 기록하거나 리턴한다. 단, 특별한 이유가 없다면 둘 다 하지 않는다.
>* Go 언에서 interface는 데이터나 데이터 구조가 아니라 동작을 정의하는 것이다.
>* io.Reader나 io.Writer 인터페이스를 사용하면 코드의 확장성을 높일 수 있다.
>* 함수에 변수를 전달할 때 꼭 필요한 경우에만 포인터로 전달한다. 나머지 경우는 변수를 값으로 전달한다.
>* 에러 변수는 string 타입이 아닌 error 타입으로 지정한다.
>* 프로덕션 머신에는 Go 코드를 테스트하지 않는다.
>* Go 언어의 기능 중 잘 이해가지 않는 것은 사용하기 전에 반드시 테스트해본다. 특히 굉장히 많은 사용자들이 쓰게 될 애플리케이션이나 유틸리티를 개발할 때는 꼭 테스트한다.
>* 실수하는 것이 두렵다면 정말로 유용한 것을 만들 수 없다. 그러니 최대한 많이 실험해보기 바란다.

##_참고자료_

---
#### * 표준 Go 패키지인 unsafe에 대한 좀 더 알고 싶다면 [공식문서](http://golang.org/pkg/unsafe) 를 참고한다.
#### * DTrace [웹사이트](http://dtrace.org) 도 방문해본다.
#### * runtime 패키지에서 제공하는 함수에 대한 자세한 정보는 [http://golang.org/pkg/runtime](http://golang.org/pkg/runtime) 을 참고한다.
#### * 쉽지 않겠지만 논문도 읽어두면 큰 도움이 된다. 특히 "On-the-fly Garbage Collection: An Exercise in Cooperation" 만큼은 꼭 한 번 읽기 바란다. 이 논문은 [https://dl.acm.org/citation.cfm?id=359655](https://dl.acm.org/citation.cfm?id=359655) 를 비롯한 여러 곳에서 다운로드할 수 있다.
#### * [https://github.com/gasche/gc-latency-experiment](https://github.com/gasche/gc-latency-experiment) 에 가면 여러 가지 프로그래밍 언어의 가비지 컬렉터에 대한 벤치마킹 코드를 볼 수 있다.
#### * 가비지 컬렉션에 대해 자세히 알고 싶다면 [http://gchandbook.org](http://gchandbook.org) 를 꼭 한 번 가보기 바란다.
#### * cgo에 대한 공식 문서인 [https://golang.org/cmd/cgo](https://golang.org/cmd/cgo) 도 방문하기 바란다.

---

# _3장 Go 언어의 기본 데이터 타입_

---
##_요약_
> 이 장에서는 맵, 배열, 슬라이스, 포인터, 상수, 루프, 날짜와 시간을 다루는 방법과 같이 Go 언어에서 제공하는 여러 가지 흥미로운 기능에 대하
> 살펴봤다. 이제 슬라이스가 배열보다 뛰어난 이유를 이해했을 것이다.
---
##_참고자료_

---
#### * time 패키지에 대한 공식 문서 [http://golang.org/pkg/time](http://golang.org/pkg/time)
#### * 표준 Go 패키지인 regexp에 대한 [공식 문서](http://golang.org/pkg/regexp)

---
##_연습문제_
> * 숫자 4의 제곱에 대한 상수 생성자 iota를 작성한다.
> * 한 주의 요일에 대한 상수 생성자 iota를 작성한다.
> * 배열을 맵으로 변화하는 프로그램을 작성한다.
> * parseTime.go를 원하는 형태로 수정하다. 이때 테스트도 반드시 수행한다.
> * 두 가지의 날짜 및 시간 포맷을 처리하도록 timeDate.go를 수정한다.
> * parseDate.go를 원하는 형태로 수정한다.

---

# _4장 합성 타입 사용법_

---
##_요약_
> 4장에서는 Go 언어에서 제공하는 여러 가지 유용한 기능을 살펴봤다. 구조체, 튜플, 스트링, 룬을 생성하고 사용하는
> 방법도 살펴보고 표준 Go 패키지인 unicode에서 제공하는 기능도 살펴봤다. 또한, 패턴 매칭과 정규 표현식,
> switch문, 표준 Go 패키지인 strings를 사용하는 방법도 소개하고 Go 언어로 간단한 키-값 스토어를 구현하는
> 방법도 알아봤다. 마지막으로 math/big 패키지에서 제공하는 타입을 이용하여 파이 값을 원하는 정확도로
> 계산하는 방법도 소개했다.

---
##_참고자료_

---
#### * 표준 Go 패키지인 regexp에 대한 [공식 문서](https://golang.org/pkg/regexp)
#### * grep(1) 유틸리티에 대한 맨 페이지
#### * math/big 패키지에 대한 [공식 문서](https://golang.org/pkg/math/big)
#### * 표준 Go 패키지인 unicode에 대한 [공식 문서](https://golang.org/pkg/unicode)
#### * 처음 볼 때는 굉장히 힘들겠지만, [The Go Programming Language Specification](https://golang.org/ref/spec) 도 참고하면 도움된다.

---
##_연습문제_
> * 주어진 IPv4 주소에서 올바르게 작성된 부분과 그렇지 않은 부분을 출력하는 Go 프로그램을 작성해본다.
> * make와 new의 차이점에 대해 이 장의 설명을 보지 않고 표현한다.
> * 문자, 바이트, 룬의 차이점을 설명한다.
> * findIPv4.go를 이용하여 로그 파일에서 가장 많이 자주 등장하는 IP 주소를 출력하는 Go 프로그램을 작성한다. 
    결과를 출력할 때 다른 유닉스 유틸리티는 사용하지 않는다.
> * 로그 파일에서 404 HTML 에러 메시지를 생성한 IPv4 주소를 찾는 Go 프로그램을 작성한다.
> * 표준 Go 패키지인 math/big을 사용하여 제곱근을 높은 정확도로 구하는 Go 프로그램을 작성한다.
    여기서 사용할 알고리즘은 마음대로 고른다.

---
# _5장 자료 구조로 Go 코드 개선하기_

---
##_요약_
> 5장에서는 Go 언어의 자료 구조에 관련된 흥미롭고 유용한 주제애 대해 다양하게 살펴봤다. 여기서는 연결 리스트,
> 해시 테이블, 큐, 스택을 구현하는 방법을 소개하고, 표준 Go 패키지인 container에서 제공하는 여러 가지 기능을 
> 살펴봤다. 5장을 읽고 나서 반드시 기억할 점은, 자료 구조마다 노드를 정의하고 구현하는 부분이 가장 중요하다는
> 점이다.

---
##_참고자료_
#### * 자체 정의한 언어로 그래프를 그리는 유틸리티인 [그래피즈(Graphviz)](http://graphviz.org)
#### * 표준 Go 패키지인 container에 대한 [공식 문서](https://golang.org/pkg/container)
#### * 자료 구조에 대해 좀 더 알고 싶은 독자는 알프레드 V. 에이호(Alfred V. Aho), 존 E. 홉크로프트(John E. Hopcroft), 제프리 D. 울만(Jeffrey D. Ullman)의 'The Design and Analysis of Computer Algorithms(Addison-Wesley Professional, 1974)'를 꼭 한번 읽어보기 바란다. 정말 좋은 책이다.
#### * 그 밖에 알고리즘과 자료 구조에 대해 잘 설명하고 있는 책을 추천하면, 존 베틀리(Jonb Bentley)의 'Programming Pearls(Addison-Wesley Professional, 1999)', 동 저자의 'More Programming Pearls(Addison-Wesley Profession, 1988)' 등이 있다. 두 책 모두 읽으면 더욱 훌륭한 프로그래머가 될 수 있다.

---
##_연습문제_ 
> * 패스워드 목록에서 하나를 고르는 방식으로 generatePassword.go의 로직을 변경한다. 이 때, 패스워드 목록은
>   슬라이스로 표현하며, 패스워드 값에 현재 시스템 시각 또는 날짜와 결합한다.
> * queue.go에서 정수 대신 부동 소수점 수를 저장할 수 있도록 수정한다.
> * stack.go에서 각 노드에 세 개의 정수형 데이터 필드를 가지도록 수정한다. 각 필드의 이름을 Value, Number,
>   Seed로 짓는다. 그러면 Nodestruct를 정의하는 외형적인 부분 말고, 나머지 코드에서 어떤 점이 크게 달라지는지 살펴본다.
> * linkedList.go에서 연결 리스트의 노드를 항상 정렬된 상태를 유지하도록 수정한다.
> * 마찬가지로, doublyLinkedList.go에서 리스트를 항상 정렬된 상태를 유지하도록 수정한다. 또한, 기존 노드를
>   삭제하는 함수를 작성한다.
> * hashTableLookup.go에서 해시 테이블의 값이 중복되지 않도록 변경한다. 이 때, lookup() 함수를 활용한다.
> * generatePassword.go에서 대문자로만 구성된 패스워드를 생성하도록 변경한다.
> * conHeap.go에서 float32 대신=, 좀 더 커스텀 구조체를 지원하도록 변경한다.
> * linkedList.go에 빠진 노드 삭제 기능을 구현한다.
> * queue.go를 이중 연결 리스트로 구현하면 더 나아질까? 직접 구현해보기 바란다.

---

---
# _6장 Go 패키지에 대해 잘 알려지지 않은 사실_

---
##_요약_
> 6장에서는 Go 언어의 함수와 Go 언어의 패키지라는 두 가지 주제를 중심으로 살펴봤다. 또한 바람직한 Go 패키지를
> 작성하는데 도움되는 팁도 소개했다. 일반 텍스트와 HTML 출력을 템플릿 기반으로 생성하는 text/template과
> html/template 패키지에 대해서도 살펴봤다. 마지막으로 표준 Go 패키지인 syscall에서 제공하는 고급 기능도
> 알아봤다.

---
##_참고자료_
#### * 표준 Go 패키지인 syscall에 대한 [공식 문서](https://golang.org/pkg/syscall) 도 한 번 읽어본다. 지금껏 내가 본 Go 문서 페이지 중에서 가장 길다.
#### * text/template 패키지에 대한 [공식 문서](https://golang.org/pkg/text/tempate) 도 방문하기 바란다.
#### * 마찬가지로 html/template 패키지에 대한 [공식 문서](https://golang.org/pkg/html/template) 페이지도 방문한다.
#### * [SQLite3 홈페이지](https://www.sqlite.org) 도 가본다.
#### * 맷 라이어(Mat Ryer)의 Writing Beautiful Packages in Go 동영상도 보기 바란다.
#### * 플랜 9(Plan 9)에 대해 자세히 알고 싶으면 [https://plat9.io/plan9](https://plat9.io/plan9) 을 참고하기 바란다.
#### * find(1) 커맨드라인 도구의 맨 페이지(man 1 find)를 찬찬히 읽어보기 바란다.

---
##_연습문제_
> * fmt.Println() 함수의 실제 구현 코드를 좀 더 자세히 살펴본다.
> * 세 개의 int 값을 정렬하는 함수를 작성한다. 이 함수를 두 가지 버전으로 구현한다. 하나는 이름이 있는 리턴값으로
>   다른 하나는 이름 있는 리턴값을 사용하지 않도록 정의한다. 둘 중 어느 것이 나은가?
> * htmlT.go에서 html/template 대신 text/template을 사용하도록 수정한다.
> * htmlT.go에서 SQLite3 데이터베이스와 통신하는데 https://github.com/feyeleanor/gosqlite3이나 
>   https://github.com/phf/go-sqlite3 패키지 중 하나를 사용하도록 수정한다.
> * htmlT.go처러 MySQL 데이터베이스로부터 데이터를 읽는 프로그램을 작성한다. htmlT.go과 다른 부분을 나열한다.

---

