package service

import (
	"fmt"
	"gf_demo_api/app/constant"
	"gf_demo_api/app/jsonapi"
	"gf_demo_api/library/ecode"
	"gf_demo_api/library/jwt"

	"github.com/gogf/gf/util/guid"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

type AuthSrv struct{}

func NewAuthSrv() *AuthSrv {
	return &AuthSrv{}
}

// 生成token
func (s *AuthSrv) MakeToken(token jsonapi.Token) (strToken string, err error) {
	g.Log().Debugf("==> AuthSrv.MakeToken [token:%v]", token)

	token.Skey = guid.S()
	g.Redis().Do("SET", s.GetRedisKey(token), token.Skey, "EX", constant.AuthRedisKey4SkeyTimeout)
	return jwt.CreateToken(token)
}

// 校验token
func (s *AuthSrv) CheckToken(strToken string) (token jsonapi.Token, err error) {
	g.Log().Debugf("==> AuthSrv.CheckToken [strToken:%v]", strToken)

	// 参数校验
	if e := gvalid.Check(strToken, "required|length:64,1024", "必填|参数格式不合法"); e != nil {
		return token, e
	}

	// parse
	rep, err := jwt.ParseToken(strToken)
	if err != nil {
		g.Log().Line().Errorf("jwt字符串解析失败 [strtoken:%v] [err:%v]", strToken, err.Error())
		err = err
		return
	} else {
		token = rep.JwtSession
		g.Log().Debugf("jwt字符串解析成功 [token:%v]", token)
	}

	// from redis
	redisKey := s.GetRedisKey(token)
	var skeyFromRedis string
	if rsp, err := g.Redis().Do("GET", redisKey); err != nil {
		g.Log().Line().Errorf("redis操作失败 [err:%v]", err.Error())
		return token, err
	} else {
		skeyFromRedis = gconv.String(rsp)
		g.Log().Debugf("缓存中获取skey成功 [redisKey:%v] [skey:%v]", redisKey, skeyFromRedis)
	}

	// check
	if gstr.ToUpper(skeyFromRedis) != gstr.ToUpper(token.Skey) {
		g.Log().Line().Errorf("skey校验失败 [redisKey:%v] [skey:%v] [token:%v]", redisKey, skeyFromRedis, token)
		return token, ecode.AccessDenied
	} else {
		g.Log().Debugf("skey校验成功 [redisKey:%v] [skey:%v] [token:%v]", redisKey, skeyFromRedis, token)

		// 开始续期
		go func() {
			g.Redis().Do("SET", redisKey, token.Skey, "EX", constant.AuthRedisKey4SkeyTimeout)
			g.Log().Debugf("session 续期完成 [redisKey:%v] [skey:%v] [Redis4SessionTimeOut:%v]", redisKey, token.Skey, constant.AuthRedisKey4SkeyTimeout)
		}()

		return token, nil
	}
}

// 删除token
func (s *AuthSrv) DelToken(token jsonapi.Token) error {
	g.Log().Debugf("==> AuthSrv.DelToken [token:%v]", token)

	_, err := g.Redis().Do("DEL", s.GetRedisKey(token))
	return err
}

// 获取rediskey auth:session:uin:source:role
func (s *AuthSrv) GetRedisKey(token jsonapi.Token) string {
	return gstr.ToLower(fmt.Sprintf(constant.AuthRedisKey4Skey, token.Uin, token.Source, token.Role))
}
