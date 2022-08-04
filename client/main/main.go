package main
import (
	"ChattingRoom/client/view"
	"ChattingRoom/client/menuSwitch"
	"fmt"
	"ChattingRoom/common/info"
) 



func main(){
	var keyWelcome int = view.LoopDispWelcome()
	menuSwitch.WelcomeSwitch(keyWelcome)

	fmt.Printf("[client-%v]:Thanks for using this software, goodbye!\n", info.CurrFuncName())
	return

}