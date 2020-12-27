package main

import "fmt"

/*
 * 버퍼 채널(Buffer Channel)(버퍼를 사용하는 채널)이란?
 * ==> Go 스케줄러에서 더 많은 요청을 처리할 수 있도록 작업을 큐에 재빨리 저장할 때
 *     이 타입의 채널을 사용한다. 또한 버퍼 채널을 '세마포어(Semaphore)' 처럼
 *     사용해 애플리케이션의 처리량을 제한할 수도 있다.
 *
 * 들어온 요청은 모두 채널로 전달되고, 각각을 하나씩 처리한다. 채널이 어떤 요청에 대한
 * 처리 작업을 끝내면 호출한 측에서 새로운 작업을 처리할 준비가 됐다는 메시지를 보낸다.
 * 따라서 채널에서 사용하는 버퍼의 크기에 따라 동시에 처리할 수 있는 요청의 수가 결정된다.
 */
func main() {
	/*
	 * numbers 채널은 최대 다섯 개의 정수를 저장하도록 설정되었다.
	 */
	numbers := make(chan int, 5)
	counter := 10

	/*
	 * numbers 채널에 10개의 정수를 집어넣고 있다. 하지만 numbers 채널의 용량은
	 * 5개의 정수만 가질 수 있기 때문에, 여기서 지정한 10개의 정수를 모두 저장할 수 없다.
	 */
	for i := 0; i < counter; i++ {
		select {
		case numbers <- i:
		default:
			fmt.Println("Not enough space for", i)
		}
	}

	/*
	 * for 루프와 select 문을 통해 numbers 채널에 담긴 내용을 읽고 있다. numbers 채널이
	 * 비었다면 default 브랜치가 실행된다.
	 */
	for i := 0; i < counter+5; i++ {
		select {
		case num := <-numbers:
			fmt.Println(num)
		default:
			fmt.Println("Nothing more to be done!")
			break
		}
	}
}
