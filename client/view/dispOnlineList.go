package view

import (
	"fmt"
	"ChattingRoom/common/obj"
)

func DispOnlineList(gOnlineList obj.ResOnlineListMes){
	fmt.Printf("\n<<<<<<<<<<<<<<<<< online list <<<<<<<<<<<<<<<<<\n")
	fmt.Printf("NO.\tID\t\tNAME\t\t\n")
	for i, ol := range gOnlineList{
		fmt.Printf("%v\t%v\t\t%v\t\t\n", i+1, ol.UserID, ol.UserName)
	}
	fmt.Printf("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^\n")
				
}