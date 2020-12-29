package main

import "fmt"

func fibo_1(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo_1(n-1) + fibo_1(n-2)
	}
}

func fibo_2(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibo_2(n-1) + fibo_2(n-2)
}

func fibo_3(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	return fn[n]
}

func main() {
	fmt.Println(fibo_1(40))
	fmt.Println(fibo_2(40))
	fmt.Println(fibo_3(40))
}
