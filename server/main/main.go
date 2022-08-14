package main

import (
	
	"fmt"
	_ "github.com/garyburd/redigo/redis"
	_ "ChattingRoom/common/obj"
	"ChattingRoom/common/info"
	"ChattingRoom/server/dao"
	"time"
	"net"

)


func main() {
	// init redis conn pool
	dao.GlobalPoolInit("tcp", ":6379", 16, 0, 300 * time.Second)
	defer dao.GlobalPoolClose()
	// listen
	listen, ListenErr := net.Listen("tcp", "192.168.157.137:8888")
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
