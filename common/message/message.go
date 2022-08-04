package message
import (
	"fmt"
	"io"
	"net"
	"ChattingRoom/common/info"
	"encoding/json"
)
// type of Message
const (
	LoginMesType = iota
	LoginResMesType
	RegisterMesType
	RegisterResMesType
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
	fmt.Printf("[debug-%v]:%v\n", info.CurrFuncName(), string(resMessageSli))

	_, writeErr := conn.Write(resMessageSli)
	if writeErr != nil{
		fmt.Printf("[debug-%v]:%v\n", info.CurrFuncName(), writeErr)
		return writeErr
	}
	return nil
}
/*****************************************************************/


// data of Message

type User struct{
	UserID int			`json:"userID"`
	UserPasswd string	`json:"userPasswd"`
	UserName string		`json:"userName"`
}
type LoginMes User

type LoginResMes struct{
	Code int 		`json:"code"`	// 500:未注册  200:注册成功
	ErrorText string	`json:"errortext"`
}

type RegisterMes struct{

}

type RegisterResMes struct{
	
}