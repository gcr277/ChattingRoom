package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/garyburd/redigo/redis"
	"go_practice/projChattingRoom/common/message"
	"go_practice/projChattingRoom/common/info"
	"io"
	"net"
	"errors"
)

func main() {

	// listen
	listen, ListenErr := net.Listen("tcp", ":8888")
	if ListenErr != nil {
		fmt.Printf("[server-%v]:[listen fail...](%v)\n", info.CurrFuncName(), ListenErr)
		return
	}
	defer listen.Close()
	// wait for connection
	for {
		conn, connErr := listen.Accept()
		if connErr != nil {
			fmt.Printf("[server-%v]:[connect fail...](%v)\n", info.CurrFuncName(), ListenErr)
		} else {
			fmt.Printf("[server-%v]:[client ip=%v]\n", info.CurrFuncName(), conn.RemoteAddr().String())
			// goroutine
			go process(conn)
		}
	}
}

/*
 * 处理客户端连接发来的消息
 */
func process(conn net.Conn) {
	defer conn.Close()
	for {
		var recvMessage message.Message
		readMessageErr := readMessage(conn, &recvMessage)
		if readMessageErr != nil {
			fmt.Printf("[server-%v]:(%v)\n", info.CurrFuncName(), readMessageErr)
			return
		}
		fmt.Printf("[debug-%v]:%v\n", info.CurrFuncName(), recvMessage)
		//
		handleMesErr := handleMes(conn, &recvMessage)
		if readMessageErr != nil {
			fmt.Printf("[server-%v]:(%v)\n", info.CurrFuncName(),handleMesErr)
			return
		}

	}
}

/*
 * 将字符切片反序列化为消息结构体
 */
func readMessage(conn net.Conn, recvMessagePtr *message.Message) error {
	buf := make([]byte, 4096)
	//fmt.Printf("[debug]:(%v)[a new buffer is waiting message fron client...]\n",conn.RemoteAddr().String())
	n, readErr := conn.Read(buf)
	if readErr != nil {
		if readErr != io.EOF{
			fmt.Printf("[server-%v]:[read error ...](%v)(%v)\n", info.CurrFuncName(), readErr, conn.RemoteAddr().String())
		}
		return readErr
	}
	//fmt.Printf("[server goroutine]:%v(%v)\n", string(buf[0:n]), conn.RemoteAddr().String())

	unmarshalErr := json.Unmarshal(buf[0:n], recvMessagePtr)
	if unmarshalErr != nil {
		fmt.Printf("[server-%v]:[unmarshal error ...](%v)\n", info.CurrFuncName(), unmarshalErr)
		return unmarshalErr
	}
	return nil
}

/*
 * 根据消息类型，调用不同的处理函数来进行反序列化
 */
 func handleMes(conn net.Conn, recvMessagePtr *message.Message)error{
	switch (*recvMessagePtr).Type {
		case message.LoginMesType:
			handleMesErr := handleLoginMes(conn, recvMessagePtr)
			if handleMesErr != nil{
				return handleMesErr
			}
		//case message.RegisterMesType:

		default:
			unknownTypeErr := errors.New("Unknown message type")
			fmt.Printf("[server-%v]:(%v)", info.CurrFuncName(), unknownTypeErr)
			return unknownTypeErr
	}
	return nil
 }

 func handleLoginMes(conn net.Conn, recvMessagePtr *message.Message)error{
	var loginMes message.LoginMes
	unmarshalErr := json.Unmarshal([]byte((*recvMessagePtr).Data), &loginMes)
	if unmarshalErr != nil {
		return unmarshalErr
	}
	fmt.Printf("[debug-%v]:(%v)\n", info.CurrFuncName(), loginMes)

	// 校验登陆信息...

	// 如果登录信息正确
	var loginResMes message.LoginResMes
	loginResMes.Code = 200
	loginResMes.Error = "***"
		// 对loginResMes序列化
	var resMessage message.Message
	resMessage.Type = message.LoginResMesType
	resMessageDataSli ,marshalErr := json.Marshal(loginResMes)
	if marshalErr != nil {
		return marshalErr
	}
	resMessage.Data = string(resMessageDataSli)
		// 对resMessage序列化
	resMessageSli, marshalErr := json.Marshal(resMessage)
	if marshalErr != nil {
		return marshalErr
	}
	fmt.Printf("[debug-%v]:%v\n", info.CurrFuncName(), string(resMessageSli))

	_, writeErr := conn.Write(resMessageSli)
	if writeErr != nil{
		fmt.Printf("[server-%v]:%v\n", info.CurrFuncName(), writeErr)
		return writeErr
	}
	// 发送，还是封装吧。。。
	

	return nil
 }