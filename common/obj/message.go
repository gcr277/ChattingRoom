package obj
import (
	"fmt"
	"io"
	"net"
	"ChattingRoom/common/info"
	"encoding/json"
)
// type of Message
const (
	LoginMesType = iota				//0
	LoginResMesType					//1

	RegisterMesType					//2
	RegisterResMesType				//3
	
	OnlineMesType					//4

	RequestOnlineListMesType		//5
	ResOnlineListMesType			//6

	ChatMesType						//7
	ResChatMesType					//8
	ResFwdChatMesType				//9
)

/***********************************************************************/
type Message struct{
	Type int		`json:"type"`
	Data string		`json:"data"`

}
// 从conn读到字符切片，反序列化为消息结构体
func (this *Message)ReadMessageStructFromConn(conn net.Conn) error {
	buf := make([]byte, 4096)
	n, readErr := conn.Read(buf)
	if readErr != nil {
		if readErr != io.EOF{
			fmt.Printf("[debug-%v]:[read error ...](%v)(%v)\n", info.CurrFuncName(), readErr, conn.RemoteAddr().String())
		}
		return readErr
	}

	unmarshalErr := json.Unmarshal(buf[0:n], this)
	if unmarshalErr != nil {
		fmt.Printf("[debug-%v]:[unmarshal error ...](%v)\n", info.CurrFuncName(), unmarshalErr)
		return unmarshalErr
	}
	return nil
}
func (this *Message)WriteMessageStructIntoConn(conn net.Conn) error {
	resMessageSli, marshalErr := json.Marshal(*this)
	if marshalErr != nil {
		return marshalErr
	}
	//fmt.Printf("[debug-%v]:%v\n", info.CurrFuncName(), string(resMessageSli))

	_, writeErr := conn.Write(resMessageSli)
	if writeErr != nil{
		fmt.Printf("[debug-%v]:%v\n", info.CurrFuncName(), writeErr)
		return writeErr
	}
	return nil
}
/*****************************************************************/

// Login Message
type LoginMes struct{
	User
}
type LoginResMes struct{
	Code int 				`json:"code"`	//200登录成功,300未注册,400密码错误
	ErrorText string		`json:"errortext"`
	SearchedUserName string	`json:"searchedUserName"`
}
// Reg Message
type RegisterMes struct{
	User
}

type RegisterResMes struct{
	Code int 				`json:"code"`	
	ErrorText string		`json:"errortext"`
}

type RequestOnlineListMes int // userID
type ResOnlineListMes []OnlineMes //

type ChatMes struct{
	SrcUser UserPublicInfo
	DstUser UserPublicInfo
	Content string
}

type ResChatMes struct{
	Code int				`json:"code"`	
	ErrorText string		`json:"errortext"`
}

type ResFwdChatMes struct{
	DstUserID int 			`json:"dstUserID"`
	Code int				`json:"code"`	
	ErrorText string		`json:"errortext"`
}
