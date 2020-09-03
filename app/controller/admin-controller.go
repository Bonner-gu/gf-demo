package controller

import (
	"gf_demo_api/app/constant"
	"gf_demo_api/app/jsonapi"
	"gf_demo_api/app/service"
	"gf_demo_api/library/ecode"
	"gf_demo_api/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

//api管理对象
type AdminController struct{}

//管理员注册
func (a *AdminController) AdminSignup(r *ghttp.Request) {
	var params *jsonapi.AdminSignupReq
	r.GetStruct(&params)
	g.Log().Debugf("======== AdminController.AdminSignup ======== [params:%v]", params)

	// 参数校验
	if err := r.Parse(params); err != nil {
		g.Log().Line().Errorf("AdminController.AdminSignup:参数错误====[err:%v]", err.Error())
		response.JsonExit(r, ecode.ParamErr.Code(), err.Error())
	}

	//注册
	status, token, err := service.AdminSrvCli.AdminSignUp(*params, r)
	if err != nil {
		g.Log().Line().Errorf("AdminController.AdminSignup:DB操作失败====[err:%v]", err.Error())
		response.JsonExit(r, ecode.MysqlManageErr.Code(), "MysqlManageErr")
	}
	if !status {
		response.JsonExit(r, ecode.CaptchaCheckErr.Code(), "CaptchaCheckErr")
	}
	r.Response.Header().Set(constant.BaseCorsHeader, token)
	response.JsonExit(r, ecode.OK.Code(), "ok")
}

//管理员登录
func (a *AdminController) AdminSignin(r *ghttp.Request) {
	var params *jsonapi.AdminSigninReq
	r.GetStruct(&params)
	g.Log().Debugf("======== AdminController.AdminSignin ======== [params:%v]", params)

	// 参数校验
	if err := r.Parse(params); err != nil {
		g.Log().Line().Errorf("AdminController.AdminSignin:参数错误====[err:%v]", err.Error())
		response.JsonExit(r, ecode.ParamErr.Code(), err.Error())
	}

	//登录
	admin, status, token, err := service.AdminSrvCli.AdminSignIn(*params)
	g.Log().Debugf("====[%v][%v][%v]====", admin, status, token)
	if err != nil {
		g.Log().Line().Errorf("AdminController.AdminSignin:DB获取数据失败====[err:%v]", err.Error())
		response.JsonExit(r, ecode.PasswordInvalid.Code(), "PasswordInvalid")
	}
	if status == false {
		response.JsonExit(r, ecode.CaptchaCheckErr.Code(), "CaptchaCheckErr")
	}
	if admin == nil {
		response.JsonExit(r, ecode.PasswordInvalid.Code(), "PasswordInvalid")
	}

	r.Response.Header().Set(constant.BaseCorsHeader, token)
	response.JsonExit(r, ecode.OK.Code(), "ok")
}

//获取管理员信息
func (a *AdminController) GetAdminInfo(r *ghttp.Request) {
	var params *jsonapi.GetAdminInfoReq
	r.GetStruct(&params)
	g.Log().Debugf("======== AdminController.GetAdminInfo ======== [params:%v]", params)

	// 参数校验
	if err := r.Parse(params); err != nil {
		g.Log().Line().Errorf("AdminController.GetAdminInfo:参数错误====[err:%v]", err.Error())
		response.JsonExit(r, ecode.ParamErr.Code(), err.Error())
	}

	//DB获取数据
	ainfo := service.AdminSrvCli.GetAdminInfo(*params)
	if ainfo == nil {
		response.JsonExit(r, ecode.AdminNoExist.Code(), "AdminNoExist")
	}

	response.JsonExit(r, ecode.OK.Code(), "ok", ainfo)
}

//修改管理员信息
func (a *AdminController) ModifyAdminInfo(r *ghttp.Request) {
	var params *jsonapi.ModifyAdminInfoReq
	r.GetStruct(&params)
	g.Log().Debugf("======== AdminController.ModifyAdminInfo ======== [params:%v]", params)

	// 参数校验
	if err := r.Parse(params); err != nil {
		g.Log().Line().Errorf("AdminController.ModifyAdminInfo:参数错误====[err:%v]", err.Error())
		response.JsonExit(r, ecode.ParamErr.Code(), err.Error())
	}

	status, err := service.AdminSrvCli.ModifyAdminByInfo(*params)
	if err != nil {
		g.Log().Line().Errorf("AdminController.ModifyAdminInfo:DB库操作失败====[err:%v]", err.Error())
		response.JsonExit(r, ecode.MysqlManageErr.Code(), "MysqlManageErr")
	}
	if !status {
		response.JsonExit(r, ecode.MysqlManageErr.Code(), "MysqlManageErr")
	}

	response.JsonExit(r, ecode.OK.Code(), "ok")
}

//修改管理员密码
func (a *AdminController) ModifyAdminPw(r *ghttp.Request) {
	var params *jsonapi.ModifyPwReq
	r.GetStruct(&params)
	g.Log().Debugf("======== AdminController.ModifyAdminPw ======== [params:%v]", params)

	// 参数校验
	if err := r.Parse(params); err != nil {
		g.Log().Line().Errorf("AdminController.ModifyAdminPw:参数错误====[err:%v]", err.Error())
		response.JsonExit(r, ecode.ParamErr.Code(), err.Error())
	}

	status, err := service.AdminSrvCli.ModifyPassword(*params)
	if err != nil {
		g.Log().Line().Errorf("AdminController.ModifyAdminPw:DB库操作失败====[err:%v]", err.Error())
		response.JsonExit(r, ecode.MysqlManageErr.Code(), "MysqlManageErr")
	}
	if !status {
		response.JsonExit(r, ecode.MysqlManageErr.Code(), "MysqlManageErr")
	}

	response.JsonExit(r, ecode.OK.Code(), "ok")
}

//修改管理员密码
func (a *AdminController) ModifyAdminNewPw(r *ghttp.Request) {
	var params *jsonapi.ModifyNewPwReq
	r.GetStruct(&params)
	g.Log().Debugf("======== AdminController.ModifyAdminNewPw ======== [params:%v]", params)

	// 参数校验
	if err := r.Parse(params); err != nil {
		g.Log().Line().Errorf("AdminController.ModifyAdminNewPw:参数错误====[err:%v]", err.Error())
		response.JsonExit(r, ecode.ParamErr.Code(), err.Error())
	}

	status, status1, err := service.AdminSrvCli.ModifyNewPassword(*params)
	if err != nil {
		g.Log().Line().Errorf("AdminController.ModifyAdminNewPw:DB库操作失败====[err:%v]", err.Error())
		response.JsonExit(r, ecode.MysqlManageErr.Code(), "MysqlManageErr")
	}
	if !status1 {
		response.JsonExit(r, ecode.CaptchaCheckErr.Code(), "CaptchaCheckErr")
	}
	if !status {
		response.JsonExit(r, ecode.MysqlManageErr.Code(), "MysqlManageErr")
	}

	response.JsonExit(r, ecode.OK.Code(), "ok")
}
