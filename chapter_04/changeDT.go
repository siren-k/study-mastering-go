package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

//logs := []string{
//"127.0.0.1 - - [16/Nov/2017:10:49:46 +0200] 325504",
//"127.0.0.1 - - [16/Nov/2018:10:16:41 +0200] \"GET /CVEN HTTP/1.1\" 200 12531 \"=\" \"Mozilla/5.0 AppleWebKit/537.36",
//"127.0.0.1 200 9412 - - [12/Nov/2017:06:26:05 +0200] \"GET \"http://www.mtsoukalous.eu/taxonomy/term/47\" 1507",
//"[12/Nov/2017:16:27:21 +0300]",
//"[12/Nov/2017:20:88:21 +0200]",
//"[12/Nov/2017:20:21 +0200]",
//}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Print("Please provide one text file to process!")
		os.Exit(1)
	}

	filename := arguments[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file %s", err)
		os.Exit(1)
	}
	defer f.Close()

	notAMatch := 0
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("error reading file %s", err)
		}

		r1 := regexp.MustCompile(`.*\[(\d\d/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\].*`)
		if r1.MatchString(line) {
			match := r1.FindStringSubmatch(line)
			d1, err := time.Parse("02/Jan/2006:15:04:05 -0700", match[1])
			if err == nil {
				newFormat := d1.Format(time.Stamp)
				fmt.Print("1>", strings.Replace(line, match[1], newFormat, 1))
			} else {
				notAMatch++
			}
			continue
		}

		r2 := regexp.MustCompile(`.*\[(\w+\-\d\d-\d\d:\d\d:\d\d:\d\d.*)\].*`)
		if r2.MatchString(line) {
			match := r2.FindStringSubmatch(line)
			d1, err := time.Parse("Jan-02-06:15:04:05 -0700", match[1])
			if err == nil {
				newFormat := d1.Format(time.Stamp)
				fmt.Print("2>", strings.Replace(line, match[1], newFormat, 1))
			} else {
				notAMatch++
			}
			continue
		}
	}
	fmt.Println(notAMatch, "lines did not match!")
}
