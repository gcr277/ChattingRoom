package process

import (
	"ChattingRoom/common/info"
	"ChattingRoom/common/message"
	"ChattingRoom/server/service"
	"encoding/json"
	"fmt"
	"net"
)

func LoginProcess(conn net.Conn, recvMessagePtr *message.Message) error {
	var loginMes message.LoginMes
	unmarshalErr := json.Unmarshal([]byte((*recvMessagePtr).Data), &loginMes)
	if unmarshalErr != nil {
		return unmarshalErr
	}
	fmt.Printf("[debug-%v]:(%v)\n", info.CurrFuncName(), loginMes)

	// 校验登陆信息...
	checkLoginErr := service.CheckLoginMessage(&loginMes)
	var loginResMes message.LoginResMes
	if checkLoginErr != nil { // 如果登录信息错误	
		loginResMes.Code = 500
		loginResMes.ErrorText = checkLoginErr.Error()
	} else {
		// 服务器端设置登录状态
		// ...service.SetLoginFlag(&loginMes)
		loginResMes.Code = 200
		loginResMes.ErrorText = ""
	}
	// 对loginResMes序列化, 构造resMessage
	var resMessage message.Message
	resMessage.Type = message.LoginResMesType
	resMessageDataSli, marshalErr := json.Marshal(loginResMes)
	if marshalErr != nil {
		return marshalErr
	}
	resMessage.Data = string(resMessageDataSli)
	// 对resMessage序列化并发送
	writeMessageErr := (&resMessage).WriteMessageStructIntoConn(conn)
	if writeMessageErr != nil {
		return writeMessageErr
	}

	return nil
}
