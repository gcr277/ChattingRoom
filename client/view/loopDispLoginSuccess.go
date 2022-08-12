package view

import (
	"ChattingRoom/common/info"
	"fmt"
)

const (
	LS_ONLINELIST = 1
	LS_SENDMES = 2
	LS_MESLIST = 3
	LS_ACCOUNTMAN = 4
	LS_EXIT = 5
)
func LoopDispLoginSuccess(userID int, userName string) int {
	var key int
loopLoginSuccess:
	for {
		fmt.Printf("--------------------- menu ----------------------\n"+
				   "|\tID:[%v]\tName:[%v]\n"+
				   "-------------------------------------------------\n",
			userID, userName)
		fmt.Printf("\t%v 显示在线用户列表\n" +
				   "\t%v 发送消息\n" +
				   "\t%v 离线信息列表\n" +
				   "\t%v 账号管理\n" +
				   "\t%v 退出系统\n" +
				   "input number:\n",
				   LS_ONLINELIST,LS_SENDMES,LS_MESLIST,LS_ACCOUNTMAN,LS_EXIT)
		fmt.Scanf("%d\n", &key)
		switch key {
		case LS_ONLINELIST, LS_SENDMES, LS_MESLIST, LS_ACCOUNTMAN, LS_EXIT:
			break loopLoginSuccess
		
		default:
			fmt.Printf("[client-%v]:illegal input!.\n", info.CurrFuncName())
		}
	}
	return key

}
