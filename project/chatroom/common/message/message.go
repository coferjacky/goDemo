package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data""`
}

//定义两个消息，后面有需要再增加

type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code  int    `json:"code"`  //返回状态 500 表示改用户没有注册     200 登录成功
	Error string `json:"error"` //返回错误信息
}

type RegisterMes struct {
	//....
}
