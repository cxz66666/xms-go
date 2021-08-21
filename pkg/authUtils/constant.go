package authUtils

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"xms/models"
	"xms/pkg/setting"
)


// XMS_AUTH_BEARER is the cookie name to store token
const XMS_AUTH_BEARER="XMS_AUTH_BEARER"


// EvaClaimTypes is used for jwt token, used by oreo and other
type EvaClaimTypes struct {
	Name string `json:"name"`
	Role uint8 `json:"role"`
	Department uint8 `json:"department"`
	UserId int `json:"user_id"`
	Type string `json:"type"`

	jwt.StandardClaims
}

//GetEvaClaimFromUser convert models.User to EvaClaimTypes (already init the jwt.StandardClaims)
func GetEvaClaimFromUser(user models.User) *EvaClaimTypes  {
	nowTime:=time.Now()
	expireTime:=nowTime.Add(setting.ServerSetting.JwtExpireTime)
	return &EvaClaimTypes{
		Name: user.Name,
		Role: uint8(user.Role),
		Department: uint8(user.Department),
		UserId: user.ID,
		Type: "staff",
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: setting.SecretSetting.JwtIssuer,
		},
	}
}

// WechatClaimTypes is alos used for jwt token, but used by wechat mini-program
type WechatClaimTypes struct {
	OpenId string `json:"open_id"`
	UnionId string `json:"union_id"`
	SessionKey string `json:"session_key"`
	Type string `json:"type"`

	jwt.StandardClaims
}

//GetWechatClaimFromString convert string to WechatClaimTypes (already init the jwt.StandardClaims)
func GetWechatClaimFromString(openId,sessionKey, unionId string) *WechatClaimTypes  {
	nowTime:=time.Now()
	expireTime:=nowTime.Add(setting.ServerSetting.JwtExpireTime)

	return &WechatClaimTypes{
		OpenId: openId,
		UnionId: unionId,
		SessionKey: sessionKey,
		Type: "wechat",
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: setting.SecretSetting.JwtIssuer,
		},
	}
}