package message

// type of Message
const (
	LoginMesType = iota
	LoginResMesType
	RegisterMesType
	RegisterResMesType
)

type Message struct{
	Type int		`json:"type"`
	Data string		`json:"data"`

}

// data of Message
type LoginMes struct{
	UserID int			`json:"userID"`
	UserPasswd string	`json:"userPasswd"`
	UserName string		`json:"userName"`
}

type LoginResMes struct{
	Code int 		`json:"code"`	// 500:未注册  200:注册成功
	Error string	`json:"error"`
}

type RegisterMes struct{

}

type RegisterResMes struct{
	
}