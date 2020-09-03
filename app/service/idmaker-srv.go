package service

import (
	"fmt"
	"gf_demo_api/app/constant"
	"gf_demo_api/app/model/t_idmaker"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

// id 生成器
type IdmakerSrv struct {
}

func NewIdmakerSrv() *IdmakerSrv {
	return &IdmakerSrv{}
}

func (s *IdmakerSrv) GetId(key string) int {
	g.Log().Debugf("==> IdmakerSrv.GetId [key:%v]", key)

	var lastId int

	redisKey := fmt.Sprintf(constant.IdmakerRedisKey4Id, key)
	r, _ := g.Redis().Do("INCR", redisKey)
	idFromRedis := gconv.Int(gconv.String(r))
	if idFromRedis <= constant.BaseId {
		// 从db获取值 同时加上基础id
		info := t_idmaker.GetOne(key)
		if info == nil {
			g.Log().Errorf("从db中获取 不存在该id [key:%v]", key)
			g.Redis().Do("DEL", redisKey, lastId)
			return 0
		} else {
			lastId = info.Value + 10000
			g.Redis().Do("SET", redisKey, lastId)
			g.Log().Debugf("从db中获取用户id成功 [redisKey:%v] [lastId:%v] [infoid:%v]", redisKey, lastId, info.Value)
		}
	} else {
		// 直接使用
		lastId = idFromRedis
		g.Log().Debugf("从缓存获取用户id成功 [redisKey:%v] [lastId:%v]", redisKey, lastId)
	}

	// 更新db中的数据
	go t_idmaker.UpdateId(key, lastId)

	return lastId
}
