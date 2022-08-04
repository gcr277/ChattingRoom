package view

import (
	"fmt"
)

func DispLogin() (userID int, userPasswd string) {
	fmt.Printf("----------------- login -----------------\n")
	fmt.Printf("input user ID:")
	fmt.Scanf("%d", &userID)
	fmt.Printf("input user password:")
	fmt.Scanf("%s", &userPasswd)
	return
}
