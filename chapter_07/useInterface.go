package main

import (
	"fmt"
	"math"
	"myInterface" // 인터페이스를 정의하는 코드를 패키지로 만들었기 때문에 import문에 myInterface 패키지를 불러왔음
)

type square struct {
	X float64
}

type circle struct {
	R float64
}

/*
 * square 타입에 사용할 Shape 인터페이스를 구현하고 있다.
 */
func (s square) Area() float64 {
	return s.X * s.X
}

func (s square) Perimeter() float64 {
	return 4 * s.X
}

/*
 * circle 타입에 사용할 Shape 인터페이스를 구현하고 있다.
 */
func (c circle) Area() float64 {
	return c.R * c.R * math.Pi
}

func (c circle) Perimeter() float64 {
	return 2 * c.R * math.Pi
}

/*
 * 하나의 매개변수(myInterface.Shape)를 받는 함수를 구현하고 있다.
 * 이 함수는 Shape 인터페이스를 구현한 값이라면 어떠한 값도 매개변수로 받을 수 있다.
 */
func Calculate(x myInterface.Shape) {
	_, ok := x.(circle)
	if ok {
		fmt.Println("Is a circle!")
	}

	v, ok := x.(square)
	if ok {
		fmt.Println("Is a square:", v)
	}

	fmt.Println(x.Area())
	fmt.Println(x.Perimeter())
}

func main() {
	x := square{X: 10}
	fmt.Println("Perimeter:", x.Perimeter())
	Calculate(x)
	y := circle{R: 5}
	Calculate(y)

	// ❯ go run useInterface.go
	// Perimeter: 40
	// Is a square: {10}
	// 100
	// 40
	// Is a circle!
	// 78.53981633974483
	// 31.41592653589793
}
