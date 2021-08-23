package user_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"xms/models"
	"xms/pkg/cache"
	"xms/pkg/response"
)

func GetMyInfo(c *gin.Context) {
	g:=response.Gin{C: c}
	userId:=g.GetUserId()
	user:=cache.GetOrCreate(cache.GetKey(cache.UserInfo,userId), func() interface{} {
		user,err:= models.GetUserById(userId)
		if err!=nil	{
			return &models.User{
				ID: -1,
			}
		}
		return user
	}).(*models.User)
	if user.ID<=0 {
		g.Error(http.StatusOK,response.ERROR_USER_NOT_FOUND,nil)
		return
	}

	url:="https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"
	if len(user.AvatarURL)>0{
		url=user.AvatarURL
	}
	g.Success(http.StatusOK,UserInfoRes{
		Name: user.Name,
		Userid: userId,
		Email: "",
		Title: user.Role,
		Group: user.Department,
		Avatar: url,
	})
	return
}

func SetUserAvatar(c *gin.Context)  {
	g:=response.Gin{C: c}
	idStr:=c.Param("id")
	id,ok:=strconv.Atoi(idStr)
	if !(ok==nil&&id>0){
		g.Error(http.StatusOK,response.ERROR_PARAM_NOT_VAILD,nil)
		return
	}
	user:=cache.GetOrCreate(cache.GetKey(cache.UserInfo,id), func() interface{} {
		user,err:= models.GetUserById(id)
		if err!=nil	{
			return &models.User{
				ID: -1,
			}
		}
		return user
	}).(*models.User)

	if user.ID<=0 {
		g.Error(http.StatusOK,response.ERROR_USER_NOT_FOUND,nil)
		return
	}

	user.AvatarURL = "https://1299271970796699.cn-hangzhou.fc.aliyuncs.com/2016-08-15/proxy/oss_upload/download/?stuId="+ user.GetAvatarURL();

	_=models.UpdateUser(user)
	cache.Remove(cache.GetKey(cache.UserInfo,id))
	g.Success(http.StatusOK,"SET_AVATAR_SUCCESS")
}

func GetUserInfoById(c *gin.Context)  {
	g:=response.Gin{C: c}
	idStr:=c.Param("id")
	id,ok:=strconv.Atoi(idStr)
	if !(ok==nil&&id>0){
		g.Error(http.StatusOK,response.ERROR_PARAM_NOT_VAILD,nil)
		return
	}
	user:=cache.GetOrCreate(cache.GetKey(cache.UserInfo,id), func() interface{} {
		user,err:= models.GetUserById(id)
		if err!=nil	{
			return &models.User{
				ID: -1,
			}
		}
		return user
	}).(*models.User)

	if user.ID<=0 {
		g.Error(http.StatusOK,response.ERROR_USER_NOT_FOUND,nil)
		return
	}
	url:="https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"
	if len(user.AvatarURL)>0{
		url=user.AvatarURL
	}

	g.Success(http.StatusOK,UserInfoRes{
		Name: user.Name,
		Userid: user.ID,
		Email: "",
		Title: user.Role,
		Group: user.Department,
		Avatar: url,
	})
	return
}

func GetUserDetailsById(c *gin.Context)  {
	g:=response.Gin{C: c}
	idStr:=c.Param("id")
	id,ok:=strconv.Atoi(idStr)
	if !(ok==nil&&id>0){
		g.Error(http.StatusOK,response.ERROR_PARAM_NOT_VAILD,nil)
		return
	}
	user:=cache.GetOrCreate(cache.GetKey(cache.UserInfo,id), func() interface{} {
		user,err:= models.GetUserById(id)
		if err!=nil	{
			return &models.User{
				ID: -1,
			}
		}
		return user
	}).(*models.User)

	if user.ID<=0 {
		g.Error(http.StatusOK,response.ERROR_USER_NOT_FOUND,nil)
		return
	}

	url:="https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"
	if len(user.AvatarURL)>0{
		url=user.AvatarURL
	}

	g.Success(http.StatusOK, UserDetailsRes{
		Id: user.ID,
		StudentId: user.StudentId,
		Name: user.Name,
		Email: "",
		Title: user.Role,
		Group : user.Department,
		ComputerCount: user.ComputerFixedCount,
		ApplicanceCount: user.ApplianceFixedCount,
		Avatar: url,
	})
	return
}

func GetUserList(c *gin.Context)  {
	//g:=response.Gin{C: c}
	//countStr:=c.Query("count")
	//pageIdStr:=c.Query("pageId")
	//roleStr:=c.Query("role")
	//departmentStr:=c.Query("department")
	//
	//
	//var  count int
	//var  pageId int
	//var  role int
	//var department int
	//
	//
	//invalid:=false
	//
	//
	//// check pageId
	//if len(pageIdStr)==0 {
	//	pageId=1
	//} else {
	//	num,ok:=strconv.Atoi(pageIdStr)
	//	if ok!=nil||num<=0 {
	//		invalid=true
	//	} else {
	//		pageId=num
	//	}
	//}
	//
	////check pageSize
	//if len(countStr)==0 {
	//	count=30
	//} else {
	//	num,ok:=strconv.Atoi(countStr)
	//	if ok!=nil||num<0||num>30 {
	//		invalid=true
	//	} else {
	//		count=num
	//	}
	//}
	//
	//
	////check queryType
	//if len(queryTypeStr)==0 {
	//	queryType=0
	//} else {
	//	num,ok:=strconv.Atoi(queryTypeStr)
	//	if ok!=nil||num<0||num>2 {
	//		invalid=true
	//	} else {
	//		queryType=num
	//	}
	//}
	//
	//if invalid{
	//	g.Error(http.StatusOK,response.ERROR_TICKET_SEARCH_INVALID_QUERY,nil)
	//	return
	//}
}