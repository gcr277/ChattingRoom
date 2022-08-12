package view

import (
	"fmt"
)

func DispRegister() (userID int, userName string, userPasswd string) {
	fmt.Printf("----------------- register -----------------\n")
	fmt.Printf("input user ID:")
	fmt.Scanf("%d", &userID)
	fmt.Printf("input user name:")
	fmt.Scanf("%s", &userName)
	fmt.Printf("input user password:")
	fmt.Scanf("%s", &userPasswd)
	return
}
