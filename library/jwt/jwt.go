package jwt

import (
	"gf_demo_api/app/constant"
	"gf_demo_api/app/jsonapi"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtSecret = constant.AuthJwtSecret
	myissuer  = constant.AuthJwtMyissuer
	ExpiresAt = constant.AuthJwtExpiresAt
)

type Claims struct {
	jwt.StandardClaims
	JwtSession jsonapi.Token
}

// 生成token
func CreateToken(JwtSession jsonapi.Token) (tokenString string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(ExpiresAt)

	claims := Claims{
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    myissuer,
		},
		JwtSession,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtSecret)
	return
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
