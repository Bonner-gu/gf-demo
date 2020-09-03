package jsonapi

type TestReq struct {
	Mk     string `v:"required|length:6,64#账号不能为空|mk长度6,64" json:"mk"`              // 机器码
	Ip     string `v:"required|ipv4#ip不能为空|ip格式不合法" json:"ip"`                      // ip
	Source string `v:"required|length:1,8#来源不能为空|来源格式不合法" json:"source"`            // source 来源
	Phone  string `v:"required-with:phone|phone#phone不能为空|phone格式不合法" json:"phone"` // phone 来源
}
