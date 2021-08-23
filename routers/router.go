package routers

import (
	"github.com/gin-gonic/gin"
	"xms/middleware"
	"xms/pkg/setting"
	"xms/routers/api/v1/activity_controller"
	"xms/routers/api/v1/auth_controller"
	"xms/routers/api/v1/bill_controller"
	"xms/routers/api/v1/news_controller"
	"xms/routers/api/v1/stats_controller"
	"xms/routers/api/v1/ticket_controller"
	"xms/routers/api/v1/user_controller"
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

		newController:=apiv1.Group("/news")
		newController.Use(middleware.Jwt(),middleware.StaffOnly())
		{
			newController.GET("/current",news_controller.GetLatestNews)
			newController.GET("/get/current",news_controller.GetLatestNews)
			newController.GET("/:id",news_controller.GetNews)
			newController.POST("/new",middleware.SeniorMemberOnly(),news_controller.PostNeNews)
			newController.PUT("/:id",middleware.SeniorMemberOnly(),news_controller.UpdateNews)
			newController.DELETE("/:id",news_controller.DeleteNews)
			newController.GET("/list",news_controller.GetNewsList)
		}

		statsController:=apiv1.Group("/stats")
		statsController.Use(middleware.Jwt(),middleware.StaffOnly())

		{
			statsController.GET("/ticket/count",stats_controller.GetTicketCount)
		}

		ticketCountroller:=apiv1.Group("/ticket")
		ticketCountroller.Use(middleware.Jwt(),middleware.StaffOnly())
		{
			ticketCountroller.GET("/search",ticket_controller.SearchTicketList)
			ticketCountroller.GET("/list",ticket_controller.GetTicketList)
			ticketCountroller.GET("/:id",ticket_controller.GetTicketInfo)
			ticketCountroller.GET("/get/:id",ticket_controller.GetTicketInfo)
			ticketCountroller.PUT("/:id",ticket_controller.UpdateTicketInfo)
			ticketCountroller.PUT("/update/:id",ticket_controller.UpdateTicketInfo)
			ticketCountroller.POST("/migrate")
			ticketCountroller.POST("/new",ticket_controller.CreateNewTicket)
			ticketCountroller.POST("/pick/:id",ticket_controller.PickupTicket)
			ticketCountroller.POST("/lock/:id",ticket_controller.LockTicket)
			ticketCountroller.POST("/unlock/:id",ticket_controller.UnlockTicket)
			ticketCountroller.POST("/cs/:id",ticket_controller.ChangeTicketStatus)
			ticketCountroller.DELETE("/:id",middleware.SeniorMemberOnly(),ticket_controller.DeleteTicket)
			ticketCountroller.DELETE("/delete/:id",middleware.SeniorMemberOnly(),ticket_controller.DeleteTicket)
			ticketCountroller.POST("/newcomment/:id",ticket_controller.PostNewComment)

		}

		userController:=apiv1.Group("/user")
		userController.Use(middleware.Jwt(),middleware.StaffOnly())
		{
			userController.GET("/info/me",user_controller.GetMyInfo)
			userController.GET("/setavatar",user_controller.SetUserAvatar)
			userController.GET("/info/:id",user_controller.GetUserInfoById)
			userController.GET("details/:id",user_controller.GetUserDetailsById)
			userController.GET("/list",user_controller.GetUserList)

		}
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