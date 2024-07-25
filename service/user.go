package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"oj/help"
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

// Login
// @Tags         公共方法
// @Summary      用户登录
// @Param        username   formData   string  false  "username"
// @Param        password   formData   string  false  "password"
// @Success      200  string json "{"code":"200","msg":,"",data:""}"
// @Router       /login [post]
func Login(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")

	if username == "" || password == "" {
		context.JSON(200, gin.H{
			"code": -1,
			"msg":  "必填信息内容为空",
		})
	}
	password = help.GetMd5(password)
	log.Println(username, password)
	data := models.UserBasic{}
	err := models.DB.Where("name = ? AND password= ?", username, password).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(200, gin.H{
				"code": -1,
				"msg":  "用户名和密码错误",
			})
		}

		context.JSON(200, gin.H{
			"code": -1,
			"msg":  "Get user fail" + err.Error(),
		})
		return
	}
	context.JSON(200, gin.H{
		"code": 200,
		"data": data,
	})
	context.JSON(200, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"token": "token",
		},
	})
}
