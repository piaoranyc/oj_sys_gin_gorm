package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"oj/define"
	"oj/models"
	"strconv"
)

// GetProblemList
// @Tags         公共方法
// @Summary      问题列表
// @Param        keyword   query   string  false  "keyword"
// @Param        page   query      int  false  "请输入当前页,默认第一页"
// @Param        size   query      int  false  "size"
// @Success      200  string json "{"code":"200","msg":,"",data:""}"
// @Router       /problem_list [get]
func GetProblemList(context *gin.Context) {
	page, err := strconv.Atoi(context.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println(err)
		return
	}
	size, err := strconv.Atoi(context.DefaultQuery("size", define.DefaultSize))
	if err != nil {
		log.Println(err)
		return
	}
	page = (page - 1) * size
	var count int64
	keyword := context.Query("keyword")
	list := make([]*models.Problem, 0)
	tx := models.GetProblemList(keyword)
	err = tx.Count(&count).Omit("content").Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Panicln("get problem list:", err)
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"msg": 200,
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})

}
