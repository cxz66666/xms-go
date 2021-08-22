package authUtils

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"xms/models"
	"xms/pkg/setting"
)



const (
	// XMS_AUTH_BEARER is the cookie name to store token
	XMS_AUTH_BEARER="XMS_AUTH_BEARER"
	Claim="EVACLAIM"

	EVACLAIM =0
	WECHATCLAIM =1
)

// Policy is the interface for EvaClaimTypes and WechatClaimTypes, it will be stored in context and used for middleware to auth
type Policy interface {
	SysAdminOnly() bool
	CoreMemberOnly() bool
	SeniorMemberOnly() bool
	StaffOnly()  bool
	WechatOnly() bool
	CheckExpired() bool
	GetType() int
}



// EvaClaimTypes is used for jwt token, used by oreo and other
type EvaClaimTypes struct {
	Name string `json:"name"`
	Role models.Role `json:"role"`
	Department models.Department `json:"department"`
	UserId int `json:"user_id"`
	Type string `json:"type"`

	jwt.StandardClaims
}
func (e *EvaClaimTypes) SysAdminOnly() bool {
	return e.Role==models.Sysadmin
}
func (e *EvaClaimTypes) CoreMemberOnly() bool {
	return e.Role==models.Sysadmin||e.Role==models.President||e.Role==models.VicePresident||e.Role==models.Minister
}

func (e *EvaClaimTypes) SeniorMemberOnly() bool {
	return e.Role==models.Sysadmin||e.Role==models.President||e.Role==models.VicePresident ||
		e.Role== models.Minister||e.Role==models.TechAdviser||e.Role==models.ViceMinister
}

func (e *EvaClaimTypes) StaffOnly() bool {
	return e.Type=="staff"
}

func (e *EvaClaimTypes) WechatOnly() bool {
	return e.Type=="wechat"
}

func (e *EvaClaimTypes) CheckExpired() bool {
	return time.Now().Unix()<e.ExpiresAt
}
func (e *EvaClaimTypes) GetType() int {
	return EVACLAIM
}


//GetEvaClaimFromUser convert models.User to EvaClaimTypes (already init the jwt.StandardClaims)
func GetEvaClaimFromUser(user models.User) *EvaClaimTypes  {
	nowTime:=time.Now()
	expireTime:=nowTime.Add(setting.ServerSetting.JwtExpireTime)
	return &EvaClaimTypes{
		Name: user.Name,
		Role: user.Role,
		Department: user.Department,
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

func (w *WechatClaimTypes) SysAdminOnly() bool {
	return false
}

func (w *WechatClaimTypes) CoreMemberOnly() bool {
	return false
}

func (w *WechatClaimTypes) SeniorMemberOnly() bool {
	return false
}

func (w *WechatClaimTypes) StaffOnly() bool {
	return  w.Type=="staff"
}

func (w *WechatClaimTypes) WechatOnly() bool {
	return w.Type=="wechat"
}

func (w *WechatClaimTypes) CheckExpired() bool {
	return time.Now().Unix()<w.ExpiresAt
}
func (w *WechatClaimTypes) GetType() int {
	return WECHATCLAIM
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