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