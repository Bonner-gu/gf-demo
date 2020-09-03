package service

import (
	"gf_demo_api/app/constant"
	"gf_demo_api/app/jsonapi"
	"gf_demo_api/app/model/t_admin"
	"gf_demo_api/library/captcha"

	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
	"github.com/gogf/gf/util/gvalid"
)

type AdminSrv struct{}

func NewAdminSrv() *AdminSrv {
	return &AdminSrv{}
}

//admin注册
func (a *AdminSrv) AdminSignUp(signupreq jsonapi.AdminSignupReq, r *ghttp.Request) (status bool, token string, errs error) {
	g.Log().Debugf("==> AdminSrv.AdminSignUp [request:%v]", signupreq)
	//图形码验证
	sta := captcha.CaptchaVerifyHandle(signupreq.Mkid, signupreq.Imgcode)
	if sta == false {
		status = sta
		return
	}
	//生成动态盐
	salt, _ := gmd5.Encrypt(gconv.String(gtime.Timestamp()) + signupreq.Phone + signupreq.Mkid)
	//生成密码
	password, _ := gmd5.Encrypt(salt + signupreq.Rightpw)

	//验证账号是否存在
	adminbyphone := t_admin.GetOneByPhone(signupreq.Phone)
	if adminbyphone != nil {
		status = false
		return
	}

	//DB库操作
	adminInfo := &t_admin.Entity{
		Aid:        gconv.Int64(IdmakerSrvCli.GetId("aid")),
		Name:       "dvso_" + grand.Str("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", 5),
		Phone:      signupreq.Phone,
		Logo:       "",
		Sign:       "",
		Pw:         password,
		Salt:       salt,
		RegisterIp: r.GetClientIp(),
		CreateTime: gtime.Now(),
	}

	if err := t_admin.CreateOne(adminInfo); err != nil {
		g.Log().Line().Errorf("AdminSrv.AdminSignUp :DB操作失败=====[error:%v]", err)
		errs = err
		return
	}

	//生成token
	tokenstr := new(jsonapi.Token)
	tokenstr.Ip = signupreq.Ip
	tokenstr.Mk = signupreq.Mkid
	tokenstr.Role = constant.AdminRole
	tokenstr.Source = signupreq.Source
	tokenstr.Uin = gconv.Int(adminInfo.Aid)
	jwttoken, err := AuthSrvCli.MakeToken(*tokenstr)
	if err != nil {
		g.Log().Line().Errorf(" AdminSrv.AdminSignUp:生成token失败=====[error:%v]", err.Error())
		errs = err
		return
	}
	status = sta
	token = jwttoken
	return
}

//admin登录
func (a *AdminSrv) AdminSignIn(signinreq jsonapi.AdminSigninReq) (ainfo *t_admin.Entity, status bool, token string, errs error) {
	g.Log().Debugf("==> AdminSrv.AdminSignIn [request:%v]", signinreq)
	//图形码验证
	status = captcha.CaptchaVerifyHandle(signinreq.Mkid, signinreq.Imgcode)

	if status == false {
		return
	}

	//DB库获取数据
	admin := t_admin.GetOneByPhone(signinreq.Phone)
	if admin == nil {
		ainfo = nil
		return
	}
	//密码验证
	password, _ := gmd5.Encrypt(admin.Salt + signinreq.Pw)
	if admin.Pw != password {
		g.Log().Debugf("==> AdminSrv.AdminSignIn:密码错误 [request:%v]", signinreq)
		ainfo = nil
		return
	}

	//生成token
	tokenstr := new(jsonapi.Token)
	tokenstr.Ip = signinreq.Ip
	tokenstr.Mk = signinreq.Mkid
	tokenstr.Role = constant.AdminRole
	tokenstr.Source = signinreq.Source
	tokenstr.Uin = gconv.Int(admin.Aid)
	jwttoken, err := AuthSrvCli.MakeToken(*tokenstr)
	if err != nil {
		g.Log().Line().Errorf(" AdminSrv.AdminSignIn:生成token失败=====[error:%v]", err.Error())
		errs = err
		return
	}

	token = jwttoken
	ainfo = admin
	return
}

//获取管理员信息
func (a *AdminSrv) GetAdminInfo(ainfo jsonapi.GetAdminInfoReq) (adminInfo *jsonapi.AdminInfoRet) {
	g.Log().Debugf("==> AdminSrv.GetAdminInfo [request:%v]", ainfo)
	admin := t_admin.GetOneByAid(ainfo.Uin)
	if admin == nil {
		adminInfo = nil
		return
	}
	adinfo := new(jsonapi.AdminInfoRet)
	adinfo.Logo = admin.Logo
	adinfo.Name = admin.Name
	adinfo.Phone = PhoneHide(admin.Phone)
	adinfo.Sign = admin.Sign

	adminInfo = adinfo

	return
}

//手机号*号保密处理
func PhoneHide(phone string) string {
	if len(phone) <= 10 {
		return phone
	}
	return phone[:3] + "****" + phone[len(phone)-4:]
}

//身份证*号保密处理
func IdCardHide(card string) string {
	if len(card) <= 17 {
		return card
	}
	return card[:6] + "********" + card[len(card)-4:]
}

//管理员信息修改
func (a *AdminSrv) ModifyAdminByInfo(modifyinfo jsonapi.ModifyAdminInfoReq) (status bool, errs error) {
	g.Log().Debugf("==> AdminSrv.ModifyAdminByInfo [request:%v]", modifyinfo)
	//获取待修改的admin信息
	admin := t_admin.GetOneByAid(modifyinfo.Uin)
	if admin == nil {
		status = false
		return
	}
	//信息验证/修改
	if modifyinfo.Name != "" {
		//参数校验
		if err := gvalid.Check(modifyinfo.Name, "required", "昵称参数不能为空"); err != nil {
			g.Log().Line().Errorf("AdminSrv.ModifyAdminByInfo:参数验证失败====[error:%v]", err)
			errs = err
			return
		}

		if admin.Name != modifyinfo.Name {
			admin.Name = modifyinfo.Name
		}
	}
	if modifyinfo.Logo != "" {
		//参数校验
		if err := gvalid.Check(modifyinfo.Logo, "required", "头像参数不能为空"); err != nil {
			g.Log().Line().Errorf("AdminSrv.ModifyAdminByInfo:参数验证失败====[error:%v]", err)
			errs = err
			return
		}

		if admin.Logo != modifyinfo.Logo {
			admin.Logo = modifyinfo.Logo
		}
	}
	if modifyinfo.Phone != "" {
		//参数校验
		if err := gvalid.Check(modifyinfo.Phone, "required", "账号（电话号码）参数不能为空"); err != nil {
			g.Log().Line().Errorf("AdminSrv.ModifyAdminByInfo:参数验证失败====[error:%v]", err)
			errs = err
			return
		}

		if admin.Phone != modifyinfo.Phone {
			admin.Phone = modifyinfo.Phone
		}
	}

	if modifyinfo.Sign != "" {
		//参数校验
		if err := gvalid.Check(modifyinfo.Sign, "required", "个性签名参数不能为空"); err != nil {
			g.Log().Line().Errorf("AdminSrv.ModifyAdminByInfo:参数验证失败====[error:%v]", err)
			errs = err
			return
		}
		if admin.Sign != modifyinfo.Sign {
			admin.Sign = modifyinfo.Sign
		}
	}

	err := t_admin.ModifyBase(admin, admin)
	if err != nil {
		g.Log().Line().Errorf(" AdminSrv.ModifyAdminByInfo:DB操作失败==== [error:%v]", err.Error())
		errs = err
		return
	}
	status = true
	return
}

//密码修改
func (a *AdminSrv) ModifyPassword(pinfo jsonapi.ModifyPwReq) (status bool, errs error) {
	g.Log().Debugf("==> AdminSrv.ModifyPassword [request:%v]", pinfo)
	//获取待修改的admin信息
	admin := t_admin.GetOneByAid(pinfo.Uin)
	if admin == nil {
		status = false
		return
	}
	//旧密码验证
	pwstr, _ := gmd5.Encrypt(admin.Salt + pinfo.Oldpw)
	if pwstr != admin.Pw {
		status = false
		return
	}
	//密码修改
	newpwstr, _ := gmd5.Encrypt(admin.Salt + pinfo.Rightpw)
	admin.Pw = newpwstr
	err := t_admin.ModifyBase(admin, admin)
	if err != nil {
		g.Log().Line().Errorf(" AdminSrv.ModifyPassword:DB操作失败==== [error:%v]", err.Error())
		errs = err
		return
	}

	//删除token
	err = AuthSrvCli.DelToken(*pinfo.Token)
	if err != nil {
		g.Log().Line().Errorf(" AdminSrv.ModifyPassword:Token删除失败==== [error:%v]", err.Error())
		errs = err
		return
	}

	status = true
	return
}

//密码重置
func (a *AdminSrv) ModifyNewPassword(pinfo jsonapi.ModifyNewPwReq) (sqlstatus bool, codestatus bool, errs error) {
	g.Log().Debugf("==> AdminSrv.ModifyNewPassword [request:%v]", pinfo)

	//图形码验证
	sta := captcha.CaptchaVerifyHandle(pinfo.Mkid, pinfo.Imgcode)
	if !sta {
		sqlstatus = sta
		return
	}
	//获取待修改的admin信息
	admin := t_admin.GetOneByAid(pinfo.Uin)
	if admin == nil {
		codestatus = false
		return
	}
	//密码修改
	pwstr, _ := gmd5.Encrypt(admin.Salt + pinfo.Rightpw)
	admin.Pw = pwstr
	err := t_admin.ModifyBase(admin, admin)
	if err != nil {
		g.Log().Line().Errorf(" AdminSrv.ModifyNewPassword:DB操作失败==== [error:%v]", err.Error())
		errs = err
		return
	}
	sqlstatus = true
	codestatus = true
	return
}
