package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "oj/docs"
	"oj/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", service.Ping)
	//用户
	r.GET("/user_detail", service.GetUserDetail)
	r.POST("/login", service.Login)
	//问题
	r.GET("/problem_list", service.GetProblemList)
	r.GET("/problem_detail", service.GetProblemDetail)
	//提交
	r.GET("/submit_list", service.GetSubmitList)
	return r
}
