package middleware

import (
	"gf_demo_api/app/service"
	"gf_demo_api/library/ecode"
	"gf_demo_api/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

// 鉴权中间件，只有登录成功之后才能通过
func Auth(r *ghttp.Request) {
	// 获取token
	strTokenFromHeader := r.Header.Get("X-Auth-Token")

	// 基础参数校验
	if e := gvalid.Check(strTokenFromHeader, "length:32,512", "token err"); e != nil {
		g.Log().Line().Errorf("token格式不合法 [middleware.auth] [err:%v] [strtoken:%v]", e.Error(), strTokenFromHeader)
		response.JsonExit(r, ecode.Unauthorized.Code(), "Unauthorized")
	}

	// check
	token, err := service.AuthSrvCli.CheckToken(strTokenFromHeader)
	if err != nil {
		g.Log().Line().Errorf("token校验失败 [middleware.auth] [err:%v] [strtoken:%v] [token:%v]", err.Error(), strTokenFromHeader, token)
		response.JsonExit(r, ecode.Unauthorized.Code(), "Unauthorized")
	} else {
		//r.SetParam("str_token", strTokenFromHeader)
		r.SetParam("token", token)
		r.SetParam("uin", token.Uin)
		r.Middleware.Next()
	}

	r.Middleware.Next()
}
