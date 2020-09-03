package jsonapi

type AdminSignupReq struct {
	Mkid    string `v:"required|length:16,64#机器码不能为空|机器码格式不合法" json:"mkid"`              //机器码
	Imgcode string `v:"required#验证的数值不能为空" json:"imgcode"`                               //验证数值
	Phone   string `v:"required|length:11,11#账号（手机号码）不能为空|账号(手机号码)长度为:min" json:"phone"` //账号(手机号)
	Newpw   string `v:"required#密码不能为空" json:"newpw"`                                    //密码
	Rightpw string `v:"required|same:Newpw#请确认密码|两次密码不一致" json:"rightpw"`                //确认密码
	Ip      string `v:"required|ipv4#ip不能为空|ip格式不合法" json:"ip"`                          // ip
	Source  string `v:"required|length:1,8#来源不能为空|来源格式不合法" json:"source"`                // 来源
}

type AdminSigninReq struct {
	Mkid    string `v:"required|length:16,64#机器码不能为空|机器码格式不合法" json:"mkid"`              //机器码
	Imgcode string `v:"required#验证的数值不能为空" json:"imgcode"`                               //验证数值
	Phone   string `v:"required|length:11,11#账号（手机号码）不能为空|账号(手机号码)长度为:min" json:"phone"` //账号(手机号)
	Pw      string `v:"required#密码不能为空" json:"pw"`                                       //密码
	Source  string `v:"required|length:1,8#来源不能为空|来源格式不合法" json:"source"`                // 来源
	Ip      string `v:"required|ipv4#ip不能为空|ip格式不合法" json:"ip"`                          // ip
}

type GetAdminInfoReq struct {
	Uin    int64  `v:"required#admin唯一标识不能为空" json:"uin"`                //admin唯一标识
	Source string `v:"required|length:1,8#来源不能为空|来源格式不合法" json:"source"` // 来源
}

type AdminInfoRet struct {
	Name  string //admin姓名
	Phone string //电话号码
	Logo  string //头像
	Sign  string //个性签名
}
type ModifyAdminInfoReq struct {
	Name   string
	Phone  string
	Logo   string
	Sign   string
	Uin    int64  `v:"required#admin唯一标识不能为空" json:"uin"`                //admin唯一标识
	Source string `v:"required|length:1,8#来源不能为空|来源格式不合法" json:"source"` // 来源
}

type ModifyPwReq struct {
	Oldpw   string `v:"required#旧密码密码不能为空" json:"oldpw"`                  //旧密码
	Newpw   string `v:"required#新密码不能为空" json:"newpw"`                    //新密码
	Rightpw string `v:"required|same:Newpw#请确认密码|两次密码不一致" json:"rightpw"` //确认密码
	Uin     int64  `v:"required#admin唯一标识不能为空" json:"uin"`                //admin唯一标识
	Token   *Token
	Source  string `v:"required|length:1,8#来源不能为空|来源格式不合法" json:"source"` // 来源
}

type ModifyNewPwReq struct {
	Mkid    string `v:"required|length:16,64#机器码不能为空|机器码格式不合法" json:"mkid"` //机器码
	Imgcode string `v:"required#验证的数值不能为空" json:"imgcode"`                  //验证数值
	Newpw   string `v:"required#新密码不能为空" json:"newpw"`                      //新密码
	Rightpw string `v:"required|same:Newpw#请确认密码|两次密码不一致" json:"rightpw"`   //确认密码
	Uin     int64  `v:"required#admin唯一标识不能为空" json:"uin"`                  //admin唯一标识
	Token   *Token
	Source  string `v:"required|length:1,8#来源不能为空|来源格式不合法" json:"source"` // 来源
}
