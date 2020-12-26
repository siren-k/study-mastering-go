package myInterface

/*
 * Shape 인터페이스는 굉장히 간단하고 직관적이다. Area()와 Perimeter()란 두 함수만
 * 구현하면 되고, 둘 다 float64 값을 리턴한다. 첫 번째 함수는 평명에 대한 도형의 면적을
 * 계산하고, 두 번째 함수는 평면에 있는 도현의 둘레를 계산한다.
 */
type Shape interface {
	Area() float64
	Perimeter() float64
}
