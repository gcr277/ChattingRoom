package main

import (
	"ChattingRoom/common/info"
	"ChattingRoom/common/message"
	"ChattingRoom/server/process"
	"errors"
	"fmt"
	"net"
)

// 处理客户端连接发来的消息
func handle(conn net.Conn) error {
	defer conn.Close()
	for {
		var recvMessage message.Message
		readMessageErr := (&recvMessage).ReadMessageStructFromConn(conn)
		if readMessageErr != nil {
			fmt.Printf("[server-%v]:(%v)\n", info.CurrFuncName(), readMessageErr)
			return readMessageErr
		}
		fmt.Printf("[debug-%v]:%v\n", info.CurrFuncName(), recvMessage)
		// 根据消息类型，调用不同的处理函数
		switch recvMessage.Type {
		case message.LoginMesType:
			loginErr := process.LoginProcess(conn, &recvMessage)
			if loginErr != nil {
				return loginErr
			}
			//case message.RegisterMesType:

		default:
			unknownTypeErr := errors.New("Unknown message type")
			fmt.Printf("[server-%v]:(%v)", info.CurrFuncName(), unknownTypeErr)
			return unknownTypeErr
		}

	}
}
