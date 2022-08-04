package process

import (
	"fmt"
	"net"
	"ChattingRoom/common/message"
	"ChattingRoom/common/info"
	_ "ChattingRoom/client/view"
	"encoding/json"
	"errors"
)

func ClientLoginProcess(conn net.Conn, user message.User)error{

	// send message with conn
	var loginMessage message.Message
	loginMessage.Type = message.LoginMesType
	loginMessageStruct := message.LoginMes{
		UserID : user.UserID,
		UserPasswd : user.UserPasswd,
	}
	// 序列化->loginMessage.Data
	loginMessageDataSli, marshalErr := json.Marshal(loginMessageStruct)
	if marshalErr != nil{
		fmt.Printf("[%v]:%v\n", info.CurrFuncName(), marshalErr)
		return marshalErr
	}
	loginMessage.Data = string(loginMessageDataSli)

	writeMessageErr := (&loginMessage).WriteMessageStructIntoConn(conn)
	if writeMessageErr != nil {
		fmt.Printf("[%v]:%v\n", info.CurrFuncName(), writeMessageErr)
		return writeMessageErr
	}
	// 读取服务器响应
	var retMessage message.Message
	readMessageErr := (&retMessage).ReadMessageStructFromConn(conn)
	if readMessageErr != nil {
		fmt.Printf("[%v]:(%v)\n", info.CurrFuncName(), readMessageErr)
		return readMessageErr
	}
	// fmt.Printf("[debug-%v]:%v\n", info.CurrFuncName(), retMessage)
	var loginResMes message.LoginResMes
	unmarshalErr := json.Unmarshal([]byte(retMessage.Data), &loginResMes)
	if unmarshalErr != nil {
		fmt.Printf("[%v]:(%v)\n", info.CurrFuncName(), unmarshalErr)
		return unmarshalErr
	}
	if loginResMes.Code != 200{
		return errors.New(loginResMes.ErrorText)
	}

	//fmt.Printf("[debug-%v]:(%v)\n", info.CurrFuncName(), loginResMes)
	return nil
}