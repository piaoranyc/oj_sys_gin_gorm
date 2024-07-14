package router

import (
	"github.com/gin-gonic/gin"
	"oj/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", service.Ping)

	return r
}
