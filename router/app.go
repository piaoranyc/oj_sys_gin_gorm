package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"oj/middlewares"

	_ "oj/docs"
	"oj/service"
)

func Router() *gin.Engine {
	r := gin.Default()

	//公用方法
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", service.Ping)
	//用户
	r.GET("/user_detail", service.GetUserDetail)
	r.POST("/login", service.Login)
	r.POST("/send-code", service.SendCode)
	r.POST("/register", service.Register)
	//排行榜
	r.GET("/rank-list", service.GetRankList)
	//问题
	r.GET("/problem_list", service.GetProblemList)
	r.GET("/problem_detail", service.GetProblemDetail)
	//提交
	r.GET("/submit_list", service.GetSubmitList)
	//管理员私有方法
	authAdmin := r.Group("/admin", middlewares.AuthAdminCheck())
	authAdmin.PUT("/problem_modify", service.ProblemModify)
	authAdmin.POST("/problem-create", service.ProblemCreate)

	//分类列表
	authAdmin.GET("/category_list", service.GetCategoryList)
	authAdmin.POST("/category_create", service.GetCategoryCreate)
	authAdmin.PUT("/category_modify", service.GetCategoryModify)
	authAdmin.DELETE("/category_delete", service.GetCategoryDelete)

	//用户私有方法
	authUser := r.Group("/user", middlewares.AuthAdminCheck())
	//code submit
	authUser.POST("/submit", service.Submit)
	return r
}
