package routers

import (
	"github.com/gin-gonic/gin"
	"xms/pkg/setting"
	"xms/routers/api/v1/activity_controller"
)
func InitRouter() *gin.Engine {
	r:=gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	apiv1:=r.Group("/api")
	{
		activity:=apiv1.Group("/activities")
		{
			activity.GET("/current",activity_controller.GetLatestActivities)
		}

		//authController:=apiv1.Group("/auth")
		//{
		//	authController.POST("/login",_)
		//	authController.GET("/setToken",_)
		//	authController.GET("/logout",_)
		//	authController.GET("/ping",_)
		//}
		//
		//billController:=apiv1.Group("/bill")
		//{
		//	billController.POST("/create",_)
		//	billController.GET("/:id",_)
		//	billController.DELETE("/:id",_)
		//	billController.PUT("/:id",_)
		//	billController.GET("/list",_)
		//	billController.GET("/stats/amount",_)
		//	billController.GET("/stats/count",_)
		//}
		//
		//newController:=apiv1.Group("/news")
		//{
		//	newController.GET("/current",_)
		//	newController.GET("/get/current",_)
		//	newController.GET("/:id",_)
		//	newController.POST("/:id",_)
		//	newController.PUT("/:id",_)
		//	newController.DELETE("/:id",_)
		//	newController.GET("/list",_)
		//}
		//
		//statsController:=apiv1.Group("/stats")
		//{
		//	statsController.GET("/ticket/count",_)
		//}
		//
		//ticketCountroller:=apiv1.Group("/ticket")
		//{
		//	ticketCountroller.GET("/search",_)
		//	ticketCountroller.GET("/list",_)
		//	ticketCountroller.GET("/:id",_)
		//	ticketCountroller.GET("/get/:id",_)
		//	ticketCountroller.PUT("/:id",_)
		//	ticketCountroller.PUT("/update/:id",_)
		//	ticketCountroller.POST("/migrate",_)
		//	ticketCountroller.POST("/new",_)
		//	ticketCountroller.POST("/pick/:id",_)
		//	ticketCountroller.POST("/lock/:id",_)
		//	ticketCountroller.POST("/unlock/:id",_)
		//	ticketCountroller.POST("/cs/:id",_)
		//	ticketCountroller.DELETE("/:id",_)
		//	ticketCountroller.DELETE("/delete/:id",_)
		//	ticketCountroller.POST("/newcomment/:id",_)
		//
		//}
		//
		//userController:=apiv1.Group("/user")
		//{
		//	userController.GET("/info/me",_)
		//	userController.GET("/setavatar",_)
		//	userController.GET("/info/:id",_)
		//	userController.GET("details/:id",_)
		//	userController.GET("/list",_)
		//
		//}
		//
		//wechatController:=apiv1.Group("/wechat")
		//{
		//	wechatController.GET("/getToken",_)
		//	wechatController.GET("/getStatus",_)
		//	wechatController.POST("/register",_)
		//	wechatController.POST("/comment",_)
		//
		//}
	}
	return r
}