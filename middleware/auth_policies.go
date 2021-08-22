package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xms/pkg/authUtils"
	"xms/pkg/response"
)
//SysAdminOnly check the policy and return 403 if forbidden
//you need use it after jwt middleware!! very important
func SysAdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		g:=response.Gin{C: c}
		claim,_:= c.Get(authUtils.Claim)
		if policy,ok:= claim.(authUtils.Policy);!ok||!policy.SysAdminOnly(){
			g.Error(http.StatusForbidden,response.ERROR_FORBIDDEN,nil)
			c.Abort()
			return
		}
		c.Next()
	}
}

//CoreMemberOnly check the policy and return 403 if forbidden
//you need use it after jwt middleware!! very important
func CoreMemberOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		g:=response.Gin{C: c}
		claim,_:= c.Get(authUtils.Claim)
		if policy,ok:= claim.(authUtils.Policy);!ok||!policy.CoreMemberOnly(){
			g.Error(http.StatusForbidden,response.ERROR_FORBIDDEN,nil)
			c.Abort()
			return
		}
		c.Next()
	}
}



//SeniorMemberOnly check the policy and return 403 if forbidden
//you need use it after jwt middleware!! very important
func SeniorMemberOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		g:=response.Gin{C: c}
		claim,_:= c.Get(authUtils.Claim)
		if policy,ok:= claim.(authUtils.Policy);!ok||!policy.SeniorMemberOnly(){
			g.Error(http.StatusForbidden,response.ERROR_FORBIDDEN,nil)
			c.Abort()
			return
		}
		c.Next()
	}
}



//StaffOnly check the policy and return 403 if forbidden
//you need use it after jwt middleware!! very important
func StaffOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		g:=response.Gin{C: c}
		claim,_:= c.Get(authUtils.Claim)
		if policy,ok:= claim.(authUtils.Policy);!ok||!policy.StaffOnly(){
			g.Error(http.StatusForbidden,response.ERROR_FORBIDDEN,nil)
			c.Abort()
			return
		}
		c.Next()
	}
}

//WechatOnly check the policy and return 403 if forbidden
//you need use it after jwt middleware!! very important
func WechatOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		g:=response.Gin{C: c}
		claim,_:= c.Get(authUtils.Claim)
		if policy,ok:= claim.(authUtils.Policy);!ok||!policy.WechatOnly(){
			g.Error(http.StatusForbidden,response.ERROR_FORBIDDEN,nil)
			c.Abort()
			return
		}
		c.Next()
	}
}

