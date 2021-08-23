package stats_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xms/models"
	"xms/pkg/cache"
	"xms/pkg/response"
)

func GetTicketCount(c *gin.Context)  {
	g:=response.Gin{C: c}
	count:=cache.GetOrCreate(cache.GetKey(cache.Stats,int(cache.TicketCount)), func() interface{} {
		num,_:=models.GetTicketCount()
		return num
	}).(int)
	g.Success(http.StatusOK,count)
	return
}