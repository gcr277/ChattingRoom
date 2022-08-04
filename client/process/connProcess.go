package process

import (
	"net"
	"fmt"
	"ChattingRoom/common/info"
)

func DialProcess() net.Conn {
	// dial to server
	conn, dialErr := net.Dial("tcp","localhost:8888")
	if dialErr != nil{
		fmt.Printf("[client-%v]:%v\n", info.CurrFuncName(), dialErr)
	}else{
		fmt.Printf("[client-%v]:connected to server!\n", info.CurrFuncName())
	}
	return conn
}