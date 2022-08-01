package main
import (
	"fmt"
	"go_practice/projChattingRoom/common/info"
) 

var userID int
var userPasswd string
func main(){
	//
	var key int
	var loop bool = true
	//
	for loop{
		fmt.Printf("----------------- welcome -----------------\n")
		fmt.Printf(	"\t1 login\n"+
					"\t2 register\n"+
					"\t3 exit\n"+
					"input number: ")

		fmt.Scanf("%d", &key)
		switch key{
			case 1 :
				fmt.Printf("----------------- login -----------------\n")
				loop = false
			case 2 :
				fmt.Printf("----------------- register -----------------\n")
				loop = false
			case 3 :
				fmt.Printf("----------------- exit -----------------\n")
				return 
			default :
				fmt.Printf("[client-%v]:illegal input, input again: ", info.CurrFuncName())
		}
		
	}
	if key == 1{
		fmt.Printf("请输入用户ID:")
		fmt.Scanf("%d", &userID)
		fmt.Printf("请输入用户密码:")
		fmt.Scanf("%s", &userPasswd)

		//login.go
		loginErr := login(userID, userPasswd)
		if loginErr != nil{
			fmt.Printf("[client-%v]:登录失败(%v)\n", info.CurrFuncName(), loginErr)
		}else{
			fmt.Printf("[client-%v]:登陆成功!\n", info.CurrFuncName())
		}
	}else if key == 2{

		// register.go
	}

}