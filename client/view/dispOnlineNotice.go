package view

import (
	"fmt"
	"ChattingRoom/common/obj"
)

func DispOnlineNotice(recvOnlineMes obj.OnlineMes){
	if recvOnlineMes.UserIsOnline{
		fmt.Printf(	"\n"+
				"<<<<<<<<<<<<<<<<< notice <<<<<<<<<<<<<<<<<\n"+
				  "\t[%v]已上线\tID:[%v]\n"+
				"^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^\n",
				recvOnlineMes.UserName, recvOnlineMes.UserID)
	}else{
		fmt.Printf(	"\n"+
				"<<<<<<<<<<<<<<<<< notice <<<<<<<<<<<<<<<<<\n"+
				  "\t[%v]已下线\tID:[%v]\n"+
				"^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^\n",
				recvOnlineMes.UserName, recvOnlineMes.UserID)
	}
}