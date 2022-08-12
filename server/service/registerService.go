package service

import (
	"ChattingRoom/common/obj"
	"ChattingRoom/server/dao"
)

func RegisterUser(registerMes *obj.RegisterMes) error {
	if !IsUserNameLegal((*registerMes).UserName){
		return obj.ERROR_USER_ILLEGAL_NAME
	}
	if !IsPasswordLegal((*registerMes).UserPasswd){
		return obj.ERROR_USER_ILLEGAL_PASSWD
	}
	var userDao dao.UserDao
	(&userDao).SetPool(dao.GlobalPool)
	setUserErr := userDao.SetNewUser(*registerMes)
	if setUserErr != nil{
		return setUserErr
	}else {
		return nil 
	}
}

func IsPasswordLegal(userPasswd string)bool{
	return true
}
func IsUserNameLegal(userName string)bool{
	return true
}