package jsonapi

type CaptchaReq struct {
	Mk     string `v:"required|length:16,64#机器码不能为空|机器码格式不合法" json:"mk"` // 机器码
	Ip     string `v:"required|ipv4#ip不能为空|ip格式不合法" json:"ip"`           // ip
	Source string `v:"required|length:1,8#来源不能为空|来源格式不合法" json:"source"` // 来源
}

type CaptchaCheckReq struct {
	Mkid  string // 机器码
	Value string // 验证值
}
