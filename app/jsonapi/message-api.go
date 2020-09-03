package jsonapi

import "github.com/gogf/gf/os/gtime"

type SendMessageInfoReq struct {
	Source  string `v:"required|length:1,8#来源不能为空|来源格式不合法" json:"source"` // 来源
	Status  int    `v:"required#发送对象不能为空" json:"titel"`
	Title   string `v:"required#标题不能为空" json:"message"`
	Message string `v:"required#内容不能为空" json:"uin"`
	Uin     int64  `v:"required#admin唯一标识不能为空" json:"uin"`
}

type MessageInfo struct {
	Mid      int64 //消息id
	Logo     string
	Title    string      //标题
	Message  string      //内容
	Status   int         //是否已读状态
	Type     int         //消息类型
	SendTime *gtime.Time //发送时间
}

//消息参数
type MessageProfileReq struct {
	Uin    int64  `v:"required#admin唯一标识不能为空" json:"uin"`
	Source string `v:"required|length:1,8#来源不能为空|来源格式不合法" json:"source"` // 来源
	Mid    int64  `v:"required#Mid参数不能为空"`
}

//消息全部已读参数
type MessageModifyReq struct {
	Uin    int64 `v:"required#admin唯一标识不能为空" json:"uin"`
	Type   int
	Page   int32
	Size   int32
	Source string `v:"required|length:1,8#来源不能为空|来源格式不合法" json:"source"` // 来源
}

//消息详情返回参数
type MessageProfilInfo struct {
	Title    string
	Message  string
	SendTime *gtime.Time
}
