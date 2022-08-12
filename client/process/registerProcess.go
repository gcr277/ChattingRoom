package process

import (
	"net"
	"ChattingRoom/common/obj"
	"ChattingRoom/common/info"
	"encoding/json"
	"fmt"
	"errors"
)

func RegisterProcess(gconn net.Conn, guser *obj.User) error{
	// send message with conn
	var registerMessage obj.Message
	registerMessage.Type = obj.RegisterMesType
	registerMessageStruct := obj.RegisterMes{ }
	registerMessageStruct.UserID  = (*guser).UserID
	registerMessageStruct.UserPasswd = (*guser).UserPasswd
	registerMessageStruct.UserName = (*guser).UserName
	// 序列化->RegisterMessage.Data
	registerMessageDataSli, marshalErr := json.Marshal(registerMessageStruct)
	if marshalErr != nil{
		fmt.Printf("[%v]:%v\n", info.CurrFuncName(), marshalErr)
		return marshalErr
	}
	registerMessage.Data = string(registerMessageDataSli)
	
	writeMessageErr := (&registerMessage).WriteMessageStructIntoConn(gconn)
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
	var registerResMes obj.RegisterResMes
	unmarshalErr := json.Unmarshal([]byte(retMessage.Data), &registerResMes)
	if unmarshalErr != nil {
		fmt.Printf("[%v]:(%v)\n", info.CurrFuncName(), unmarshalErr)
		return unmarshalErr
	}
	if registerResMes.Code != obj.CODE_REGISTER_SUCCESS{
		return errors.New(registerResMes.ErrorText)
	}
	
	return nil
}