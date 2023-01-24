package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"mini_tiktok/pkg/consts"
	"time"
)

type Claims struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

func CreateToken(userId int64) (string, error) {
	expireTime := time.Now().Add(24 * 7 * time.Hour) //过期时间为7天
	nowTime := time.Now()                            //当前时间
	claims := Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间戳
			IssuedAt:  nowTime.Unix(),    //当前时间戳
			Issuer:    "mini-tiktok",     //颁发者签名
		},
	}
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenStruct.SignedString([]byte(consts.SecretKey))
}

func CheckToken(token string) (*Claims, bool) {
	tokenObj, _ := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return consts.SecretKey, nil
	})
	fmt.Println(tokenObj)
	if key, _ := tokenObj.Claims.(*Claims); tokenObj.Valid {
		return key, true
	} else {
		return nil, false
	}
}
