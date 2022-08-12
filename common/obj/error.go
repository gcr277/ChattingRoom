package obj

import (
	"errors"
)

var(
	CODE_LOGIN_SUCCESS = 200
	CODE_REGISTER_SUCCESS = 201

	CODE_USER_NOTEXIST = 300
	ERROR_USER_NOTEXIST = errors.New("[USER]user does not exist!")

	CODE_USER_EXISTS = 301
	ERROR_USER_EXISTS = errors.New("[USER]user already exists!") 	

	CODE_USER_WRONG_PASSWD = 400
	ERROR_USER_WRONG_PASSWD = errors.New("[USER]wrong password!") 	

	CODE_USER_ILLEGAL_PASSWD = 401
	ERROR_USER_ILLEGAL_PASSWD = errors.New("[USER]illegal password!")	

	CODE_USER_ILLEGAL_NAME = 402
	ERROR_USER_ILLEGAL_NAME = errors.New("[USER]illegal username!")	

	CODE_UNKNOWN_IN_SERVER = 500
	ERROR_UNKNOWN_IN_SERVER = errors.New("[SERVER]unknown error in server!")

	CODE_UNABLE_POPCHAN = 600
	ERROR_UNABLE_POPCHAN = errors.New("[CLIENT]can not pop from channel!")

	CODE_UNABLE_PUSHCHAN = 601
	ERROR_UNABLE_PUSHCHAN = errors.New("[CLIENT]can not push into channel!")

	CODE_DST_RECV_SUCCESS = 700
	CODE_SERVER_RECV_FAIL = 701
	ERROR_SERVER_RECV_FAIL = errors.New("[CLIENT]can not send to server!")
	CODE_SERVER_RECV_SUCCESS_DST_RECV_FAIL = 702
	ERROR_SERVER_RECV_SUCCESS_DST_RECV_FAIL = 
			errors.New("[SERVER]sent to server but dst did not receive!")
	

)