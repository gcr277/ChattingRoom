package main

import (
	
	"fmt"
	_ "github.com/garyburd/redigo/redis"
	_ "ChattingRoom/common/message"
	"ChattingRoom/common/info"
	"net"

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
			go handle(conn)
		}
	}
}
