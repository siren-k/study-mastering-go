package main

import (
	"fmt"
	"os"
	"reflect"
)

type a struct {
	X int
	Y float64
	Z string
}

type b struct {
	F int
	G int
	H string
	I float64
}

/*
 * 리플렉션(Reflection)은 Go 언어에서 제공하는 고급 기능으로서, 주어진 오브젝트에 대한 타입 뿐만
 * 아니라 구조체에 대한 정보를 동적으로 알아내는데 사용된다. Go에서는 이러한 리플렉션 기능을 reflect
 * 패키지를 통해 제공한다. 그런데 리플렉션은 Go 프로그램을 작성할 때마다 항상 쓰는 기능이 아니다.
 * 그렇다면 리플렉션을 언제 그리고 왜 쓰는지 궁금할 것이다.
 *
 * 리플렉션은 fmt, text/template, html/template과 같은 패키지를 구현하는데 사용된다. fmt
 * 패키지에서 리플렉션을 이용하면 모든 데이터 타입을 구체적으로 지정하지 않아도 된다. 설사 모든 타입에
 * 일일이 대응하도록 코드를 작성할 정도로 인내심이 강하더라도, 코드에 입력될 수 있는 모든 종류의 타입을
 * 예상한다는 것은 현실적으로 불가능하다. 이럴 때 리플렉션 기능을 활용하면, 처음 보는 타입이 들어와도
 * 그 정보를 확인해서 적절히 처리하도록 구현할 수 있다.
 *
 * 결론적으로 최대한 범용성을 제공하도록 구현하거나, 코드를 작성하는 시점에는 없지만 나중에 새로 정의할
 * 데이터 타입을 다루는 코드를 작성하려면 리플렉션을 사용할 수 밖에 없다. 또한, 공통적으로 구현된
 * 인터페이스가 없는, 서로 다른 타입으로 정의된 여러 값들을 다룰 때도 리플렉션을 활용하면 편하다.
 *
 * reflect 패키지에서 주로 사용하는 타입은 reflect.Value와 reflect.Type이다. reflect.Value는
 * 어떤 타입이 가진 값을 저장하는데 사용되고, reflect.Type은 Go 언어에서 지원하는 타입을
 * 표현하는데 사용된다.
 *
 * 리플렉션은 Go 언어에서 제공하는 강력한 기능이라는 점은 의심할 여지가 없다. 하지만 모든 도구가 그렇듯이,
 * 리플렉션도 함부로 남용하면 안되는데 그 이유로 세 가지를 꼽을 수 있다.
 * - 첫 번쨰 이유는 리플렉션을 너무 많이 사용하면 코드를 이해하고 관리하기 힘들어진다. 이러한 문제는 문서를
 *   잘 작성함으로써 어느 정도 해결할 수 있다. 하지만 개발자는 중요한 문서를 작성할 때도 시간을 할애하기
 *   싫어하는 경향이 있다.
 * - 두 번째 이유는 리플렉션을 사용하면 실행 속도가 느려진다. 일반적으로 구체적인 데이터 타입을 다루도록
 *   작성된 코드가 리플렉션으로 데이터 타입을 동적으로 다루는 코드보다 휠씬 빠르다. 또한 이러한 동적 코드로
 *   인해 도구를 이용하여 코드를 리팩토링하거나 분석하기 힘들어진다.
 * - 마지막 세 번째 이유는 리플렉션에 관련된 에러는 빌드 시간에 잡을 수 없고, 프로그램을 실행하다가 뻗으면
 *   그제서야 발견하게 된다. 다시 말해 리플렉션에 관련된 에러로 인해 프로그램 전체가 뻗어버릴 수 있다.
 *   심지어 프로그램을 다 개발하고 몇 개월 또는 몇 년이 지나서야 발견되는 경우도 있다. 이러한 문제에
 *   대처하기 위한 한 가지 방법은 문제의 소지가 있는 함수를 호출하기 전에 충분히 테스타하는 코드를 추가하는
 *   것이다. 하지만 이로 인해 코드가 늘어나서 속도는 더욱 느려진다.
 */
func main() {
	/*
	 * x란 변수를 선언하고 reflect.ValueOf(&x).Elem() 함수를 호출한다. 그런 다음 xRefl.Type()를
	 * 호출해서 그 변수의 타입을 알아낸 다음, 이를 xType에 정의한다. 만약 알고 싶은 것이 변수의 데이터
	 * 타입뿐이라면 그냥 reflect.TypeOf(x)만 호출해도 된다.
	 */
	x := 100
	xRefl := reflect.ValueOf(&x).Elem()
	xType := xRefl.Type()
	fmt.Printf("The type of x is %s.\n", xType)
	fmt.Printf("The type of x is %s.\n", reflect.TypeOf(x))

	/*
	 * 두 가지 타입에 대한 변수 두 개를 정의하는데, 둘 중 한 변수에 대해서만 검사한다. 프로그램을 실행할 때
	 * 커맨드 라인 인수를 하나도 지정하지 않으면 첫 번째 변수만 검사하고, 그렇지 않으면 두 번째 변수를 검사한다.
	 * 이렇게 구성한 이유는 실전에서 사전에 주어진 정보가 없는 struct 변수를 프로그램을 실행하는 동안
	 * 처리하는 방법을 소개한다.
	 */
	A := a{100, 200.12, "Struct a"}
	B := b{1, 2, "Struct b", -1.2}
	var r reflect.Value

	arguments := os.Args
	if len(arguments) == 1 {
		r = reflect.ValueOf(&A).Elem()
	} else {
		r = reflect.ValueOf(&B).Elem()
	}

	iType := r.Type()
	fmt.Printf("i Type: %s\n", iType)
	// NumField() 함수는 reflect.Value 구조체에 있는 필드의 개수를 리턴한다.
	fmt.Printf("The %d fields of %s are:\n", r.NumField(), iType)

	for i := 0; i < r.NumField(); i++ {
		// Field() 함수는 매개변수로 지정한 구조체의 필드를 리턴한다.
		// Interface() 함수는 reflect.Value 구조체에 있는 필드의 값을 인터페이스로 리턴한다.
		fmt.Printf("Field name: %s ", iType.Field(i).Name)
		fmt.Printf("with type: %s ", r.Field(i).Type())
		fmt.Printf("and value %v\n", r.Field(i).Interface())
	}
}
