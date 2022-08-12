package service

import (
	"fmt"
	"ChattingRoom/common/info"
)

func DumpForDst(srcUserID int, dstUserID int)error{

	fmt.Printf("[debug-%v]:message from %v to %v dumped!\n", 
					info.CurrFuncName(), srcUserID, dstUserID)
	return nil
}
