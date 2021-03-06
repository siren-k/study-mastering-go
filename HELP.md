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
    슬라이스로 표현하며, 패스워드 값에 현재 시스템 시각 또는 날짜와 결합한다.
> * queue.go에서 정수 대신 부동 소수점 수를 저장할 수 있도록 수정한다.
> * stack.go에서 각 노드에 세 개의 정수형 데이터 필드를 가지도록 수정한다. 각 필드의 이름을 Value, Number,
    Seed로 짓는다. 그러면 Nodestruct를 정의하는 외형적인 부분 말고, 나머지 코드에서 어떤 점이 크게 달라지는지 살펴본다.
> * linkedList.go에서 연결 리스트의 노드를 항상 정렬된 상태를 유지하도록 수정한다.
> * 마찬가지로, doublyLinkedList.go에서 리스트를 항상 정렬된 상태를 유지하도록 수정한다. 또한, 기존 노드를
    삭제하는 함수를 작성한다.
> * hashTableLookup.go에서 해시 테이블의 값이 중복되지 않도록 변경한다. 이 때, lookup() 함수를 활용한다.
> * generatePassword.go에서 대문자로만 구성된 패스워드를 생성하도록 변경한다.
> * conHeap.go에서 float32 대신=, 좀 더 커스텀 구조체를 지원하도록 변경한다.
> * linkedList.go에 빠진 노드 삭제 기능을 구현한다.
> * queue.go를 이중 연결 리스트로 구현하면 더 나아질까? 직접 구현해보기 바란다.

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
    다른 하나는 이름 있는 리턴값을 사용하지 않도록 정의한다. 둘 중 어느 것이 나은가?
> * htmlT.go에서 html/template 대신 text/template을 사용하도록 수정한다.
> * htmlT.go에서 SQLite3 데이터베이스와 통신하는데 https://github.com/feyeleanor/gosqlite3이나 
    https://github.com/phf/go-sqlite3 패키지 중 하나를 사용하도록 수정한다.
> * htmlT.go처러 MySQL 데이터베이스로부터 데이터를 읽는 프로그램을 작성한다. htmlT.go과 다른 부분을 나열한다.

---
# _7장 리플렉션과 인터페이스_

---
##_요약_
> 이 장에서는 Go 언어에서 일종의 계약서 역할을 하는 인터페이스와 타입 메소드, 타입 어써션, 리플랙션에 대해 배웠다.
> 또한 Go 언어에서 OOP 방식으로 코드를 작성하는 방법도 배웠다.
>
> 리플렉션은 Go 언어에서 제공하느 굉장히 강력한 기능이지만, 프로그램 속도가 느려질 수 있다. 실행 시간에 하나의
> 계층이 더 추가되기 때문이다. 또한 리플렉션을 조심해서 사용하지 않으면 프로그램이 뻗어버릴 수 있다.
>
> 이 장에서 가장 중요한 사실 하나만 고른다면, Go 언어가 OOP 언어는 아니지만, 자바나 C++과 같은 전통 OOP 언어에서
> 제공하는 기능을 부분적으로 흉내는 낼 수 있다는 점이다. 다시 말해 프로젝트 전체를 OOP 방식으로 소프트웨어를 개발하고
> 싶다면, Go가 아닌 다른 언어를 사용하는 것이 좋다. 하지만 OOP를 만병통치약처럼 생각하면 안 된다. Go와 같은
> 프로그래밍 언어로도 충분히 성능 좋고 깔끔하고 견고하게 프로그램을 설계할 수 있다.

---
##_참고자료_
#### * 표준 Go 패키지인 reflect에 대한 [공식 문서](https://golang.org/pkg/reflect) 이 패키지는 이 장에서 소개한 것보다 훨씬 많은 기능을 제공한다.
#### * 리플렉션이 정말 좋아서 더 많은 내용을 알고 싶다면, 미첼 하시모토(Mitchell Hashimoto)의 [리플렉트워크(Reflectwalk) 라이브러리](https://github.com/mitchellh/reflectwalk) 를 참고하기 바란다. reflectwalk 라이브러리를 이용하면 리플렉션을 이용해 복잡한 값을 탐색할 수 있다. 여유가 된다면 이 라이브러리를 구현한 코드도 참고한다.

---
##_연습문제_
> * 인터페이스를 직접 정의하고, 이를 다른 프로그램에서 사용해본다. 그러고 나서 직접 작성한 인터페이스를 이용하면
    어떤 점이 좋은지 설명한다.
> * 육면체나 구와 같은 3차원 도형의 부피를 계산하는데 사용할 인터페이스를 정의한다.
> * 선분의 길이나 평면에 있는 두 점 사이의 거리를 계산하는데 사용할 인터페이스를 정의한다.
> * 자신이 작성한 코드에서 리플렉션을 사용해본다.
> * Go 언어의 맵에 대해 리플렉션을 사용해본다.
> * 수학을 잘 한다면, 실수와 복소수에 대한 사칙 연산을 구현하기 위한 인터페이스를 작성한다. 이 때, 표준 Go 타입인
    complex64나 complex128은 사용하지 말고 복소수를 표현하기 위한 구조체를 직접 정의한다.

---
# _8장 유닉스 시스템에게 작업 지시하기_

---
##_요약_
> 8장에서는 파일을 읽고, 파일에 쓰고, flag 패키지를 사용하는 방법을 비롯한 시스템 프로그래밍과 관련된 여러 가지
> 주제를 살펴봤다. 물론 시스템 프로그래밍과 관련해 8장에서 소개하지 못한 내용도 많이 있다. 가령 디렉토리 다루기,
> 파일 복사/삭제/이름 바꾸기, 유닉스 사용자, 그룹, 프로세스 다루는 방법, PATH와 같은 환경 변수 다루는 방법,
> 유닉스 파일 접근 권한, 스파스 파일(Sparse File) 생성 방버, JSON 데이터 읽고 저장하는 방법, 파일 잠그기 및 생성하기,
> 로그 파일 사용하고 순환시키기, os.Stat()에서 리턴하는 구조체에 대한 정보 등이 있다.
> 
> 8장 마지막에서는 두 가지 고급 유틸리티를 Go 언어로 구현하는 방법을 살펴봤다. 첫 번째 유틸리티는 레지스터 상태를
> 검사하는 방법을 소개하고, 두 번째 유틸리티는 시스템 콜을 추적하는 기법을 보여준다.

---
##_참고자료_
#### * io 패키지에 대한 [공식 문서](https://golang.org/pkg/io)
#### * Glot 플로팅 라이브러리에 대해 자세히 알고 싶다면 [공식 웹 페이지](https://github.com/Arafatk/glot) 를 참고한다.
#### * 표준 패키지인 encoding/binary에 대해 자세히 알고 싶다면 [공식 페이지](https://golang.org/pkg/encoding/binary) 를 참고한다.
#### * encoding/gob 패키지에 대한 [공식 문서](https://golang.org/pkg/encoding/gob) 도 참고하기 바란다.
#### * eBPF에 대한 자세한 정보는 [http://www.brendangregg.com/ebpf.html](http://www.brendangregg.com/ebpf.html) 을 참고한다. 또한 [https://www.youtube.com/watch?v=JRFNIKUROPE](https://www.youtube.com/watch?v=JRFNIKUROPE) 와 [https://www.youtube.com/watch?v=w8nFRoFJ6EQ](https://www.youtube.com/watch?v=w8nFRoFJ6EQ) 도 한 번 보기 바란다.
#### * 엔디안(Endian)에 대한 개념을 설명하는 자료는 많다. 그 중 하나가 [위키피디아](https://en.wikipedia.org/wiki/Endianness) 가 있다.
#### * flag 패키지의 [공식 문서](https://golang.org/pkg/flag) 도 참고한다.

---
##_연습문제_
> * 세 개의 인수(텍스트 파일 이름과 스트링 두 개)를 받는 프로그램을 작성한다. 이 프로그램은 인수로 지정한 파일에서
    첫 번째 인수로 스트링이 나타난 자리에 두 번째 스트링으로 교체한다. 보안을 위해 최종 결과는 화면에 출력한다. 다시
    말해 원보 텍스트 파일은 그대로 유지한다.
> * encoding/gob 패키지로 구조체에 대한 슬라이스와 Go 맵을 직렬화하고 역직렬화하는 프로그램을 작성한다.
> * 자신이 원하는 세 가지 시그널을 처리하는 프로그램을 작성한다.
> * 텍스트 파일에 나온 탭 문자를 모두 커맨드라인 인수로 지정한 숫자만큼의 스페이스로 교체하는 프로그램을 작성한다.
    여기서도 결과를 파일이 아닌 화면에 출력한다.
> * 텍스트 파일을 한 줄씩 읽어서 각 줄에 나온 공백 문자를 strings.TrimSpace() 함수로 제거하는 프로그램을 작성한다.
> * kvSaveLoad.go에서 하ㅏ의 커맨드라인 인수만 지원하도록 수정한다. 이 인수는 데이터를 불러오고 저장할 파일의 이름이다.
> * wc(1) 유틸리티를 Go 언어로 구현한다. 여기서 지원하는 커맨드라인 옵션에 대해서는 wc(1)의 맨 페이지를 참고한다.
> * goFind.go에서 일반 파일만 출력하도록 수정한다. 즉, 디렉토리, 소켓, 심볼릭 링크는 출력하지 않는다.
> * Glot으로 함수에 대한 도표를 그리는 프로그램을 작성한다.
> * traceSyscall.go에서 모니터링하는 시스템 콜이 호출될 때마다 출력하도록 수정한다.
> * cat.go에서 파일을 전체를 스캔하지 않고 곧바로 그 내용을 복사하도록 io.Copy(os.Stdout, f)만 수행하도록 수정한다.
> * 파일을 한 단어씩 읽을 때 bufio.NewScanner와 bufio.ScanWords를 사용할 수도 있다. 구체적인 구현 방법을
    찾아서 byWord.go의 새 버전을 생성한다.

---
# _9장 Go 언어의 동시성 - Go 루틴, 채널, 파이프라인_

---
##_요약_
> 9장에서는 Go 언어만이 제공하는 독창적인 기능(Go 루틴, 채널, 파이프라인)에 대해 살펴봤다. 또한 sync 패키지를
> 이용하여 각각의 Go 루틴이 작업을 마칠 때까지 충분한 시간을 두고 기다리는 방법도 배웠다. 마지막으로 Go 함수의
> 매개변수로 채널을 사용하는 방법도 살펴봤다. 이를 이용하면 원하는 형태로 데이터가 흐르는 파이프라인을 구성할 수 있다.

---
##_참고자료_
#### * sync 패키지에 대한 [공식 문서](https://golang.org/pkg/sync)
#### * sync 패키지에 대한 공식 문서를 다시 살펴보자. 이번에는 10장에서 소개할 sync.Mutex와 sync.RWMutex 타입을 중점적으로 살펴본다.

---
##_연습문제_
> * 여러 개의 텍스트 파일을 읽는 파이프라인을 생산하고, 각 파일에서 입력으로 지정한 문구가 나타나는 횟수를 계산한 뒤,
>   모든 파일에서 그 문구가 나타난 총 횟수를 계산하는 프로그램을 작성한다.
> * 주어진 범위 안에 있는 모든 자연수의 제곱을 합하는 파이프라인을 생성한다.
> * simple.go 예제에서 time.Sleep(1 * time.Second) 문장을 삭제하면 어떤 결과가 나오는지 살펴본다.
>   그리고 그 이유를 설명한다.
> * pipeline.go에서 다섯 개의 함수에 대한 파이프라인과 이에 맞는 수만큼 채널을 생성하도록 수정한다.
> * pipeline.go에서 first() 함수의 out 채널을 닫지 않으면 어떤 일이 발생하는지 살펴본다.

---
# _10장 Go 언어의 동시성 - 고급 주제_

---
##_요약_
> 이 장에서는 Go 루틴에 관련된 여러 가지 중요한 주제를 다뤘다. 그 중에서도 특히 select문이 얼마나 강력한지
> 살펴봤다. 또한 표준 Go 패키지인 context를 사용하는 방법도 소개했다.
> 
> select문에서 제공하는 기능을 고려하면, Go 프로그램의 구성 요소를 채널로 연결하는 것이 가장 Go 언어다운 방식이다.
> 
> 동시성 프로그래밍에 관련된 규칙은 없지만 여러 가지가 있지만, 그 중에서도 가장 중요한 것은 특별한 이유가 없다면
> 함부로 공유하지 않는 것이다. 공유 데이터는 동시성 프로그래밍에서 발생할 수 있는 지저분한 버그의 근원이다.
> 
> 10장에서 배운 내용 중에서 반드시 명심할 점이 있다. 같은 프로세스에 있는 여러 스레드가 데이터를 교환하기 위한
> 유일한 수단으로 공유 메모리를 사용하긴 하나, Go 언어에서는 그보다 뛰어난, Go 루틴으로 서로 통신하는 기능을
> 제공한다. 따라서 Go 프로그램에서 공유 메모리를 사용하기 전에 먼저 Go 언어의 방식으로 표현할 수 없는지 고려하자.
> 아무리 생각해도 공유 메모리를 사용할 수 밖에 없다면 공유 메모리보다는 모니터 Go 루틴을 사용하는 편이 좋다.

---
##_참고자료_
#### * sync 패키지에 대한 [공식 문서](https://golang.org/pkg/sync)
#### * context 패키지에 대한 [공식 문서](https://golang.org/pkg/context)
#### * Go 스케줄러의 구현에 대해 자세히 알고 싶다면 [https://golang.org/src/runtime/proc.go](http://golang.org/src/runtime/proc.go) 를 참고한다.
#### * Go 스케줄러에 대한 설계 문서를 보고 싶다면 [https://golang.org/s/go11sched](https://golang.org/s/go11sched) 를 참고한다.

---
##_연습문제_
> * wc(1)의 동시성 버전을 버퍼 채널을 이용해 구현한다.
> * 그런 다음 wc(1)의 동시성 버전을 공유 메모리를 이용해 구현한다.
> * 마지막으로 wc(1)의 동시성 버전을 모니터 Go 루틴을 이용해 구현한다.
> * workerPool.go에서 결과를 파일에 저장하도록 수정한다. 파일을 처리하는 과정에서 뮤텍스와 크리티컬 섹션을
>   사용하거나, 데이터를 디스크에 작성할 때 모니터 Go 루틴을 사용한다.
> * workerPool.go에서 전역 변수인 size의 값을 1로 지정하면 어떤 일이 발생할까? 그리고 왜 그런 일이 생길까?
> * workerPool.go에서 wc(1) 커맨드라인 유틸리티의 기능을 구현하도록 수정한다.
> * workerPool.go에서 버퍼 채널인 clients와 data의 크기를 커맨드라인 인수로 정의할 수 있도록 수정한다.
> * find(1) 커맨드라인 유틸리티의 동시성 버전을 모니터 Go 루틴으로 구현한다.
> * simpleContext.go에서 function1(), function2(), function3() 함수에서 사용한 익명 함수를
>   별도의 함수로 구현하도록 수정한다. 가장 크게 바꿔야 할 부분은 어디인가?
> * simpleContext.go에서 function1(), function2(), function3() 함수가 Context 변수를
>   자체적으로 정의하지 않고, 모두 외부에서 생성된 변수를 사용하도록 수정한다.
> * useContext.go에서 context.WithTimeout() 대신 context.WithDeadline()이나 context.WithCancel()을
>   사용하도록 수정한다.
> * 마지막으로, sync.Mutex 타입의 뮤텍스를 이용해 find(1) 커맨드라인 유틸리티에 대한 동시성 버전을 Go 언어로
>   구현한다.
 
---
# _11장 코드 테스팅, 최적화, 프로파일링_

---
##_요약_
> 11장에서는 코드 테스팅, 최적화, 프로파일링에 대해 살펴봤다. 마지막 부분에서는 실행되지 않는 코드 영역을 찾고
> 크로스 컴파일하는 방법에 대해 배웠다. 그리고 go test 커맨드를 이용해 Go 코드를 테스트하고 벤치마킹하는 방법과
> 예제 함수를 통해 문서에 부가 정보를 추가하는 방법도 소개했다.
> 
> Go 프로파일러와 go tool trace의 완성도는 그리 높지 않지만 프로파일링과 코드 트레이싱(추적)에 대한 개념은
> 반드시 알아두는게 좋다. 새로운 기법을 익히기 위한 가장 좋은 방법은 직접 사용해보는 것이다.

---
##_참고자료_
#### * [그래피즈 웹사이트](http://graphviz.org)
#### * testing 패키지에 대한 [문서](https://golang.org/pkg/testing/)
#### * godoc 유틸리티에 대한 [문서](https://godoc.org/golang.org/x/tools/cmd/godoc)
#### * 표준 Go 패키지인 runtime/pprof에 대한 [문서](https://golang.org/pkg/runtime/pprof.go)
#### * net/http/pprof 패키지의 [소스 코드](https://golang.org/src/net/http/pprof/pprof.go)
#### * net/http/pprof 패키지에 대한 [문서](https://golang.org/pkg/net/http/pprof/)
#### * [Go 1.10 버전과 1.9 버전 사이의 차이점](https://golang.org/doc/go1.10)
#### * pprof 도구에 대한 [페이지](https://github.com/google/pprof)
#### * 고퍼폰(GopherCon) 2017에서 미첼 하시모토(Mitchell Hashimoto)가 발표한 [고급 테스팅 기법 동영상](https://www.youtube.com/watch?v=8hQG7QlcLBk)
#### * testing 패키지에 대한 [소스 코드](https://golang.org/src/testing/testing.go)
#### * profile 패키지에 대한 [웹 페이지](https://github.com/pkg/profile)
#### * go fix 도구에 대한 [웹 페이지](https://golang.org/cmd/fix/)

---
##_연습문제_
> * 8장, '유닉스 시스템 프로그래밍'에서 만든 byWord.go에 대한 테스트 함수를 작성한다.
> * 8장, '유닉스 시스템 프로그래밍'에서 만든 readSize.go에 대한 벤치마크 함수를 작성한다.
> * documentMe.go와 documentMe_test.go에서 발생한 문제를 해결한다.
> * go tool pprof 유틸리티의 텍스트 인터페이스를 이용해 profileMe.go에서 생성된 memoryProfile.out
>   파일을 분석한다.
> * 그런 다음, go tool pprof 유틸리티의 웹 인터페이스를 이용해 profileMe.go에서 생성된 memoryProfile.out
>   파일을 분석한다.

---
# _12장 Go 언어를 이용한 네트워크 프로그래밍의 기초_

---
##_요약_
> 이 장에서는 Go 언어로 웹 클라이언트와 웹 서버를 구현하는 방법뿐만 아니라 웹 사이트를 만드는 방법도 살펴봤다.
> 또한 http.Response, http.Request, http.Transport 구조체를 이용해 HTTP 연결에 대한 매개변수를
> 정의하는 방법도 소개했다.
> 
> 또한 Go 프로그램에서 유닉스 머신의 네트워크 설정 정보를 가져오고, DNS를 조회하고, 도메인에 대한 NS와 MX 레코드를
> 가져오는 방법도 살펴봤다.
> 
> 마지막으로 와이어샤크와 티샤크에 대해서도 간략히 소개했다. 네트워크 트래픽을 수집하고 분석하는 유틸리티 중에서도
> 대표적으로 손꼽히는 도구다. 또한 이 장의 서두에서 nc(1) 유틸리티 사용법도 간략히 소개했다.

---
##_참고자료_
#### * 아피치 웹 서버의 [공식 웹 페이지](http://httpd.apache.org/)
#### * 엔진엑스(Nginx) 웹 서버의 [공식 웹 페이지](http://nginx.org/)
#### * 인터넷과 TCP/IP와 다양한 서비스에 대핸 좀 더 알고 싶다면 RFC 문서부터 읽는 것이 좋다. 이러한 문서를 제공하는 사이트 중 하나로 [http://www.rfc-archive.org](http://www.rfc-archive.org) 를 추천한다.
#### * 와이어샤크와 티샤크 [홈페이지](https://www.wireshark.org/) 도 참고한다.
#### * 표준 Go 패키지인 net의 [홈페이지](https://goloang.org/pkg/net/) 도 참고한다. 공식 Go 문서 중에서도 내용이 가장 방대하다.
#### * 표준 Go 패키지인 net/http의 [홈페이지](https://goloang.org/pkg/net/http/) 도 참고한다.
#### * Go 코드를 하나도 작성하지 않고 웹 사이트를 만들고 싶다면 휴고(Hugo) 유틸리티를 써보기 바란다. 휴고는 Go 언어로 구현한 것이다. 자세한 사항은 [https://gohugo.io/](https://gohugo.io/) 를 참고한다. 그런데 휴고가 정작 유용한 부분은 그 기능보다는 휴고를 구현한 [코드](https://github.com/gohugoio/hugo) 에 있다. Go 프로그래머 입장에서 꼭 한 번 읽어볼만 하다.
#### * net/http/httptrace 패키지에 대한 [공식 문서](https://golang.org/pkg/net/http/httptrace) 도 참고한다.
#### * net/http/pprof 패키지에 대한 [공식 문서](https://golang.org/pkg/net/http/pprof) 도 참고한다.
#### * nc(1) 커맨드라인 유틸리티에서 제공하는 기능과 다양한 옵션에 대해 자세히 알고 싶다면 맨페이지를 읽어보기 바란다.
#### * 데이브 체니(Dave Cheney)가 개발한 httpstat 유틸리티는 [https://github.com/davecheney/httpstat](https://github.com/davecheney/httpstat) 에서 볼 수 있다. net/http/httptrace 패키지로 HTTP 트레이싱을 구현하는 좋은 예도 나와 있다.
#### * ab(1)에 대한 자세한 사항은 [메뉴얼 페이지](https://httpd.apache.org/docs/2.4/programs/ab.html) 를 참고한다.

---
##_연습문제_
> * 앞에 나온 코드를 보지 않고 Go 언어로 웹 클라이언트를 작성한다.
> * MXrecords.go와 NSrecords.go를 합쳐서 두 기능을 모두 제공하는 하나의 유틸리티로 만든다. 각 기능은
    커맨드라인 인수로 지정한다.
> * MXrecords.go와 NSrecords.go에서 IP 주소를 입력받을 수 있도록 수정한다.
> * HTML 출력을 외부 파일에 저장할 수 있도록 advancedWebClient.go를 수정한다.
> * ab(1)의 간략한 버전을 Go 루틴으로 직접 구현한다.
> * 원본 키-값 스토어에 있는 DELETE와 LOOKUP 연산을 제공하도록 kvWeb.go를 수정한다.
> * io.Copy(os.Stdout, response.Body) 문장을 실행하지 않는 플래그를 추가하도록 httpTrace.go를
>   수정한다.
 
---
# _13장 네트워크 프로그래밍 - 서버와 클라이언트 만들기_

---
##_요약_
> 이 장에서는 TCP와 UDP 프로토콜에 대한 서버와 클라이언트를 구현하는 방법을 비롯한 다양한 주제를 살펴봤다.
> 이러한 애플리케이션은 TCP/IP 네트워크를 통해 작동했다.

---
##_참고자료_
#### * 표준 Go 패키지인 net의 [공식 문서](https://golang.org/pkg/net.) 는 Go 언어에 관련된 공식 문서 중에서 가장 방대한 문서다.
#### * 이 책에서 RPC에 대해 소개했지만 gRPC는 다루지 않았다. gRPC는 오픈 소스로 제공되는 고성능 RPC 프레임워크이다. gRPC에 대한 Go 언어로 작성된 패키지는 [https://github.com/grpc/grpc-go](https://github.com/grpc/grpc-go) 에서 볼 수 있다.
#### * IPv4에 대한 ICMP 프로토콜은 RFC 792에 정의되어 있다. 이 문서는 여러 곳에서 제공하고 있는데 그 중 하나로 [https://tools.ietf.org/html/rfc792](https://tools.ietf.org/html/rfc792) 를 추천한다.
#### * 웹 소켓(Web Socket)이란 클리이언트와 원격 호스트가 양방향 통신을 하기 위한 프로토콜이다. Go 언어로 구현된 웹 소켓은 [http://github.com/gorilla/websocket](http://github.com/gorilla/websocket) 에서 볼 수 있다. 웹 소켓에 대한 자세한 사항은 [http://www.rfc-editor.org/rfc/rfc6455.txt](http://www.rfc-editor.org/rfc/rfc6455.txt) 를 참고한다.
#### * 네트워크 프로그래밍에 관심이 많고 TCP 원본 패킷을 다루는 방법을 알고 싶다면 [gopacket](https://github.com/google/gopacket) 라이브러리를 참고하면 유용한 정보를 많이 볼 수 있다.
#### * [https://github.com/mdlayher/raw](https://github.com/mdlayher/raw) 에서 제공하는 raw 패키지를 사용하면 네트워크 디바이스의 디바이스 드라이버 수준에서 데이터를 읽고 쓸 수 있다.

---
##_연습문제_
> * FTP 클라이언트를 Go 언어로 구현한다.
> * 그런 다음 Go 언어로 FTP 서버를 작성한다. FTP 클라이언트와 FTP 서버 중 어느 것이 구현하기 더 힘든가? 
>   또 그 이유는 무엇인가?
> * nc(1) 유틸리티를 Go 언어로 작성한다. 이렇게 복잡한 유틸리티를 작성하기 위한 비법 한 가지를 알려주면,
>   원본 유틸리티에서 가장 기본적인 기능부터 개발한 다음, 여러 가지 옵션을 추가해 나간다.
> * TCPserver.go에서 한 패킷은 날짜를 리턴하고 다른 패킷은 시각을 리턴하도록 수정한다.
> * TCPserver.go에서 여러 클라이언트를 순차적으로 처리하고록 수정한다. 주의할 점은 여러 요청을 동시에
>   처리할 때와 다른 방식으로 구현해야 한다. 쉽게 말해, for 루트를 사용해 Accept()를 여러 번 호출하도록
>   작성해야 한다.
> * fiboTCP.go와 같은 TCP 서버는 특정한 시그널을 받으면 종료한다. 따라서 8장, '유닉스 시스템 프로그래밍'에서
>   설명한 내용을 참고해 fiboTCP.go에서 주어진 시그널을 처리하는 코드를 추가한다.
> * kvTCP.go에서 save() 함수를 sync.Mutex로 보호하도록 수정한다. 그리고 꼭 이렇게 할 필요가 있는지
>   설명한다.
> * Go 언어로 조그만 웹 서버를 직접 구현한다. 이 때, http.ListenAdnServe() 함수 대신 일반 TCP로
>   구현한다.
 
---
# _그 이상을 알고 싶다면?_

---
> 어떠한 프로그래밍 책도 완벽할 수 없으며, 이 책도 예외가 아니다. 지금까지 다룬 주제 중에서 빠진 부분은
> 없을까? 당연히 있다. 왜 그럴까? 이 책에서 다룬 것보다 훨씬 많은 주제들이 있기 때문이다. 따라서 모든 
> 주제를 다루다보면 책을 출간할 수 조차 없을 것이다. 이는 프로그램의 명세서를 작성하는 것과 비슷하다. 추가하고
> 싶은 새롭고 멋진 기능은 끝없이 존재하기 마련이다. 따라서 중간에 한 번은 매듭짓지 않으면 영원히 프로그램을
> 개발할 수도, 최종 사용자에게 전달될 수도 없을 것이다. 여기서 다루지 못한 주제 중 어떤 것은 이 책의 2판에서
> 추가될 지도 모른다.
> 
> 한 가지 다행인 것은 이 책을 읽고 나면 어떤 주제든 직접 찾아서 익힐 역략을 갖게 된다는 것이다. 제대로 집필된
> 책이라면 어떠한 프로그래밍 책을 봐도 이러한 효과를 얻을 수 있다. 이 책의 주된 목적은 독자들로 하여금 Go 언어로
> 프로그램을 작성하고, 그 과정에서 여러 가지 경험을 쌓게 하는 것이다. 이를 위해 실습하면서 시행착오를 겪어보는
> 것만큰 좋은 방법은 없다. 프로그래밍 언어를 배우기 위한 유일한 방법이기도 하다. 따라서 보다 어려운 기능을
> 끊임없이 개발해보기 바란다.
> 
> 이로써 또 하나의 Go 언어 책이 끝았다. 하지만 독자에게는 이제부터가 시작이다. 이제 본격적으로 Go 언어로 소프트웨어를
> 작성하고 새로운 주제도 더 많이 익혀보기 바란다.