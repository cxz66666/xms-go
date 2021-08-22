package response

import (
	"errors"
	"github.com/gin-gonic/gin"
	"xms/models"
	"xms/pkg/authUtils"
)

type Gin struct {
	C *gin.Context
}

type Status uint8
const (
	OK Status =iota
	ERROR
)
type SuccessResponseDto struct {
	
	Status string `json:"status"`

	Data interface{} `json:"data"`
}

func SuccessDto(data interface{}) SuccessResponseDto {
	return SuccessResponseDto{
		Status: "success",
		Data: data,
	}
}


type ErrorResponseDto struct {
	
	Status string `json:"status"`
	
	Data interface{} `json:"data"`
}

func ErrorDto(data interface{}) ErrorResponseDto {
	return ErrorResponseDto{
		Status: "error",
		Data: data,
	}
}

// Response provide a early return function to lower invasive, but at most time we use Success and Error to instead it
func (g *Gin)Response(httpCode int,status Status ,data interface{})  {
	if status==OK{
		 g.C.JSON(httpCode, SuccessDto(data))
	} else if status==ERROR {
		g.C.JSON(httpCode, ErrorDto(data))
	} else {
		g.C.JSON(httpCode,gin.H{
			"data":"internal error",
		})
	}
}

//Success is a more simple function to return data
func (g *Gin)Success(httpCode int,data interface{})  {
	g.C.JSON(httpCode, SuccessDto(data))
}

//Error is a more simple function to return data, and it will set abort for context
func (g *Gin) Error(httpCode int, errorCode int, err error) {
	g.C.JSON(httpCode, ErrorDto(gin.H{
		"code":errorCode,
		"msg":GetMsg(errorCode),
		"err":err,
	}))
	//Important!!!!!
	g.C.Abort()
}

//GetUserId get userId from context, must be used after context
func (g *Gin)GetUserId() int {
	if claim,ok:=g.C.Get(authUtils.Claim);!ok{
		return -1
	} else {
		eva,ok:=claim.(*authUtils.EvaClaimTypes)
		if !ok {
			return -1
		}
		return eva.UserId
	}
}

//GetDepartment get department from context, must be used after context
func (g *Gin)GetDepartment() (models.Department,error) {
	if claim,ok:=g.C.Get(authUtils.Claim);!ok{
		return models.DN,errors.New("don't have token")
	} else {
		eva,ok:=claim.(*authUtils.EvaClaimTypes)
		if !ok {
			return models.DN,errors.New("parse error")
		}
		return eva.Department,nil
	}
}

func (g *Gin)GetWechatSessionKeyEncrypted() string {
	if claim,ok:=g.C.Get(authUtils.Claim);!ok{
		return ""
	} else {
		we,ok:=claim.(*authUtils.WechatClaimTypes)
		if !ok {
			return ""
		}
		return we.SessionKey
	}
}
func (g *Gin)GetWechatOpenId() string {
	if claim,ok:=g.C.Get(authUtils.Claim);!ok{
		return ""
	} else {
		we,ok:=claim.(*authUtils.WechatClaimTypes)
		if !ok {
			return ""
		}
		return we.OpenId
	}
}