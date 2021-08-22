package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xms/pkg/logging"
	"xms/pkg/response"
	"xms/service/ipblocker"
)

// IPBlock is a middleware for login api to block the fail login after fail to 5 times
func IPBlock() gin.HandlerFunc {
	return func(c *gin.Context) {
		g:=response.Gin{C: c}
		ip:=c.ClientIP()
		if result:=ipblocker.IsLoginable(ip);!result{
			g.Error(http.StatusOK,response.ERROR_IP_BLOCK,nil)
			c.Abort()
			return
		}
		c.Next()

		status,exist:=c.Get("LOGINSTATUS")
		if statusStr,ok:=status.(string);!exist||!ok||statusStr!="success"{
			logging.InfoF("ip %s try to login and fail",ip)
			ipblocker.Fail(ip)
		} else {
			ipblocker.Success(ip)
		}
	}
}