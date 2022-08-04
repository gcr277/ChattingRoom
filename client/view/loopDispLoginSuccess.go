package view

import (
	"ChattingRoom/common/info"
	"fmt"
)

const (
	LS_ONLINELIST = 1
	LS_SENDMES = 2
	LS_MESLIST = 3
	LS_EXIT = 4
)
func LoopDispLoginSuccess(userID int, userName string) int {
	var key int
loopLoginSuccess:
	for {
		fmt.Printf("----------------- login success -----------------\n"+
				   "|\tID:[%v]\tName:[%v]\n"+
				   "-------------------------------------------------\n",
			userID, userName)
		fmt.Printf("\t1 显示在线用户列表\n" +
				   "\t2 发送消息\n" +
				   "\t3 信息列表\n" +
				   "\t4 退出系统\n" +
				   "input number: ")
		fmt.Scanf("%d\n", &key)
		switch key {
		case LS_ONLINELIST, LS_SENDMES, LS_MESLIST, LS_EXIT:
			break loopLoginSuccess
		
		default:
			fmt.Printf("[client-%v]:illegal input, input again: ", info.CurrFuncName())
		}
	}
	return key

}
