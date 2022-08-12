package process

import (
	_"fmt"
	_"ChattingRoom/common/info"
	"ChattingRoom/common/obj"
	"ChattingRoom/server/service"
	"encoding/json"
	"net"
	"errors"
)

func ForwardOrDumpProcess(conn net.Conn, recvMessagePtr *obj.Message, resFwdChatMesChan chan string)(fodErr error){
	var resChatMes obj.ResChatMes
	
	var chatMes obj.ChatMes
	unmarshalErr := json.Unmarshal([]byte((*recvMessagePtr).Data), &chatMes)
	if unmarshalErr != nil {
		resChatMes.Code = obj.CODE_UNKNOWN_IN_SERVER
		resChatMes.ErrorText = obj.ERROR_UNKNOWN_IN_SERVER.Error()
		fodErr = unmarshalErr
	}
	//fmt.Printf("[debug-%v]:(%v)\n", info.CurrFuncName(), chatMes)
	dstStruct, isOnline := GetOnlineStructByID(chatMes.DstUser.UserID)

	
	if isOnline{ // 直接向目标转发
		fwdErr := ForwardProcess(dstStruct.ConnOfUser, &chatMes, resFwdChatMesChan)
		//fmt.Printf("**************%v\n", fwdErr)
		if fwdErr != nil { //转发失败，转储到离线消息
			dumpErr := service.DumpForDst(chatMes.SrcUser.UserID, chatMes.DstUser.UserID)
			if dumpErr != nil {
				resChatMes.Code = obj.CODE_UNKNOWN_IN_SERVER
				resChatMes.ErrorText = obj.ERROR_UNKNOWN_IN_SERVER.Error()
				fodErr = dumpErr
			}else {
				resChatMes.Code = obj.CODE_SERVER_RECV_SUCCESS_DST_RECV_FAIL
				resChatMes.ErrorText = obj.ERROR_SERVER_RECV_SUCCESS_DST_RECV_FAIL.Error()
				fodErr = obj.ERROR_SERVER_RECV_SUCCESS_DST_RECV_FAIL
			}
		}else{ // 转发成功
			resChatMes.Code = obj.CODE_DST_RECV_SUCCESS
			resChatMes.ErrorText = ""
			fodErr = nil
		}
		
	}else { // 转储到目标用户的离线消息文件
		dumpErr := service.DumpForDst(chatMes.SrcUser.UserID, chatMes.DstUser.UserID)
		if dumpErr != nil {
			resChatMes.Code = obj.CODE_UNKNOWN_IN_SERVER
			resChatMes.ErrorText = obj.ERROR_UNKNOWN_IN_SERVER.Error()
			fodErr = dumpErr
		}else {
			resChatMes.Code = obj.CODE_SERVER_RECV_SUCCESS_DST_RECV_FAIL
			resChatMes.ErrorText = obj.ERROR_SERVER_RECV_SUCCESS_DST_RECV_FAIL.Error()
			fodErr = obj.ERROR_SERVER_RECV_SUCCESS_DST_RECV_FAIL
		}
	}

	var resMes obj.Message
	resMes.Type = obj.ResChatMesType
	resMesDataSli, marshalErr := json.Marshal(resChatMes)
	if marshalErr != nil {
		fodErr =  marshalErr
	}
	resMes.Data = string(resMesDataSli)

	writeMessageErr := (&resMes).WriteMessageStructIntoConn(conn)
	if writeMessageErr != nil {
		fodErr =  writeMessageErr
	}
	return fodErr
}

func ForwardProcess(connOfDst net.Conn, chatMesPtr *obj.ChatMes, resFwdChatMesChan chan string)error{
	var fwdMessage obj.Message
	fwdMessage.Type = obj.ChatMesType
	fwdMessageDataSli, marshalErr := json.Marshal(*chatMesPtr)
	if marshalErr != nil {
		return marshalErr
	}
	fwdMessage.Data = string(fwdMessageDataSli)
	// 对fwdMessage序列化并发送
	writeMessageErr := (&fwdMessage).WriteMessageStructIntoConn(connOfDst)
	if writeMessageErr != nil {
		return writeMessageErr
	}

	select{
		case data := <- resFwdChatMesChan:
			var resFwdChatMes obj.ResFwdChatMes
			unmarshalErr := json.Unmarshal([]byte(data), &resFwdChatMes)
			if unmarshalErr != nil {
				return unmarshalErr
			}
			if resFwdChatMes.Code == obj.CODE_DST_RECV_SUCCESS{
				return nil
			}else{
				return errors.New(resFwdChatMes.ErrorText)
			}
		//default:
	}
	//return obj.ERROR_UNKNOWN_IN_SERVER
}
