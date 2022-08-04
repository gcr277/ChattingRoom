package process

import (
	"net"
	"fmt"
	"ChattingRoom/common/message"
	"ChattingRoom/common/info"
)

func ServerPushProcess(conn net.Conn){
	var serverPushMes message.Message
	for {
		readMessageErr := (&serverPushMes).ReadMessageStructFromConn(conn)
		if readMessageErr != nil {
			fmt.Printf("[%v]:(%v)\n", info.CurrFuncName(), readMessageErr)
			return
		}
		fmt.Printf("[debug-%v]:(%v)\n", info.CurrFuncName(), serverPushMes)
		//...
		
	}
}