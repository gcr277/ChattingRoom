package view

import (
	"ChattingRoom/common/info"
	"fmt"
)
const (
	W_LOGIN = 1
	W_REGISTER = 2
	W_EXIT = 3
)
func LoopDispWelcome() int {
	var key int
loopWelcome:
	for {
		fmt.Printf("----------------- welcome -----------------\n")
		fmt.Printf("\t1 login\n" +
			"\t2 register\n" +
			"\t3 exit\n" +
			"input number: ")
		fmt.Scanf("%d\n", &key)

		switch key {
		case W_LOGIN, W_REGISTER, W_EXIT:
			break loopWelcome
		default:
			fmt.Printf("[client-%v]:illegal input!\n", info.CurrFuncName())
		}
	}
	return key
}
