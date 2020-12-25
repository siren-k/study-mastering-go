package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("usage: switch number")
		os.Exit(1)
	}

	number, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("This value is not an integer:", number)
	} else {
		switch {
		case number < 0:
			fmt.Println("Less than zero!")
		case number > 0:
			fmt.Println("Bigger than zero!")
		default:
			fmt.Println("Zero!")
		}
	}

	asString := arguments[1]
	switch asString {
	case "5":
		fmt.Println("Five!")
	case "0":
		fmt.Println("Zero!")
	default:
		fmt.Println("Do not care!")
	}

	var negative = regexp.MustCompile(`-`)
	var floatingPoint = regexp.MustCompile(`\d?.\d`)
	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)
	switch {
	case negative.MatchString(asString):
		fmt.Println("Negative Number")
	case floatingPoint.MatchString(asString):
		fmt.Println("Floating Point!")
	case email.MatchString(asString):
		fmt.Println("It is an email")
		fallthrough // 현재 case 뒤에 나오는 case 문장도 실행
	default:
		fmt.Println("Something else!")
	}

	var aType error = nil
	switch aType.(type) {
	case nil:
		fmt.Println("It is nil interface!")
	default:
		fmt.Println("Not nil interface!")
	}
}