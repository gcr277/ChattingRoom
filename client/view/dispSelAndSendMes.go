package view

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func DispSelAndSendMes()(int, string){
	var num int
	var content string
	fmt.Printf("----------------- send message -----------------\n")
	for {
		fmt.Printf("select user (NO.):\n")
		fmt.Scanf("%d\n", &num)
		if selectNumIsLegal(num){
			break
		}else{
			fmt.Printf("illegal num!\n")
		}
	}
	fmt.Printf("input content:\n")
	reader := bufio.NewReader(os.Stdin)
	content, _ = reader.ReadString('\n')
	content = strings.TrimSuffix(content, "\n")
	return num, content
}
func selectNumIsLegal(int)bool{
	return true
}