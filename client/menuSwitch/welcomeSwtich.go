package menuSwitch

import (
	"fmt"
	"ChattingRoom/common/info"
	"ChattingRoom/common/message"
	"ChattingRoom/client/process"
	"ChattingRoom/client/view"
	"net"
)

var (
	User message.User
	Conn net.Conn
)

func WelcomeSwitch(keyWelcome int){
	if keyWelcome == view.W_LOGIN || keyWelcome == view.W_REGISTER{
		Conn = process.DialProcess()
		defer Conn.Close()
	}
	switch keyWelcome{
		case view.W_LOGIN: // login
			User.UserID, User.UserPasswd = view.DispLogin()
			loginErr := process.ClientLoginProcess(Conn, User)
			if loginErr != nil{
				fmt.Printf("[client-%v]:登录失败(%v)\n", info.CurrFuncName(), loginErr)
				
			}else{
				fmt.Printf("[client-%v]:登陆成功!\n", info.CurrFuncName())
				keyLoginSuccess := view.LoopDispLoginSuccess(User.UserID, User.UserName)
				LoginSuccessSwitch(keyLoginSuccess)
				// 再开一个协程来接收服务器推送
				go process.ServerPushProcess(Conn)
			}
		case view.W_REGISTER: // register.go
			
		case view.W_EXIT:
			// exit
		default: //impossible
			fmt.Printf("fatal error!!!!\n")
			return 
	}
	return
}
