package constant

import "time"

const (
	AuthRedisKey4Skey        = "gf_demo_api:auth:session:uin:%d:source:%s:role:%s"
	AuthRedisKey4SkeyTimeout = 86400

	AuthJwtMyissuer  = "test.com"
	AuthJwtExpiresAt = 30 * 24 * time.Hour
)

var AuthJwtSecret = []byte("Yb279t=#iR5B$+lnJ+xtAIxtAX$6Tm-gf_demo_api-C[")
