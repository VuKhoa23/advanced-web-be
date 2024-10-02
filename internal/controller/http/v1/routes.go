package v1

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func MapRoutes(router *gin.Engine, studentHandler *StudentHandler) {
	v1 := router.Group("/api/v1")
	{
		students := v1.Group("/students")
		{
			students.GET("/", studentHandler.GetAll)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
