package main

import (
	"ChattingRoom/common/info"
	"ChattingRoom/common/obj"
	"ChattingRoom/server/process"
	"errors"
	"fmt"
	"net"
)
var resFwdChatMesChan chan string = make(chan string, 1)
// 处理客户端连接发来的消息
func handle(conn net.Conn) error {
	defer conn.Close()
	for {
		var recvMessage obj.Message
		readMessageErr := (&recvMessage).ReadMessageStructFromConn(conn)
		if readMessageErr != nil {
			fmt.Printf("[server-%v]:(%v)\n", info.CurrFuncName(), readMessageErr)
			return readMessageErr
		}
		fmt.Printf("[debug-%v]:%v\n", info.CurrFuncName(), recvMessage)


		//channels for concurrent communication
		

		// 根据消息类型，调用不同的处理函数
		switch recvMessage.Type {
		case obj.LoginMesType:
			loginErr := process.LoginProcess(conn, &recvMessage)
			if loginErr != nil {
				return loginErr
			}
		case obj.RegisterMesType:
			registerErr := process.RegisterProcess(conn, &recvMessage)
			if registerErr != nil {
				return registerErr
			}
		case obj.RequestOnlineListMesType:
			resOLErr := process.ResOnlineListProcess(conn, &recvMessage)
			if resOLErr != nil {
				fmt.Printf("[debug-%v]:%v\n", info.CurrFuncName(), resOLErr)
			}
		case obj.ChatMesType:
			forwardOrDumpErr := process.ForwardOrDumpProcess(conn, &recvMessage, resFwdChatMesChan)
			if forwardOrDumpErr != nil {
				fmt.Printf("[debug-%v]:%v\n", info.CurrFuncName(), forwardOrDumpErr)
			}
		case obj.ResFwdChatMesType:
			resFwdChatMesChan <- recvMessage.Data
			
		default:
			unknownTypeErr := errors.New("Unknown message type")
			fmt.Printf("[server-%v]:(%v)", info.CurrFuncName(), unknownTypeErr)
			return unknownTypeErr
		}

	}
	
}
