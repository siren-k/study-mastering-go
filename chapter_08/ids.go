package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	// 유저 ID는 os.Getuid() 함수만 호출하면 간단히 찾을 수 있다.
	fmt.Println("User id:", os.Getuid())

	/*
	 * 사용자가 속한 그룹 ID를 찾아 화면에 출력한다.
	 */
	var u *user.User
	u, _ = user.Current()
	fmt.Print("Group ids:")
	groupIds, _ := u.GroupIds()
	for _, i := range groupIds {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
