package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"net/http"
	"os"
)

/*
 * SQLite3 테이블에서 읽은 레코드를 표현한다.
 */
type HtmlTemplateEntry struct {
	Number int
	Double int
	Square int
}

/*
 * DATA와 tFile이란 전역 변수를 선언하고 있다. 이 변수는 각각 템플릿 파일에
 * 전달할 데이터와 템플릿 파일의 이름을 표현한다.
 */
var DATA []HtmlTemplateEntry
var tFile string

/*
 * myHandler() 함수의 코드는 놀라울 정도로 간결하고 효과적으로 구성돼 있다. 특히 함수의 길이가
 * 엄청나게 짧다. 주요 작업은 template.ExecuteTemplate() 함수에서 처리한다. 이 함수의
 * 첫 번째 매개변수는 HTTP 클라이언트에 대한 연결 정보를 담는 변수를 지정하고, 두 번째 매개변수는
 * 데이터를 표현하는데 적용할 템플릿 파일을 지정하고, 세 번째 매개변수는 처리할 데이터를 표현한
 * 구조체 슬라이스다.
 */
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Host: %s\n", r.Host, r.URL.Path)
	myT := template.Must(template.ParseGlob(tFile))
	myT.ExecuteTemplate(w, tFile, DATA)
}

func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Println("Need Database File + Template File!")
		return
	}

	database := arguments[1]
	tFile = arguments[2]

	// sql.Open() 함수는 지정한 데이터베이스를 연결한다.
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		fmt.Println(nil)
		return
	}

	// 데이터베이스에서 실행할 커맨드는 db.Exec()로 처리한다.
	// 커맨드를 실행한 결과는 받지 않는다.
	fmt.Println("Emptying database table.")
	_, err = db.Exec("DELETE FROM data")
	if err != nil {
		fmt.Println(nil)
		return
	}

	// 데이터베이스 커맨드에서 몇 가지 매개변수만 바꿔서 여러 번 실행하려면 db.Prepare() 함수를
	// 먼저 호출할 뒤에 Exec()를 호출한다.
	fmt.Println("Populating", database)
	stmt, _ := db.Prepare("INSERT INTO data(number, double, square) values (?, ?, ?)")
	for i := 20; i < 50; i++ {
		_, _ = stmt.Exec(i, 2*i, i*i)
	}

	/*
	 * db.Query()를 호출하고 이어서 Next()와 Scan()을 여러 차례 호출하는 방식으로
	 * 지정한 테이블로부터 데이터를 읽고 있다. 데이터를 읽는 동시에 이를 구조체 슬라이스에 저장한다.
	 */
	rows, err := db.Query("SELECT * FROM data")
	if err != nil {
		fmt.Println(nil)
		return
	}

	var n int
	var d int
	var s int
	for rows.Next() {
		err = rows.Scan(&n, &d, &s)
		temp := HtmlTemplateEntry{Number: n, Double: d, Square: s}
		DATA = append(DATA, temp)
	}

	/*
	 * 웹 서버를 설정한다. 여기서 http.HandleFunc() 함수는 이 프로그램 안에 웹 서버를 하나 만들고 있다. 여기서
	 * '/'라는 URL에 대한 핸들러(myHandler())를 지정했다. '/'는 모든 URL에 해당한다. 이렇게 하면 별도의
	 * 정적 또는 동적 페이지를 생성하지 않아도 된다.
	 */
	http.HandleFunc("/", myHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ❯ sqlite3 htmlT.db
	// SQLite version 3.32.3 2020-06-18 14:16:19
	// Enter ".help" for usage hints.
	//     sqlite> create table data (
	// ...> number integer primary key,
	// ...> double integer,
	// ...> square integer);
	// sqlite> ^D
	// ❯ ls -l htmlT.db
	// -rw-r--r--  1 benjamin  staff  8192 Dec 26 10:34 htmlT.db
	//
	// ❯ go run htmlT.go htmlT.db html.gohtml.html
	// Emptying database table.
	// Populating htmlT.db
	// Host: localhost:8080
	// %!(EXTRA string=/)Host: localhost:8080
	// %!(EXTRA string=/favicon.ico)
}
