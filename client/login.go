package main

import (
	"fmt"
	"net"
	"go_practice/projChattingRoom/common/message"
	"go_practice/projChattingRoom/common/info"
	"encoding/json"
)

func login(userID int, userPasswd string)error{
	
	//fmt.Printf("userID = %v userPasswd = %v\n", userID, userPasswd)
	// connect to server
	conn, dialErr := net.Dial("tcp","localhost:8888")
	if dialErr != nil{
		fmt.Printf("[client-%v]:%v\n", info.CurrFuncName(), dialErr)
		return dialErr
	}
	defer conn.Close()

	// 下面可以封装为sendMes(conn, type, dataStruct)error
	// send message with conn
	var loginMessage message.Message
	loginMessage.Type = message.LoginMesType
	loginMessageStruct := message.LoginMes{
		UserID : userID,
		UserPasswd : userPasswd,
	}
	// 序列化->loginMessage.Data
	loginMessageDataSli, marshalErr := json.Marshal(loginMessageStruct)
	if marshalErr != nil{
		fmt.Printf("[client-%v]:%v\n", info.CurrFuncName(), marshalErr)
		return marshalErr
	}
	loginMessage.Data = string(loginMessageDataSli)

	// 序列化
	loginMessageSli, marshalErr := json.Marshal(loginMessage)
	if marshalErr != nil{
		fmt.Printf("[client-%v]:%v\n", info.CurrFuncName(), marshalErr)
		return marshalErr
	}
	fmt.Printf("[client debug]:%v\n", string(loginMessageSli))
	_, writeErr := conn.Write(loginMessageSli)
	if writeErr != nil{
		fmt.Printf("[client-%v]:%v\n", info.CurrFuncName(), writeErr)
		return writeErr
	}
	return nil
}