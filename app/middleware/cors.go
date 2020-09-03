package middleware

import (
	"gf_demo_api/app/constant"

	"github.com/gogf/gf/net/ghttp"
)

// 允许接口跨域请求
func CORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()

	// 配置允许跨域的域名 默认全部
	//corsOptions.AllowDomain = []string{"test.com", "api.test.com"}
	corsOptions.AllowHeaders = "Origin,Content-Type,Accept,User-Agent,Cookie,Authorization,X-Auth-Token,X-Requested-With"
	// 配置header
	corsOptions.ExposeHeaders = constant.BaseCorsHeader

	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}
