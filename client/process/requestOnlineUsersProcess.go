package process

import (
	"net"
	"ChattingRoom/common/obj"
	"ChattingRoom/common/info"
	"fmt"
	"encoding/json"
)

func RequestOnlineUsers(gconn net.Conn, userID int, gOnlineListPtr *obj.ResOnlineListMes)error{
	var reqMes obj.Message
	reqMes.Type = obj.RequestOnlineListMesType
	var reqMesStru obj.RequestOnlineListMes = obj.RequestOnlineListMes(userID)
	reqMesDataSli, marshalErr := json.Marshal(reqMesStru)
	if marshalErr != nil{
		fmt.Printf("[%v]:%v\n", info.CurrFuncName(), marshalErr)
		return marshalErr
	}
	reqMes.Data = string(reqMesDataSli)
	writeMessageErr := (&reqMes).WriteMessageStructIntoConn(gconn)
	if writeMessageErr != nil {
		fmt.Printf("[%v]:%v\n", info.CurrFuncName(), writeMessageErr)
		return writeMessageErr
	}
	select {
	case retMessageData := <- GOLListChan:
		unmarshalErr := json.Unmarshal([]byte(retMessageData), gOnlineListPtr)
		if unmarshalErr != nil {
			fmt.Printf("[%v]:(%v)\n", info.CurrFuncName(), unmarshalErr)
			return unmarshalErr
		}
	// default:
	// 	return obj.ERROR_UNABLE_POPCHAN
	}
	
	return nil

}