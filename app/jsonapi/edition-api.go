package jsonapi

type EditionReq struct {
	Uin            int64
	Title          string `v:"required#任务名不能为空"`
	VersionNum     string `v:"required#当前更新的版本号不能为空"` //当前更新的版本号
	IOSAddress     string //IOS的下载地址
	AndroidAddress string //Android下载地址
	Mark           string `v:"required#版本备注信息不能为空"` //版本备注信息
}

type GetEditionReq struct {
	Page       int32
	Size       int32
	VersionNum string
	Mark       string
	Title      string
}
