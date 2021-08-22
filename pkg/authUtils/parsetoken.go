package authUtils

import (
	"github.com/dgrijalva/jwt-go"
	"xms/pkg/setting"
)

//ParseToken will try to parse the jwt token (maybe EvaClaimTypes or WechatClaimTypes)
func ParseToken(token string) (Policy, error) {
	tokenClaims,err:=jwt.ParseWithClaims(token,&EvaClaimTypes{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(setting.SecretSetting.JwtKey),nil
	})
	//try to parse EvaClaimTypes
	if tokenClaims!=nil {
		if claims,ok:=tokenClaims.Claims.(*EvaClaimTypes);ok&&tokenClaims.Valid {
			return claims,nil
		}
	}
	tokenClaims,err=jwt.ParseWithClaims(token,&WechatClaimTypes{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(setting.SecretSetting.JwtKey),nil
	})
	//try to parse WechatClaimTypes
	if tokenClaims!=nil {
		if claims,ok:=tokenClaims.Claims.(*WechatClaimTypes);ok&&tokenClaims.Valid {
			return claims,nil
		}
	}
	return nil,err
}