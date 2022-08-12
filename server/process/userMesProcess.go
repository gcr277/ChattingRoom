package process

import (
	"ChattingRoom/common/info"
	"ChattingRoom/common/obj"
	"ChattingRoom/server/service"
	"encoding/json"
	"fmt"
	"net"
)
/*************************************************/
func LoginProcess(conn net.Conn, recvMessagePtr *obj.Message) error {
	var loginMes obj.LoginMes
	unmarshalErr := json.Unmarshal([]byte((*recvMessagePtr).Data), &loginMes)
	if unmarshalErr != nil {
		return unmarshalErr
	}
	fmt.Printf("[debug-%v]:(%v)\n", info.CurrFuncName(), loginMes)

	// 校验登录信息...
	checkLoginErr := service.CheckLoginMessage(&loginMes)
	var loginResMes obj.LoginResMes
	switch checkLoginErr{
	case nil:
		loginResMes.Code = obj.CODE_LOGIN_SUCCESS
		loginResMes.SearchedUserName = loginMes.UserName
		
	case obj.ERROR_USER_NOTEXIST:
		loginResMes.Code = obj.CODE_USER_NOTEXIST
		loginResMes.ErrorText = obj.ERROR_USER_NOTEXIST.Error()
	case obj.ERROR_USER_WRONG_PASSWD:
		loginResMes.Code = obj.CODE_USER_WRONG_PASSWD
		loginResMes.ErrorText = obj.ERROR_USER_WRONG_PASSWD.Error()
	default:
		loginResMes.Code = obj.CODE_UNKNOWN_IN_SERVER
		loginResMes.ErrorText = obj.ERROR_UNKNOWN_IN_SERVER.Error()
	}
	// 对loginResMes序列化, 构造resMessage
	var resMessage obj.Message
	resMessage.Type = obj.LoginResMesType
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
	// 1.服务器端更新在线列表
	// 2.为本用户建立ResFwd收发管道
	// 3.向其他用户通知此ID已上线
	if loginResMes.Code == obj.CODE_LOGIN_SUCCESS {
		AddOnlineStruct(OnlineStruct{loginMes.UserID, loginMes.UserName, conn})
		
		var resFwdChan chan obj.ResFwdChatMes = make(chan obj.ResFwdChatMes, 1)
		G_ResFwdMap.Store(loginMes.UserID, resFwdChan)
	
		OnlineNotify(loginMes.UserID, loginMes.UserName, true)
	}
	

	// 检查在线列表  ID/总数
	// sum, tmpSli := IterateOnlineMap()
	// for _, tmpStru := range tmpSli{
	// 	fmt.Printf("[debug-%v]:(%v)(%v)\n", info.CurrFuncName(),tmpStru.UserID,sum)
	// }
	
	return nil
}

/*************************************************/
func RegisterProcess(conn net.Conn, recvMessagePtr *obj.Message)error{
	var registerMes obj.RegisterMes
	unmarshalErr := json.Unmarshal([]byte((*recvMessagePtr).Data), &registerMes)
	if unmarshalErr != nil {
		return unmarshalErr
	}
	fmt.Printf("[debug-%v]:(%v)\n", info.CurrFuncName(), registerMes)
	// 尝试注册...
	registerUserErr := service.RegisterUser(&registerMes)
	var registerResMes obj.RegisterResMes
	switch registerUserErr{
	case nil:
		registerResMes.Code = obj.CODE_REGISTER_SUCCESS
		// 注册成功...
	case obj.ERROR_USER_EXISTS:
		registerResMes.Code = obj.CODE_USER_EXISTS
		registerResMes.ErrorText = obj.ERROR_USER_EXISTS.Error()
	case obj.ERROR_USER_ILLEGAL_PASSWD:
		registerResMes.Code = obj.CODE_USER_ILLEGAL_PASSWD
		registerResMes.ErrorText = obj.ERROR_USER_ILLEGAL_PASSWD.Error()
	case obj.ERROR_USER_ILLEGAL_NAME:
		registerResMes.Code = obj.CODE_USER_ILLEGAL_NAME
		registerResMes.ErrorText = obj.ERROR_USER_ILLEGAL_NAME.Error()
	default:
		registerResMes.Code = obj.CODE_UNKNOWN_IN_SERVER
		registerResMes.ErrorText = obj.ERROR_UNKNOWN_IN_SERVER.Error()
	}

	// registerResMes, 构造resMessage
	var resMessage obj.Message
	resMessage.Type = obj.RegisterResMesType
	resMessageDataSli, marshalErr := json.Marshal(registerResMes)
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


func OnlineNotify(userID int, userName string, userStatus bool)error{
	var mes obj.Message
	mes.Type = obj.OnlineMesType
	var onlineNotifyMes obj.OnlineMes = obj.OnlineMes{
		UserPublicInfo : obj.UserPublicInfo{
			UserID : userID,
			UserName : userName,
		},
		UserIsOnline : userStatus,
	}
	mesDataSli, marshalErr := json.Marshal(onlineNotifyMes)
	if marshalErr != nil {
		return marshalErr
	}
	mes.Data = string(mesDataSli)

	_, onlineList := IterateOnlineMap()
	for _, olStruct := range onlineList{
		if olStruct.UserID != userID{
			writeMessageErr := (&mes).WriteMessageStructIntoConn(olStruct.ConnOfUser)
			if writeMessageErr != nil {
				fmt.Printf("[debug-%v]:(%v onlineNoti send to %v)(%v)\n", info.CurrFuncName(), userID, olStruct.UserID, writeMessageErr)
			}
		}
	}
	return nil
}

func ResOnlineListProcess(conn net.Conn, recvMessagePtr *obj.Message)error{
	var reqMes obj.RequestOnlineListMes
	unmarshalErr := json.Unmarshal([]byte((*recvMessagePtr).Data), &reqMes)
	if unmarshalErr != nil {
		return unmarshalErr
	}
	//fmt.Printf("[debug-%v]:(%v)\n", info.CurrFuncName(), reqMes)
	// 遍历得到自己除外的在线列表
	num, onlineListServer:= IterateOnlineMap()
	onlineListClient := make(obj.ResOnlineListMes, num - 1)
	for i,j := 0, 0; j<num;  {
		if onlineListServer[j].UserID != int(reqMes){
			onlineListClient[i].UserID = onlineListServer[j].UserID
			onlineListClient[i].UserName = onlineListServer[j].UserName
			onlineListClient[i].UserIsOnline = true
			i++
			j++
		}else {
			j++
		}
	}
	// registerResMes, 构造resMessage
	var resMessage obj.Message
	resMessage.Type = obj.ResOnlineListMesType
	resMessageDataSli, marshalErr := json.Marshal(onlineListClient)
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