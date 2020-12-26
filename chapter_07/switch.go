package main

import "fmt"

type Square struct {
	X float64
}

type Circle struct {
	R float64
}

type Rectangle struct {
	X float64
	Y float64
}

/*
 * 매개변수 x에 지정된 타입을 구분하는 방법을 보여준다. 여기서 핵심은 x에 대한
 * 타입을 리턴하는 x.(type)이란 구문에 있다. fmt.Printf() 함수에서 %v를
 * 이러한 타입 값을 가져올 수 있다.
 */
func tellInterface(x interface{}) {
	switch v := x.(type) {
	case Square:
		fmt.Println("This is a Square!")
	case Circle:
		fmt.Printf("%v is a Circle!\n", v)
	case Rectangle:
		fmt.Println("This is a Rectangle!")
	default:
		fmt.Printf("Unknown type %T!\n", v)
	}
}

func main() {
	x := Circle{R: 10}
	tellInterface(x)

	y := Rectangle{X: 4, Y: 1}
	tellInterface(y)

	z := Square{X: 4}
	tellInterface(z)

	tellInterface(10)

	// ❯ go run switch.go
	// {10} is a Circle!
	// This is a Rectangle!
	// This is a Square!
	// Unknown type int!
}
