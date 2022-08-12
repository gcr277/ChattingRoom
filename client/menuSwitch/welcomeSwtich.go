package menuSwitch

import (
	"fmt"
	"ChattingRoom/common/info"
	"ChattingRoom/common/obj"
	"ChattingRoom/client/process"
	"ChattingRoom/client/view"
	"net"
)

var (
	GUser obj.User
	GConn net.Conn
	GOnlineList obj.ResOnlineListMes
)

func WelcomeSwitch(keyWelcome int){
	if keyWelcome == view.W_LOGIN || keyWelcome == view.W_REGISTER{
		var dialErr error
		GConn, dialErr = process.DialProcess()
		if dialErr != nil{
			fmt.Printf("[client-%v]:登录失败(%v)\n", info.CurrFuncName(), dialErr)
			return 
		}
		defer GConn.Close()
	}
	switch keyWelcome{
		case view.W_LOGIN: // login
			GUser.UserID, GUser.UserPasswd = view.DispLogin()
			loginErr := process.ClientLoginProcess(GConn, &GUser)
			if loginErr != nil{
				fmt.Printf("[client-%v]:登录失败(%v)\n", info.CurrFuncName(), loginErr)
			}else{
				fmt.Printf("[client-%v]:登录成功!\n", info.CurrFuncName())
				// 
				// 再开一个协程来接收服务器推送
				go process.ServerPushProcess(GConn, GUser)
				
				LoginSuccess()
				
			}
		case view.W_REGISTER: // register.go
			GUser.UserID, GUser.UserName, GUser.UserPasswd = view.DispRegister()
			registerErr := process.RegisterProcess(GConn, &GUser)
			if registerErr != nil{
				fmt.Printf("[client-%v]:注册失败(%v)\n", info.CurrFuncName(), registerErr)
			}else{
				fmt.Printf("[client-%v]:注册成功!\n", info.CurrFuncName())
				// ...
			}
		case view.W_EXIT:
			// exit
		default: //impossible
			fmt.Printf("fatal error!!!!\n")
			return 
	}
	return
}
