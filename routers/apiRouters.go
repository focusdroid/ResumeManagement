package routers

import (
	"ResumeManagement/controllers/api"
	"ResumeManagement/middleware"
	"github.com/gin-gonic/gin"
)

func ApiRouters(r *gin.Engine) {
	//docs.SwaggerInfo.BasePath = "/api/v1"
	apiRouter := r.Group("", middleware.InitMiddleware)
	{
		apiRouter.POST("/login", api.APIController{}.Login)
		apiRouter.POST("/register", api.APIController{}.Register)
		apiRouter.POST("/sendMail", api.APIController{}.SendMail)
		apiRouter.GET("/refreshToken", api.APIController{}.RefreshToken)
	}
}
