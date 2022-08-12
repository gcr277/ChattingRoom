package obj

type User struct{
	UserID int			`json:"userID"`
	UserName string		`json:"userName"`
	UserPasswd string	`json:"userPasswd"`
}

type UserPublicInfo struct{
	UserID int			`json:"userID"`
	UserName string		`json:"userName"`
}

type OnlineMes struct{
	UserPublicInfo
	UserIsOnline bool		`json:"userStatus"`
}