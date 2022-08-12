package service

import (
	"ChattingRoom/common/obj"
	"ChattingRoom/server/dao"
	_"errors"
)


func CheckLoginMessage(loginMes *obj.LoginMes) error {
	var userDao dao.UserDao
	(&userDao).SetPool(dao.GlobalPool)
	userInDB, searchErr := userDao.GetUserByID((*loginMes).UserID)
	if searchErr != nil{
		return searchErr
	}
	if (*loginMes).UserPasswd != userInDB.UserPasswd{
		return obj.ERROR_USER_WRONG_PASSWD
	}
	(*loginMes).UserName = userInDB.UserName
	return nil
}