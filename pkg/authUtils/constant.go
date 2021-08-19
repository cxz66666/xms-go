package authUtils

import (
	jwt "github.com/dgrijalva/jwt-go"
)
// EvaClaimTypes is used for jwt token, used by oreo and other
type EvaClaimTypes struct {
	Name string `json:"name"`
	Role uint8 `json:"role"`
	Department uint8 `json:"department"`
	UserId int `json:"user_id"`
	Type string `json:"type"`

	jwt.StandardClaims
}


// WechatClaimTypes is alos used for jwt token, but used by wechat mini-program
type WechatClaimTypes struct {
	OpenId string `json:"open_id"`
	UnionId string `json:"union_id"`
	SessionKey string `json:"session_key"`

	jwt.StandardClaims
}