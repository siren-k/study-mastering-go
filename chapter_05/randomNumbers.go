package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func randomNumber(min, max int) int {
	return rand.Intn(max-min) + min
}

/*
 * 난수 생성에 대해 관심이 많은 독자는 도널드 E. 커누스(Donald E. Knuth) 교수의
 * 'The Art of Computer Programming(Addison-Wesley Professional, 2011)'의 2권부터
 * 읽어보기 바란다.
 *
 * Go 언어에서 보안에 안전하게 난수를 생성하고 싶다면, crypto/rand 패키지를 사용한다. 이 패키지는 암호학
 * 관점에서 안전한 의사난수 생성기(Cryptographically Secure Pseudorandom Number Generator)를
 * 구현한 것이다. crypto/rand 패키지에 대한 자세한 사항은 공식 문서(https://golang.org/pkg/crypto/rand)를 참고한다.
 */
func main() {
	MIN := 0
	MAX := 100
	TOTAL := 100
	SEED := time.Now().Unix()

	arguments := os.Args
	/*
	 * switch 블록이 구현하려는 로직은 다소 간단한다. 커맨드라인 인수의 개수에 따라 빠진
	 * 인수에 대한 초기값을 지정하거나, 사용자가 지정한 값으로 초기화한다. 간결한 구성을 위해
	 * strconv.Atoi()와 strconv.ParseInt() 함수의 error 변수는 무시하도록
	 * 언더스코어 문자를 지정했다. 실전에서는 strconv.Atoi()와 strconv.ParseInt()
	 * 함수의 error 변수를 반드시 처리해야 한다.
	 *
	 * 마지막으로 strconv.ParseInt()를 이용하여 SEED 변수에 새로운 값을 지정한 이유는,
	 * rand.Seed() 함수가 매개변수를 int64 타입으로 받기 때문이다. strconv.ParseInt()의
	 * 첫 번째 매개변수는 파싱할 스트링을 지정하고, 두 번째 매개변수는 생성된 숫자의 진수(기저(Base))를
	 * 지정하고, 세 번째 매개변수는 생성된 숫자의 비트 크기를 지정한다.
	 *
	 * 예제에서는 64비트를 사용하는 십진수 정수를 생성하므로 두 번쨰 매개변수를 10으로 지정하고,
	 * 세 번째 매개변수를 64로 지정했다. 이 때 부호가 없는 정수를 파싱하려면 strconv.ParseUint()
	 * 함수를 사용해야 한다는 점을 주의하자
	 */
	switch len(arguments) {
	case 2:
		fmt.Println("Usage: ./randomNumbers MIN MAX TOTAL SEED")
		MIN, _ := strconv.Atoi(arguments[1])
		MAX = MIN + 100
	case 3:
		fmt.Println("Usage: ./randomNumbers MIN MAX TOTAL SEED")
		MIN, _ = strconv.Atoi(arguments[1])
		MAX, _ = strconv.Atoi(arguments[2])
	case 4:
		fmt.Println("Usage: ./randomNumbers MIN MAX TOTAL SEED")
		MIN, _ = strconv.Atoi(arguments[1])
		MAX, _ = strconv.Atoi(arguments[2])
		TOTAL, _ = strconv.Atoi(arguments[3])
	case 5:
		fmt.Println("Usage: ./randomNumbers MIN MAX TOTAL SEED")
		MIN, _ = strconv.Atoi(arguments[1])
		MAX, _ = strconv.Atoi(arguments[2])
		TOTAL, _ = strconv.Atoi(arguments[3])
		SEED, _ = strconv.ParseInt(arguments[4], 10, 64)
	default:
		fmt.Println("Using default values!")
	}

	rand.Seed(SEED)
	for i := 0; i < TOTAL; i++ {
		myrand := randomNumber(MIN, MAX)
		fmt.Print(myrand)
		fmt.Print(" ")
	}
	fmt.Println()

	// ❯ go run randomNumbers.go
	// Using default values!
	// 90 92 3 76 41 85 56 56 31 50 6 21 27 59 50 29 76 11 12 14 88 4 66 99 89 23 59 66 17 97 50 65 78 27 61 47 25 49 21 22 66 95 22 14 47 15 78 15 94 75 33 99 62 14 95 44 35 16 26 42 7 27 94 11 38 12 5 6 25 5 95 35 21 43 79 29 59 76 94 22 18 50 30 91 56 13 73 48 22 58 6 43 17 45 99 63 78 0 24 99
	// ❯ go run randomNumbers.go 1 3 2
	// Usage: ./randomNumbers MIN MAX TOTAL SEED
	// 2 1
	// ❯ go run randomNumbers.go 1 5 10 10
	// Usage: ./randomNumbers MIN MAX TOTAL SEED
	// 3 1 4 4 1 1 4 4 4 3
	// ❯ go run randomNumbers.go 1 5 10 10
	// Usage: ./randomNumbers MIN MAX TOTAL SEED
	// 3 1 4 4 1 1 4 4 4 3
}
