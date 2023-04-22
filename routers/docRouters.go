package routers

import (
	"ResumeManagement/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func DocsRouter(c *gin.Engine) {
	docs.SwaggerInfo.BasePath = "localhost:8080/"
	docs.SwaggerInfo.Description = "线上地址www.asmie.liv/api"
	docRouter := c.Group("")
	{
		docRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
}
