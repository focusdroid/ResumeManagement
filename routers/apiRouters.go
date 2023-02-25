package routers

import (
	"ResumeManagement/controllers/api"
	"ResumeManagement/middleware"
	"github.com/gin-gonic/gin"
)

func ApiRouters(r *gin.Engine) {
	//docs.SwaggerInfo.BasePath = "/api/v1"
	apiRouter := r.Group("", middleware.InitMiddleware, middleware.InitMiddlewareBlacklist)
	{
		apiRouter.POST("/login", api.APIController{}.Login)       // 登录
		apiRouter.POST("/register", api.APIController{}.Register) // 注册
		apiRouter.POST("/sendMail", api.APIController{}.SendMail) //
		apiRouter.GET("/refreshToken", api.APIController{}.RefreshToken)
		apiRouter.GET("/isLine", api.APIController{}.IsLine)
	}
}
