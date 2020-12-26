package main

import (
	"fmt"
	"os"
	"reflect"
)

// t1과 t2는 모두 int 기반의 타입이다. 따라서 실질적으로 int와 같지만 Go 언어에서는 전혀 다른 타입인 것처럼 취급한다.
// 이 프로그램을 파싱한 뒤에 나오는 내부 표현식에서는 각각 main.t1과 main.t2라고 부른다.
type t1 int
type t2 int

type A struct {
	X    int
	Y    float64
	Text string
}

/*
 * A라는 구조체 타입을 정의하고 compareStruct()란 함수를 구현했다. 이 함수의 목적은 A 타입을 가진 변수 두 개가
 * 서로 같은지 확인하는 것이다. 코드에서 볼 수 있듯이, compareStruct()는 이 작업을 reflection.go에 나온
 * 코드를 이용하여 처리하고 있다.
 */
func (A1 A) compareStruct(A2 A) bool {
	r1 := reflect.ValueOf(&A1).Elem()
	r2 := reflect.ValueOf(&A2).Elem()

	for i := 0; i < r1.NumField(); i++ {
		if r1.Field(i).Interface() != r2.Field(i).Interface() {
			return false
		}
	}
	return true
}

func printMethods(i interface{}) {
	r := reflect.ValueOf(i)
	t := r.Type()
	fmt.Println("Type to examine: %s\n", t)

	for j := 0; j < r.NumMethod(); j++ {
		m := r.Method(j).Type()
		fmt.Println(t.Method(j).Name, "-->", m)
	}
}

func main() {
	x1 := t1(100)
	x2 := t2(100)
	fmt.Printf("The type of x1 is %s\n", reflect.TypeOf(x1))
	fmt.Printf("The type of x2 is %s\n", reflect.TypeOf(x2))

	var p struct{}
	r := reflect.New(reflect.ValueOf(&p).Type()).Elem()
	fmt.Printf("The type of r is %s\n", reflect.TypeOf(r))

	a1 := A{1, 2.1, "A1"}
	a2 := A{1, -2, "A2"}

	// a1.compareStruct(a1)을 호출하면 true가 리턴된다. 자기 자신과 비교했기 때문에 당연하다.
	if a1.compareStruct(a1) {
		fmt.Println("Equal!")
	}

	// 하지만, a1.compareStruct(a2)를 호출하면 false를 리턴하는데, a1과 a2의 값이 서로 다르기 때문이다.
	if !a1.compareStruct(a2) {
		fmt.Println("Not Equal!")
	}

	var f *os.File
	printMethods(f)

	// ❯ go run advRefl.go
	// The type of x1 is main.t1
	// The type of x2 is main.t2
	// The type of r is reflect.Value
	// Equal!
	// Not Equal!
	// Type to examine: %s
	//  *os.File
	// Chdir --> func() error
	// Chmod --> func(os.FileMode) error
	// Chown --> func(int, int) error
	// Close --> func() error
	// Fd --> func() uintptr
	// Name --> func() string
	// Read --> func([]uint8) (int, error)
	// ReadAt --> func([]uint8, int64) (int, error)
	// ReadFrom --> func(io.Reader) (int64, error)
	// Readdir --> func(int) ([]os.FileInfo, error)
	// Readdirnames --> func(int) ([]string, error)
	// Seek --> func(int64, int) (int64, error)
	// SetDeadline --> func(time.Time) error
	// SetReadDeadline --> func(time.Time) error
	// SetWriteDeadline --> func(time.Time) error
	// Stat --> func() (os.FileInfo, error)
	// Sync --> func() error
	// SyscallConn --> func() (syscall.RawConn, error)
	// Truncate --> func(int64) error
	// Write --> func([]uint8) (int, error)
	// WriteAt --> func([]uint8, int64) (int, error)
	// WriteString --> func(string) (int, error)
}
