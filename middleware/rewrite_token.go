package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"xms/pkg/authUtils"
)

// RewriteToken set the token to Authorization, and add the "Bearer " if it not exists
func RewriteToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader:=c.Request.Header.Get("Authorization")
		if !(len(authHeader)>0&&strings.HasPrefix(authHeader,"Bearer ")){
			//put the cookie to header if exist
			cookie,err:=c.Cookie(authUtils.XMS_AUTH_BEARER)
			if err==nil{
				//dont't forget to "add Bearer "
				if strings.Contains(cookie,"Bearer ") {
					c.Request.Header.Set("Authorization",cookie)
				} else {
					c.Request.Header.Set("Authorization", "Bearer "+cookie)
				}
			}
		}
		c.Next()
	}
}
