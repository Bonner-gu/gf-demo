package middleware

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

// 允许接口跨域请求
func SUM(r *ghttp.Request) {
	r.Middleware.Next()

	// 开始统计耗时
	startime := gconv.Int64(r.Get("request_time"))
	endtime := gtime.TimestampMilli()
	dur := endtime - startime
	g.Log().Debugf("请求完成: [middleware.sum] [start:%v] [end:%v] [dur:%v ms] [url:%v]", gtime.NewFromTimeStamp(startime).String(), gtime.NewFromTimeStamp(endtime).String(), dur, r.GetUrl())
}
