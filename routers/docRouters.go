package routers

import (
	"ResumeManagement/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func DocsRouter(c *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/v1"
	docRouter := c.Group("")
	{
		docRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
}
