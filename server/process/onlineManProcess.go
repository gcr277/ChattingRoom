package process

import (
	"net"
	_"fmt"
	_"ChattingRoom/common/info"
	"sync"
)
type OnlineStruct struct{
	UserID int
	UserName string
	ConnOfUser net.Conn
}

var (
	G_OnlineMap sync.Map
)
const (
	APPRO_ONLINE_NUM = 500
)

func AddOnlineStruct(olStruct OnlineStruct){
	G_OnlineMap.Store(olStruct.UserID, olStruct)
}
func IterateOnlineMap()(int, []OnlineStruct){
	cnt := 0
	olSli := make([]OnlineStruct, 0, APPRO_ONLINE_NUM)
	G_OnlineMap.Range(func(key, value interface{}) bool {
		cnt++
		olSli = append(olSli, value.(OnlineStruct))
		//fmt.Printf("[debug-%v]:(%v)\n", info.CurrFuncName(), value.(OnlineStruct))
		return true
	})
	return cnt, olSli
}
func GetOnlineStructByID(userID int)(OnlineStruct, bool) {
	value, contains := G_OnlineMap.Load(userID)
	return value.(OnlineStruct), contains
}
func DelOnlineStructByID(userID int) {
	G_OnlineMap.Delete(userID)
}