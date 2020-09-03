package router

import (
	"gf_demo_api/app/controller"
	"gf_demo_api/app/middleware"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()

	// 某些浏览器直接请求favicon.ico文件，特别是产生404时
	s.SetRewrite("/favicon.ico", "/resource/image/favicon.ico")

	// api
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Base, middleware.CORS, middleware.SUM)

		// 公共接口
		commController := new(controller.CommController)
		group.GET("/comm/captcha", commController, "Captcha")
		group.GET("/comm/captchacheck", commController, "CaptchaCheck")

		//admin登录、注册接口
		adminController := new(controller.AdminController)

		group.POST("/admin/signup", adminController, "AdminSignup")
		group.POST("/admin/signin", adminController, "AdminSignin")
		group.POST("/admin/newpw", adminController, "ModifyAdminNewPw")

		//admin信息操作
		group.Middleware(middleware.Base, middleware.CORS, middleware.SUM, middleware.Auth)
		group.GET("/admin/profile", adminController, "GetAdminInfo")
		group.POST("/admin/modifyinfo", adminController, "ModifyAdminInfo")
		group.POST("/admin/modifypw", adminController, "ModifyAdminPw")

	})
}
