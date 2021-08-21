package auth_controller

import (
	"strconv"
	"xms/models"
)

// LoginReq is the http post request data
type LoginReq struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Appid string `json:"appid" form:"appid"`
	Appsecret string `json:"appsecret" form:"appsecret"`
}


type LoginRes struct {
	Stuid string `json:"stuid"`
	Name string `json:"name"`
	Group string `json:"group"`
	Token string `json:"token"`
	Authority string `json:"authority"`
}

func NewLoginRes(user models.User,_token string) *LoginRes  {
	var auth string
	switch user.Role {
		case models.Sysadmin:
			auth="admin"
		case models.Staff:
			auth="user"
		case models.Retired:
			auth="user"
		default:
			auth="senior"
	}
	return &LoginRes{
		Stuid: strconv.Itoa(user.StudentId),
		Name: user.Name,
		Group: user.Department.ToDisplayName(),
		Token: _token,
		Authority:auth,
	}
}