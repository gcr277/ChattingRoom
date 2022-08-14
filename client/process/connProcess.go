package process

import (
	"net"
	"fmt"
	"ChattingRoom/common/info"
)

func DialProcess() (net.Conn, error) {
	// dial to server
	fmt.Printf("输入服务器IP:port :\n")
	var ip string
	fmt.Scanf("%s\n", &ip)
	conn, dialErr := net.Dial("tcp",ip)
	if dialErr != nil{
		fmt.Printf("[client-%v]:%v\n", info.CurrFuncName(), dialErr)
	}else{
		fmt.Printf("[client-%v]:connected to server!\n", info.CurrFuncName())
	}
	return conn, dialErr
}