package jsonapi

import "github.com/gogf/gf/os/gtime"

type CompanyStatusReq struct {
	Company  string
	Name     string
	Idcard   string
	Industry string
	Tel      string
	Page     int32
	Size     int32
	Status   int
	Source   string `v:"required|length:1,8#来源不能为空|来源格式不合法" json:"source"` // 来源
}

type CompanyInfo struct {
	Companyid  int64
	Logo       string
	Name       string
	Idcard     string
	Phone      string
	Company    string
	Industry   string
	Address    string
	Email      string
	Website    string
	LicenseUrl string
	Sorce      string
	ApplyIp    string
	CreateTime string
	Examiner   string
	Status     int
}

type ModifyCompany struct {
	Uin       int64
	Companyid int64  `v:"required#公司id不能为空" json:"companyid"`               //公司id
	Status    int    `v:"required#审核结果不能为空" json:"stauts"`                  //审核状态（1：通过，-1：驳回）
	Source    string `v:"required|length:1,8#来源不能为空|来源格式不合法" json:"source"` // 来源
}

type MessageInfoReq struct {
	SendId   int64
	ReId     int64
	Logo     string
	Title    string
	Message  string
	Status   int
	Type     int
	SendTime *gtime.Time
}

type CompanyMailListReq struct {
	Companyid int64  `v:"required#公司id不能为空" json:"companyid"`               //公司id
	Source    string `v:"required|length:1,8#来源不能为空|来源格式不合法" json:"source"` // 来源
}

type CompanyMemberReq struct {
	Companyid int64 `v:"required#公司id不能为空" json:"companyid"` //公司id
	Bid       int64 `v:"required#部门id不能为空" json:"branchid"`  //公司部门id

	Name   string //搜索词
	Tel    string //搜索词
	Idcard string //搜索词

	Page int32 //页数
	Size int32 //条数

	Source string `v:"required|length:1,8#来源不能为空|来源格式不合法" json:"source"` // 来源
}

type MemberInfo struct {
	Logo     string
	Name     string
	Idcard   string
	Phone    string
	Role     string //角色
	Status   int    //工作状态（1：工作，-1：离职）
	Email    string
	JoinTime *gtime.Time //加入时间
	Post     string      //职务
}

//添加公司行业
type AddIndustryReq struct {
	Name   string `v:"required#行业名称参数不能为空" json:"name"`
	Source string `v:"required|length:1,8#来源不能为空|来源格式不合法" json:"source"` // 来源
}

//修改公司行业
type ModifyIndustryReq struct {
	Id     int64 `v:"required#行业编号id不能为空" json:"id"`
	Name   string
	Status int
	Source string `v:"required|length:1,8#来源不能为空|来源格式不合法" json:"source"` // 来源
}

//删除公司行业
type DelIndustryReq struct {
	Id     int64  `v:"required#行业编号id不能为空" json:"id"`
	Source string `v:"required|length:1,8#来源不能为空|来源格式不合法" json:"source"` // 来源

}
