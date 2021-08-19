package authUtils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"xms/pkg/setting"
)

// GetStaffToken generate a token for a simple user, and return the token by string and error (if it exists)
func GetStaffToken(model EvaClaimTypes) (string,error) {
	nowTime:=time.Now()
	expireTime:=nowTime.Add(setting.ServerSetting.JwtExpireTime)
	model.StandardClaims=jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		Issuer: setting.SecretSetting.JwtIssuer,
	}
	tokenClaims:=jwt.NewWithClaims(jwt.SigningMethodHS256,model)

	token,err:=tokenClaims.SignedString(setting.SecretSetting.JwtKey)
	return token,err
}

func GetAdminToken(model EvaClaimTypes)  {
	//TODO I need continue after finish model
}