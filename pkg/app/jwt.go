package app

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/diy0663/goblog-service/global"
	"github.com/diy0663/goblog-service/pkg/util"
)

type Claims struct {
	Appkey    string `json:"app_key"`
	Appsecret string `json:"app_secret"`
	// 嵌入结构体,它是 jwt-go 库中预定义的
	jwt.StandardClaims
}

// todo 这里为何要用 []byte ???
func GetJWTSecret() []byte {
	return []byte(global.JwtSetting.Secret)
}

// 生成token
func GenerateToken(appKey, appSecret string) (string, error) {
	// 现在的时间
	nowTime := time.Now()
	// 未来过期的具体时间
	expireTime := nowTime.Add(global.JwtSetting.Expire)
	claims := Claims{
		Appkey:    util.EncodeMD5(appKey),
		Appsecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JwtSetting.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err

}

// 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err

}
