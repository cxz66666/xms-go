package auth_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
	"xms/models"
	"xms/pkg/authUtils"
	"xms/pkg/cache"
	"xms/pkg/response"
	"xms/pkg/setting"
	"xms/pkg/utils"
)

func Login(c *gin.Context)  {
	g:=response.Gin{C: c}
	var loginReq LoginReq
	if err:=c.ShouldBind(&loginReq);err!=nil{
		g.Error(http.StatusOK,response.ERROR_AUTH_PARAM_FAIL,err)
	}

	loginStatus:="fail"
	defer c.Set("LOGINSTATUS",loginStatus)

	if loginReq.Username==setting.AdminSetting.Username{
		if loginReq.Password==setting.AdminSetting.Password{
			//important
			loginStatus="success"
			adminToken,err:=authUtils.GetAdminToken()
			if err!=nil{
				g.Error(http.StatusOK,response.ERROR_TOKEN_GENERATE_FAIL,err)
				return
			}
			c.SetCookie(authUtils.XMS_AUTH_BEARER,"Bearer "+adminToken,int(time.Hour.Seconds()),"/","",false,true)
			g.Success(http.StatusOK,NewLoginRes(models.User{
				Name: "系统维护管理员",
				StudentId: 10086,
				Department: models.DN,
				Role: models.Sysadmin,
			},adminToken))
			return
		} else {
			g.Error(http.StatusOK,response.ERROR_ADMIN_INVALID_PASSWORD,nil)
			return
		}
	}

	secret:=utils.Password2Secret(loginReq.Password,setting.SecretSetting.SaltA,setting.SecretSetting.SaltB)
	stuId,err:=strconv.Atoi(loginReq.Username)
	if err!=nil{
		g.Error(http.StatusOK,response.ERROR_LOGIN_USERNAME,err)
		return
	}

	uid:=cache.GetOrCreate(cache.GetKey(cache.StuIdToUId,stuId), func() interface{} {
		return models.StuId2Id(stuId)
	}).(int)

	user:=cache.GetOrCreate(cache.GetKey(cache.UserInfo,uid), func() interface{} {

		user,err:= models.GetUserById(uid)
		if err!=nil	{
			return &models.User{
				ID: -1,
			}
		}
		return user
	}).(*models.User)

	if user.ID<0{
		g.Error(http.StatusOK,response.ERROR_AUTH_UNKNOWN_ERROR,nil)
		return
	}
	if user.Secret!=secret {
		g.Error(http.StatusOK,response.ERROR_AUTH_INVALID_PASSWORD,nil)
		return
	}

	loginStatus="success"
	token,err:=authUtils.GetStaffToken(*user)
	if err!=nil {
		g.Error(http.StatusOK,response.ERROR_TOKEN_GENERATE_FAIL,err)
		return
	}
	c.SetCookie(authUtils.XMS_AUTH_BEARER,"Bearer "+token,int(time.Hour.Seconds()),"/","",false,true)
	g.Success(http.StatusOK,NewLoginRes(*user,token))
	return

}

func SetToken(c *gin.Context)  {
	g:=response.Gin{C: c}

	authHeader:=c.GetHeader("Authorization")
	if len(authHeader)==0||!strings.HasPrefix(authHeader,"Bearer ") {
		g.Error(http.StatusOK,response.ERROR_AUTH_NO_VALID_HEADER,nil)
		return
	}
	c.SetCookie(authUtils.XMS_AUTH_BEARER,authHeader,int(time.Hour.Seconds()),"/","",false,true)

	g.Success(http.StatusOK,nil)
}


//Logout check whether login use middleware before this function
func Logout(c *gin.Context) {
	g:=response.Gin{C: c}
	c.SetCookie(authUtils.XMS_AUTH_BEARER,"",-1,"/","",false,true)
	g.Success(http.StatusOK,nil)
}

//Ping just test the connection
func Ping(c *gin.Context)  {
	g:=response.Gin{C: c}
	g.Success(http.StatusOK,"pa")
}