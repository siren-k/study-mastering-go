package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"os"
	"time"
)

// 첫 번째 핸들러 함수 구현
// 핸들러 함수(Handler Function)는 설정에 따라 한 개 이상의 URL에 대해 서비스 제공
func myHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

// 두 번째 핸들러 함수 구현
// 컨텐츠를 동적으로 생성
func timeHandler1(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

/*
 * 표준 Go 패키지인 net/http/pprof를 이용하면 HTTP 서버 프로그램을 프로파일링을 할 수 있다.
 * 따라서 net/http/pprof를 불러오면 /debug/pprof/ URL 아래에 여러 가지 핸들러가 설치된다.
 * 프로파일러가 http://localhost:8080 주소에 대해 작동하면 다음과 같은 링크도 함께 지원된다.
 * - http://localhost:8080/debug/pprof/goroutine
 * - http://localhost:8080/debug/pprof/heap
 * - http://localhost:8080/debug/pprof/threadcreate
 * - http://localhost:8080/debug/pprof/block
 * - http://localhost:8080/debug/pprof/mutex
 * - http://localhost:8080/debug/pprof/profile
 * - http://localhost:8080/debug/pprof/trace?seconds=5
 */
func main() {
	// 기본 포트 8001
	PORT := ":8001"
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Using default port number: ", PORT)
	} else {
		PORT = ":" + arguments[1]
	}

	// http.NewServeMux를 이용하여 프로그램이 지원하는 경로를 등록한다. 이렇게 하는 주된 이유는
	// http.NewServeMux를 사용하기 위해서는 HTTP 엔드포인트를 직접 정의해야 하기 때문이다.
	// 또한 현재 지원되는 HTTP 엔드포인트의 일부분을 정의할 수도 있다. 만약 http.NewServeMux를
	// 사용하지 않는다면 HTTP 엔드포인트가 자동으로 등록되는데 이 때에서는 net/http/pprof 패키지
	// 앞에 '_' 문자를 붙여서 불러와야 한다.
	r := http.NewServeMux()
	r.HandleFunc("/time", timeHandler1)
	r.HandleFunc("/", myHandler1)

	// 프로파일링에 관련된 HTTP 엔드포인트를 정의
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	// http.ListenAndServe() 함수에 원하는 포트를 지정해서 호출하면 웹 서버가 구동된다.
	err := http.ListenAndServe(PORT, r)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ❯ go tool pprof http://localhost:8001/debug/pprof/profile
	// Fetching profile over HTTP from http://localhost:8001/debug/pprof/profile
	// Saved profile in /Users/benjamin/pprof/pprof.samples.cpu.001.pb.gz
	// Type: cpu
	// Time: Dec 30, 2020 at 8:13am (KST)
	// Duration: 30s, Total samples = 0
	// No samples were found with the default sample value type.
	// Try "sample_index" command to analyze different sample values.
	// Entering interactive mode (type "help" for commands, "o" for options)
	// (pprof) top
	// Showing nodes accounting for 0, 0% of 0 total
	//       flat  flat%   sum%        cum   cum%
	// (pprof)
	// (pprof)
	// (pprof)
	// (pprof) help
	//   Commands:
	//     callgrind        Outputs a graph in callgrind format
	//     comments         Output all profile comments
	//     disasm           Output assembly listings annotated with samples
	//     dot              Outputs a graph in DOT format
	//     eog              Visualize graph through eog
	//     evince           Visualize graph through evince
	//     gif              Outputs a graph image in GIF format
	//     gv               Visualize graph through gv
	//     kcachegrind      Visualize report in KCachegrind
	//     list             Output annotated source for functions matching regexp
	//     pdf              Outputs a graph in PDF format
	//     peek             Output callers/callees of functions matching regexp
	//     png              Outputs a graph image in PNG format
	//     proto            Outputs the profile in compressed protobuf format
	//     ps               Outputs a graph in PS format
	//     raw              Outputs a text representation of the raw profile
	//     svg              Outputs a graph in SVG format
	//     tags             Outputs all tags in the profile
	//     text             Outputs top entries in text form
	//     top              Outputs top entries in text form
	//     topproto         Outputs top entries in compressed protobuf format
	//     traces           Outputs all profile samples in text form
	//     tree             Outputs a text rendering of call graph
	//     web              Visualize graph through web browser
	//     weblist          Display annotated source in a web browser
	//     o/options        List options and their current values
	//     quit/exit/^D     Exit pprof
	//
	//   Options:
	//     call_tree        Create a context-sensitive call tree
	//     compact_labels   Show minimal headers
	//     divide_by        Ratio to divide all samples before visualization
	//     drop_negative    Ignore negative differences
	//     edgefraction     Hide edges below <f>*total
	//     focus            Restricts to samples going through a node matching regexp
	//     hide             Skips nodes matching regexp
	//     ignore           Skips paths going through any nodes matching regexp
	//     mean             Average sample value over first value (count)
	//     nodecount        Max number of nodes to show
	//     nodefraction     Hide nodes below <f>*total
	//     noinlines        Ignore inlines.
	//     normalize        Scales profile based on the base profile.
	//     output           Output filename for file-based outputs
	//     prune_from       Drops any functions below the matched frame.
	//     relative_percentages Show percentages relative to focused subgraph
	//     sample_index     Sample value to report (0-based index or name)
	//     show             Only show nodes matching regexp
	//     show_from        Drops functions above the highest matched frame.
	//     source_path      Search path for source files
	//     tagfocus         Restricts to samples with tags in range or matched by regexp
	//     taghide          Skip tags matching this regexp
	//     tagignore        Discard samples with tags in range or matched by regexp
	//     tagshow          Only consider tags matching this regexp
	//     trim             Honor nodefraction/edgefraction/nodecount defaults
	//     trim_path        Path to trim from source paths before search
	//     unit             Measurement units to display
	//
	//   Option groups (only set one per group):
	//     cumulative
	//       cum              Sort entries based on cumulative weight
	//       flat             Sort entries based on own weight
	//     granularity
	//       addresses        Aggregate at the address level.
	//       filefunctions    Aggregate at the function level.
	//       files            Aggregate at the file level.
	//       functions        Aggregate at the function level.
	//       lines            Aggregate at the source code line level.
	//   :   Clear focus/ignore/hide/tagfocus/tagignore
	//
	//   type "help <cmd|option>" for more information
	// (pprof) list myHandler1
	// Total: 0
	// (pprof)

	// http://localhost:8001/debug/pprof/

	// ==> 웹 서버 애플리케이션의 성능도 테스트하고 싶다면 ab(1) 유틸리티를 사용한다. 이 유틸리티는
	//     '아파치 HTTP 서버 벤치마킹 툴'이란 이름으로 더 많이 알려졌다. 이를 통해 트래픽을 생성해서
	//     wwwProfile.go를 벤치마킹할 수 있다. 또한 다음과 같이 go tool pprof를 이용하여
	//     보다 정확한 데이터를 구할 수 있다.
	// ❯ ab -k -c 10 -n 100000 "http://127.0.0.1:8001/time"
	// This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
	// Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
	// Licensed to The Apache Software Foundation, http://www.apache.org/
	//
	// Benchmarking 127.0.0.1 (be patient)
	// Completed 10000 requests
	// Completed 20000 requests
	// Completed 30000 requests
	// Completed 40000 requests
	// Completed 50000 requests
	// Completed 60000 requests
	// Completed 70000 requests
	// Completed 80000 requests
	// Completed 90000 requests
	// Completed 100000 requests
	// Finished 100000 requests
	//
	//
	// Server Software:
	// Server Hostname:        127.0.0.1
	// Server Port:            8001
	//
	// Document Path:          /time
	// Document Length:        111 bytes
	//
	// Concurrency Level:      10
	// Time taken for tests:   3.853 seconds
	// Complete requests:      100000
	// Failed requests:        0
	// Keep-Alive requests:    100000
	// Total transferred:      25200000 bytes
	// HTML transferred:       11100000 bytes
	// Requests per second:    25953.39 [#/sec] (mean)
	// Time per request:       0.385 [ms] (mean)
	// Time per request:       0.039 [ms] (mean, across all concurrent requests)
	// Transfer rate:          6386.97 [Kbytes/sec] received
	//
	// Connection Times (ms)
	//               min  mean[+/-sd] median   max
	// Connect:        0    0   0.0      0       1
	// Processing:     0    0   2.5      0     248
	// Waiting:        0    0   2.5      0     248
	// Total:          0    0   2.5      0     248
	//
	// Percentage of the requests served within a certain time (ms)
	//   50%      0
	//   66%      0
	//   75%      0
	//   80%      0
	//   90%      1
	//   95%      1
	//   98%      3
	//   99%      4
	//  100%    248 (longest request)
}
