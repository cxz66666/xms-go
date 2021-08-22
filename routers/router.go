package routers

import (
	"github.com/gin-gonic/gin"
	"xms/middleware"
	"xms/pkg/setting"
	"xms/routers/api/v1/activity_controller"
	"xms/routers/api/v1/auth_controller"
	"xms/routers/api/v1/bill_controller"
)
func InitRouter() *gin.Engine {
	r:=gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	apiv1:=r.Group("/api")
	apiv1.Use(middleware.RewriteToken())
	{
		activity:=apiv1.Group("/activities")
		activity.Use(middleware.Jwt(),middleware.StaffOnly())
		{
			activity.GET("/current",activity_controller.GetLatestActivities)
		}

		authController:=apiv1.Group("/auth")
		{
			authController.POST("/login",middleware.IPBlock(),auth_controller.Login)
			authController.GET("/setToken",middleware.Jwt(),middleware.StaffOnly(),auth_controller.SetToken)
			authController.GET("/logout",middleware.Jwt(),middleware.StaffOnly(),auth_controller.Logout)
			authController.GET("/ping",auth_controller.Ping)
		}

		billController:=apiv1.Group("/bill")
		billController.Use(middleware.Jwt(),middleware.StaffOnly())
		{
			billController.POST("/create",bill_controller.CreateNewBill)
			billController.GET("/:id",bill_controller.GetBill)
			billController.DELETE("/:id",middleware.SeniorMemberOnly(),bill_controller.DeleteBill)
			billController.PUT("/:id",middleware.SeniorMemberOnly(),bill_controller.UpdateBill)
			billController.GET("/list",bill_controller.GetBillList)
			billController.GET("/stats/amount",bill_controller.GetStatsAmount)
			billController.GET("/stats/count",bill_controller.GetStatsCount)
		}

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