package process

import (
	"fmt"
	"net"
	"ChattingRoom/common/obj"
	"ChattingRoom/common/info"
	_ "ChattingRoom/client/view"
	"encoding/json"
	"errors"
)

func ClientLoginProcess(gconn net.Conn, guser *obj.User)error{

	// send message with conn
	var loginMessage obj.Message
	loginMessage.Type = obj.LoginMesType
	loginMessageStruct := obj.LoginMes{ }
	loginMessageStruct.UserID  = (*guser).UserID
	loginMessageStruct.UserPasswd = (*guser).UserPasswd
	// 序列化->loginMessage.Data
	loginMessageDataSli, marshalErr := json.Marshal(loginMessageStruct)
	if marshalErr != nil{
		fmt.Printf("[%v]:%v\n", info.CurrFuncName(), marshalErr)
		return marshalErr
	}
	loginMessage.Data = string(loginMessageDataSli)
	
	writeMessageErr := (&loginMessage).WriteMessageStructIntoConn(gconn)
	if writeMessageErr != nil {
		fmt.Printf("[%v]:%v\n", info.CurrFuncName(), writeMessageErr)
		return writeMessageErr
	}
	// 读取服务器响应
	var retMessage obj.Message
	readMessageErr := (&retMessage).ReadMessageStructFromConn(gconn)
	if readMessageErr != nil {
		fmt.Printf("[%v]:(%v)\n", info.CurrFuncName(), readMessageErr)
		return readMessageErr
	}
	 fmt.Printf("[debug-%v]:%v\n", info.CurrFuncName(), retMessage)
	var loginResMes obj.LoginResMes
	unmarshalErr := json.Unmarshal([]byte(retMessage.Data), &loginResMes)
	if unmarshalErr != nil {
		fmt.Printf("[%v]:(%v)\n", info.CurrFuncName(), unmarshalErr)
		return unmarshalErr
	}
	if loginResMes.Code != obj.CODE_LOGIN_SUCCESS{
		return errors.New(loginResMes.ErrorText)
	}
	(*guser).UserName = loginResMes.SearchedUserName
	//fmt.Printf("[debug-%v]:(%v)\n", info.CurrFuncName(), loginResMes)
	return nil
}