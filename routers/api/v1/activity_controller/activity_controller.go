package activity_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xms/models"
	"xms/pkg/cache"
	"xms/pkg/helper"
	"xms/pkg/response"
	"xms/service/activity_service"
)

func GetLatestActivities(c *gin.Context)  {

	g:=response.Gin{C: c,}
	notes,err:=activity_service.GetCurrentNotes()
	if err!=nil{
		g.Error(http.StatusOK,response.ERROR_DATABASE_QUERY,err)
		return
	}

	activityList:=make([]ActivityPiece,0,len(notes))

	for _,item:=range notes {
		activityList = append(activityList, ActivityPiece{
			Id: item.ID,
			UpdateAt: item.CreatedTime,
			UserId: item.OperatorId,
			NoteString: helper.ToHumanReadableString(item),
		})
	}

	for i,item:=range activityList {
		user:=cache.GetOrCreate(cache.GetKey(cache.UserInfo,item.UserId), func() interface{} {
			user,_:= models.GetUserById(item.UserId)
			return user
		}).(*models.User)
		if len(user.Name)>0 {
			activityList[i].UserName=user.Name
			activityList[i].Content=fmt.Sprintf("%s %s",user.Name,item.NoteString)
		} else {
			activityList[i].UserName="未知用户"
			activityList[i].Content=fmt.Sprintf("%s %s","未知用户",item.NoteString)
		}
	}

	g.Success(http.StatusOK,ActivityRes{
		PageId: 1,
		Count: len(activityList),
		Content: activityList,
	})
}