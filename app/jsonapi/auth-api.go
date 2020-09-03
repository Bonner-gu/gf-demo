package jsonapi

type Token struct {
	Source string `json:"source"` // 来源
	Ip     string `json:"ip"`     // IP地址
	Mk     string `json:"mk"`     // 机器码
	Uin    int    `json:"uin"`    // 用户唯一标识
	Role   string `json:"role"`   // 用户角色
	Skey   string `json:"skey"`   // 登陆态
}
