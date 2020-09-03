package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/guid"
)

// 设置默认参数
func Base(r *ghttp.Request) {
	// 设置请求开始时间 1ms
	r.SetParam("request_time", gtime.TimestampMilli())

	// 获取机器码
	var mk string
	mkFromHeader := r.Header.Get("mk")
	if mkFromHeader == "" {
		mk = guid.S()
	} else {
		mk = mkFromHeader
	}

	r.SetParam("mk", mk)

	// 设置ip
	r.SetParam("ip", r.GetClientIp())

	// 设置客户端参数
	var source string
	sourceFromParam := r.Get("source")
	switch gconv.String(sourceFromParam) {
	case "PC":
		source = "PC"
	case "H5":
		source = "H5"
	case "ANDROID":
		source = "ANDROID"
	case "IOS":
		source = "IOS"
	case "WX":
		source = "WX"
	default:
		source = "PC"
	}
	r.SetParam("source", source)

	// 设置角色
	var Role string
	RoleFromParam := r.Get("role")
	switch gconv.String(RoleFromParam) {
	case "ADMIN":
		Role = "ADMIN"
	default:
		Role = "USER"
	}
	r.SetParam("role", Role)

	r.Middleware.Next()
}
