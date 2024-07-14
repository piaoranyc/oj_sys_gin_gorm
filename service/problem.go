package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oj/define"
	"oj/models"
)

// GetProblemList
// @Tags         公共方法
// @Summary      问题列表
// @Param        page   query      int  false  "请输入当前页,默认第一页"
// @Param        size   query      int  false  "size"
// @Param        keyword   query      int  false  "keyword"
// @Success      200  string json "{"code":"200","msg":,"",data:""}"
// @Router       /problem_list [get]
func GetProblemList(context *gin.Context) {
	page := context.DefaultQuery("page", define.DefaultPage)
	size := context.DefaultQuery("size", define.DefaultSize)
	keyword := context.Query("keyword")
	models.GetProblemList(page, size, keyword)
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "get problem list",
	})

}
