package view

import (
	"fmt"
	"ChattingRoom/common/obj"
)

func DispNewMessage(recvChatMesPtr *obj.ChatMes){
	fmt.Printf("\n<<<<<<<<<<<<<<<<< new message <<<<<<<<<<<<<<<<<\n")
	fmt.Printf("FROM:ID[%v]\tNAME[%v]\t\n"+"%v\n", 
			recvChatMesPtr.SrcUser.UserID, recvChatMesPtr.SrcUser.UserName,
			recvChatMesPtr.Content)
	fmt.Printf("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^\n")
}