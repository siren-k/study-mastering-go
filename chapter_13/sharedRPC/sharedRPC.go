package sharedRPC

type MyFloats struct {
	A1, A2 float64
}

// sharedRPC 패키지는 MyInterface 인터페이스와 MyFloats 구조체를 정의한다.
// 둘 다 클라이언트와 서버에서 사용하지만 인터페이스를 구현하는 일은 RPC 서버에서만 한다.
type MyInterface interface {
	Multiply(arguments *MyFloats, reply *float64) error
	Power(arguments *MyFloats, reply *float64) error
}

// ❯ cp -R sharedRPC ~/.go/src
// ❯ go install sharedRPC
