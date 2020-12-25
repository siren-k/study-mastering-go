package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func randomString(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	MIN := 0
	MAX := 94
	SEED := time.Now().Unix()
	var LENGTH int64 = 8

	arguments := os.Args
	switch len(arguments) {
	case 2:
		LENGTH, _ = strconv.ParseInt(os.Args[1], 10, 64)
	default:
		fmt.Println("Using default values!")
	}

	rand.Seed(SEED)

	startChar := "!"
	var i int64 = 1
	for {
		myRand := randomString(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		fmt.Print(newChar)
		if i == LENGTH {
			break
		}
		i++
	}
	fmt.Println()

	// ❯ go run generatePassword.go
	// Using default values!
	// +q!Y22?!
	// ❯ go run generatePassword.go
	// Using default values!
	// 2%1C'Zu{
	// ❯ go run generatePassword.go 20
	// Z$D~TAT8Fx2'6vC0vR1I
	// ❯ go run generatePassword.go 20
	// 5{a'n:W4sN+)eKYT2JTa
}
