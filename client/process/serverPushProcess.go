package process

import (
	"net"
	"fmt"
	"ChattingRoom/common/obj"
	"ChattingRoom/common/info"
	"ChattingRoom/client/view"
	"encoding/json"
)
var (
	GOLListChan chan string = make(chan string, 1)
	ResChatMesChan chan string = make(chan string, 1)

)
func ServerPushProcess(gconn net.Conn, gUser obj.User){
	var serverPushMes obj.Message
	for {
		//fmt.Printf("[debug-%v]:goroutine waiting for serverpush\n",info.CurrFuncName())
		readMessageErr := (&serverPushMes).ReadMessageStructFromConn(gconn)
		if readMessageErr != nil {
			fmt.Printf("[%v]:(%v)\n", info.CurrFuncName(), readMessageErr)
		}
		//fmt.Printf("[debug-%v]:(%v)\n", info.CurrFuncName(), serverPushMes)
		switch serverPushMes.Type{
		case obj.OnlineMesType:  // 被动接受的推送直接处理
			var recvOnlineMes obj.OnlineMes
			unmarshalErr := json.Unmarshal([]byte(serverPushMes.Data), &recvOnlineMes)
			if unmarshalErr != nil {
				fmt.Printf("[%v]:(%v)\n", info.CurrFuncName(), unmarshalErr)
			}
			view.DispOnlineNotice(recvOnlineMes)
			
		case obj.ResOnlineListMesType: // 主动请求后返回的推送放入管道，由请求方取出
			GOLListChan <- serverPushMes.Data
		case obj.ChatMesType:
			var resFwdChatMes obj.ResFwdChatMes
			resFwdChatMes.DstUserID = gUser.UserID
			var recvChatMes obj.ChatMes
			unmarshalErr := json.Unmarshal([]byte(serverPushMes.Data), &recvChatMes)
			if unmarshalErr != nil {
				fmt.Printf("[%v]:(%v)\n", info.CurrFuncName(), unmarshalErr)
				resFwdChatMes.Code = obj.CODE_SERVER_RECV_SUCCESS_DST_RECV_FAIL
				resFwdChatMes.ErrorText = obj.ERROR_SERVER_RECV_SUCCESS_DST_RECV_FAIL.Error()
			}else {
				resFwdChatMes.Code = obj.CODE_DST_RECV_SUCCESS
				resFwdChatMes.ErrorText = ""
				view.DispNewMessage(&recvChatMes)
			}
			
			var SendMessage obj.Message
			SendMessage.Type = obj.ResFwdChatMesType
			// 序列化->SendMessage.Data
			SendMessageDataSli, marshalErr := json.Marshal(resFwdChatMes)
			if marshalErr != nil{
				fmt.Printf("[%v]:%v\n", info.CurrFuncName(), marshalErr)
			}
			SendMessage.Data = string(SendMessageDataSli)
			writeMessageErr := (&SendMessage).WriteMessageStructIntoConn(gconn)
			if writeMessageErr != nil {
				fmt.Printf("[%v]:%v\n", info.CurrFuncName(), writeMessageErr)
			}
		case obj.ResChatMesType:
			ResChatMesChan <- serverPushMes.Data
			
		default :

		}
	}
}