package authUtils

import (
	"github.com/dgrijalva/jwt-go"
	"xms/models"
	"xms/pkg/setting"
	"xms/pkg/utils"
)

// GetStaffToken generate a token for a simple user, and return the token by string and error (if it exists)
func GetStaffToken(user models.User) (string,error) {

	model:=GetEvaClaimFromUser(user)
	tokenClaims:=jwt.NewWithClaims(jwt.SigningMethodHS256,model)

	token,err:=tokenClaims.SignedString([]byte(setting.SecretSetting.JwtKey))
	return token,err
}

// GetAdminToken generate a token for sysAdmin user, and return the token by string and error (if it exists)
func GetAdminToken() (string,error) {
	user:=models.NewAdminUser()
	return GetStaffToken(*user)
}

// GetWechatToken generate a token for wechat mini program to use, and return the token by string and error
func GetWechatToken(openId, sessionKey, unionId string) (string,error) {
	//encrypt the sessionKey
	encryptedKey,err:=utils.AesCBCEncrypt([]byte(sessionKey),[]byte(setting.SecretSetting.AesKey),[]byte(setting.SecretSetting.AesIv))

	model:=GetWechatClaimFromString(openId,string(encryptedKey),unionId)
	tokenClaims:=jwt.NewWithClaims(jwt.SigningMethodHS256,model)

	token,err:=tokenClaims.SignedString(setting.SecretSetting.JwtKey)
	return token,err
}