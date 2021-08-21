package response

import (
	"github.com/gin-gonic/gin"
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

//Error is a more simple function to return data
func (g *Gin) Error(httpCode int, errorCode int, err error) {
	g.C.JSON(httpCode, ErrorDto(gin.H{
		"code":errorCode,
		"msg":GetMsg(errorCode),
		"err":err,
	}))
}