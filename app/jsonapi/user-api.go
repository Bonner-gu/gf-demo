package jsonapi

type UserStatusReq struct {
	Name   string
	Tel    string
	Idcard string
	Status int
	Page   int32
	Size   int32
	Source string `v:"required|length:1,8#来源不能为空|来源格式不合法" json:"source"` // 来源
}

type UserInfo struct {
	Uid        int64
	Logo       string
	Name       string
	Idcard     string
	Phone      string
	Source     string
	RegisterIp string
	CreateTime string
	Status     int
}

type ModifyUser struct {
	Uin    int64
	Uid    int64  `v:"required#用户id不能为空" json:"uid"`                     //用户流水id
	Source string `v:"required|length:1,8#来源不能为空|来源格式不合法" json:"source"` // 来源
	Status int    `v:"required#操作状态不能为空" json:"status"`
}

type UserFormReq struct {
	Name    string //搜索词
	Company string //搜索词
	Label   string //搜索词
	Page    int32  //页数
	Size    int32  //条数
	Status  int    //表单状态
	Source  string `v:"required|length:1,8#来源不能为空|来源格式不合法" json:"source"` // 来源
}

type UserFormInfo struct {
	Uid       int64  `orm:"uid,primary"   json:"uid"`        //用户唯一标识
	UserName  string `orm:"nickname"    json:"name"`         //用户名
	Cid       int64  `orm:"cid,primary"   json:"cid"`        // 公司唯一标识
	Company   string `orm:"company"   json:"company"`        // 公司名
	Status    string `orm:"status"    json:"status"`         //表单审核状态
	LabelName string `orm:"name"    json:"name"`             //标签名
	SummitWay string `orm:"submit_way"    json:"submit_way"` //用户的提交方式
	Info      string `orm:"info"    json:"info"`             //任务描述
	Mark      string `orm:"mark"    json:"mark"`             //任务描述
	Integral  int    `orm:"integral"    json:"integral"`     //任务积分
}
