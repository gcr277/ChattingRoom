package menuSwitch

import (
	"fmt"
	"ChattingRoom/common/info"
	_"ChattingRoom/common/obj"
	"ChattingRoom/client/process"
	"ChattingRoom/client/view"
	_"net"
)

func LoginSuccess(){
	for {
		keyLoginSuccess := view.LoopDispLoginSuccess(GUser.UserID, GUser.UserName)
		switch keyLoginSuccess {
		case view.LS_ONLINELIST:
			reqOLErr := process.RequestOnlineUsers(GConn, GUser.UserID, &GOnlineList)
			if reqOLErr != nil {
				fmt.Printf("[%v]:(%v)\n", info.CurrFuncName(), reqOLErr)
			}else{
				view.DispOnlineList(GOnlineList)
			}
			
		case view.LS_SENDMES:
			reqOLErr := process.RequestOnlineUsers(GConn, GUser.UserID, &GOnlineList)
			if reqOLErr != nil {
				fmt.Printf("[%v]:(%v)\n", info.CurrFuncName(), reqOLErr)
			}else{
				view.DispOnlineList(GOnlineList)
				selectNum, content := view.DispSelAndSendMes()
				sendErr := process.SendMessage(GConn, GUser, GOnlineList[selectNum-1], &content)
				if sendErr != nil {
					fmt.Printf("[%v]:(%v)\n", info.CurrFuncName(), sendErr)
				}
			}
		case view.LS_MESLIST:
	
		case view.LS_EXIT:
			return
		}
	}
}