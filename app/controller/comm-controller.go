package controller

import (
	"gf_demo_api/app/jsonapi"
	"gf_demo_api/library/captcha"
	"gf_demo_api/library/ecode"
	"gf_demo_api/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/mojocn/base64Captcha"
)

// 用户API管理对象
type CommController struct{}

// 获取图片验证码 返回base64
func (c *CommController) Captcha(r *ghttp.Request) {
	var params *jsonapi.CaptchaReq
	r.GetStruct(&params)
	g.Log().Debugf("======== CommController.Captcha ======== [params:%v]", params)

	// 参数校验
	if err := r.Parse(params); err != nil {
		g.Log().Line().Errorf("参数错误 [err:%v]", err.Error())
		response.JsonExit(r, ecode.ParamErr.Code(), err.Error())
	}

	// 开始生成图片验证码
	err, rsp := captcha.GenerateCaptchaHandler(&captcha.ConfigJsonBody{
		CaptchaType: "math",
		DriverMath: &base64Captcha.DriverMath{
			Height:          30,
			Width:           180,
			NoiseCount:      3,
			ShowLineOptions: 3,
		},
	})
	if err != nil {
		response.JsonExit(r, ecode.CaptchaCreateErr.Code(), "CaptchaCreateErr")
	} else {
		response.JsonExit(r, ecode.OK.Code(), "OK", rsp)
	}
}

//图形码验证
func (CommController) CaptchaCheck(r *ghttp.Request) {
	var params *jsonapi.CaptchaCheckReq
	r.GetStruct(&params)
	g.Log().Debugf("======== CommController.Captcha ======== [params:%v]", params)

	// 参数校验
	if err := r.Parse(params); err != nil {
		g.Log().Line().Errorf("参数错误 [err:%v]", err.Error())
		response.JsonExit(r, ecode.ParamErr.Code(), err.Error())
	}
	sta := captcha.CaptchaVerifyHandle(params.Mkid, params.Value)

	response.JsonExit(r, ecode.OK.Code(), "ok", sta)
}
