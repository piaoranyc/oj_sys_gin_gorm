package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"oj/define"
	"oj/models"
	"strconv"
)

// GetSubmitList
// @Tags         公共方法
// @Summary      提交列表
// @Param        problem_identity   query   string  false  "problem_identity"
// @Param        page   query      int  false  "请输入当前页,默认第一页"
// @Param        size   query      int  false  "size"
// @Param        user_identity   query   string  false  "user_identity"
// @Param        status   query   int  false  "status"
// @Success      200  string json "{"code":"200","msg":,"",data:""}"
// @Router       /submit_list [get]
func GetSubmitList(context *gin.Context) {
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
	problemIdentity := context.Query("problem_identity")
	userIdentity := context.Query("user_identity")

	status, _ := strconv.Atoi(context.Query("status"))
	list := make([]*models.SubmitBasic, 0)
	tx := models.GetSubmitList(problemIdentity, userIdentity, status)
	log.Println("tx:", tx)
	err = tx.Count(&count).Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Panicln("get submit list:", err)
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"msg": 200,
		"data": map[string]interface{}{
			"code":  200,
			"list":  list,
			"count": count,
		},
	})

}
