package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"xms/pkg/authUtils"
	"xms/pkg/response"
)
//Jwt is the Authentication middleware, it will write the claim to context if it's valid, and it will return error and abort if not valid
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		g:=response.Gin{C: c}
		authHeader:=c.Request.Header.Get("Authorization")
		code:=-1
		var policy authUtils.Policy
		var err error
		if len(authHeader)==0 {
			code=response.ERROR_NOT_LOGIN
		} else {
			parts := strings.SplitN(authHeader, " ", 2)
			if !(len(parts) == 2 && parts[0] == "Bearer"){
				code=response.ERROR_AUTH_NO_VALID_HEADER   //token格式不对
			} else {
				policy,err=authUtils.ParseToken(parts[1])
				//token不正确
				if err!=nil{
					code=response.ERROR_TOKEN_NOT_VAILD
				} else if !policy.CheckExpired(){ //token过期
					code=response.ERROR_TOKEN_EXPIRED
				}
			}
		}

		if code>0 {
			g.Error(http.StatusOK,code,nil)
			c.Abort()
			return
		}
		//set the policy to context
		c.Set(authUtils.Claim,policy)
		c.Next()
	}
}