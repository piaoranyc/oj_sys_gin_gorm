package service

import (
	"github.com/gin-gonic/gin"
	"oj/models"
)

// GetUserDetail
// @Tags         公共方法
// @Summary      用户详情
// @Param        identity   query   string  false  "user identity"
// @Success      200  string json "{"code":"200","msg":,"",data:""}"
// @Router       /user_detail [get]
func GetUserDetail(context *gin.Context) {
	identity := context.Query("identity")
	if identity == "" {
		context.JSON(200, gin.H{
			"code": -1,
			"msg":  "用户标识不能为空",
		})
		return
	}
	data := models.UserBasic{}
	err := models.DB.Omit("password").Where("identity = ?", identity).Find(&data).Error
	if err != nil {
		context.JSON(200, gin.H{
			"code": -1,
			"msg":  "get user" + identity + " detail fail",
		})
		return
	}
	context.JSON(200, gin.H{
		"code": 200,
		"data": data,
	})
}
