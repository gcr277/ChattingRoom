package process

import (
	"net"
	"ChattingRoom/common/obj"
	"ChattingRoom/common/info"
	"fmt"
	"encoding/json"
)

func SendMessage(gconn net.Conn, gUser obj.User, 
	ol obj.OnlineMes, contentPtr *string)error{

	var chatMes obj.ChatMes
	chatMes.SrcUser.UserID = gUser.UserID
	chatMes.SrcUser.UserName = gUser.UserName
	chatMes.DstUser.UserID = ol.UserID
	chatMes.DstUser.UserName = ol.UserName
	chatMes.Content = *contentPtr

	var SendMessage obj.Message
	SendMessage.Type = obj.ChatMesType
	
	// 序列化->SendMessage.Data
	SendMessageDataSli, marshalErr := json.Marshal(chatMes)
	if marshalErr != nil{
		fmt.Printf("[%v]:%v\n", info.CurrFuncName(), marshalErr)
		return marshalErr
	}
	SendMessage.Data = string(SendMessageDataSli)
	writeMessageErr := (&SendMessage).WriteMessageStructIntoConn(gconn)
	if writeMessageErr != nil {
		fmt.Printf("[%v]:%v\n", info.CurrFuncName(), writeMessageErr)
		return writeMessageErr
	}

	// 已在serverPush协程读取服务器回执并放入管道
	select {
		case data := <- ResChatMesChan:
			var resChatMes obj.ResChatMes
			unmarshalErr := json.Unmarshal([]byte(data), &resChatMes)
			if unmarshalErr != nil {
				fmt.Printf("[%v]:(%v)\n", info.CurrFuncName(), unmarshalErr)
			}
			if resChatMes.Code == obj.CODE_DST_RECV_SUCCESS{
				fmt.Printf("[%v]:对方已收到信息！\n", info.CurrFuncName())
			}else if resChatMes.Code == obj.CODE_SERVER_RECV_SUCCESS_DST_RECV_FAIL{
				fmt.Printf("[%v]:对方未收到信息，服务器已转储！\n", info.CurrFuncName())
			}else {
				fmt.Printf("[%v]:(%v)\n", info.CurrFuncName(), resChatMes.ErrorText)
			}
		//default:
	}

	

	return nil
}