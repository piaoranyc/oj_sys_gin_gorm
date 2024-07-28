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
	token, err := help.GenerateJwt(data.Identity, data.Name)
	if err != nil {
		context.JSON(200, gin.H{
			"code": -1,
			"msg":  "generate token fail" + err.Error(),
		})
	}
	context.JSON(200, gin.H{
		"code": 200,
		"data": data,
	})
	context.JSON(200, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}

// SendCode
// @Tags         公共方法
// @Summary      发送验证码
// @Param        email   formData   string  true  "email"
// @Success      200  string json "{"code":"200","msg":,"",data:""}"
// @Router       /send-code [post]
func SendCode(context *gin.Context) {
	email := context.PostForm("email")
	if email == "" {
		context.JSON(200, gin.H{

			"code": -1,
			"msg":  "不为空",
		})
		return

	}

	code := "12345"
	err := help.SendCode(email, code)
	if err != nil {
		context.JSON(200, gin.H{
			"code": -1,
			"msg":  err.Error() + "发送验证码失败",
		})
		return
	}
	context.JSON(200, gin.H{
		"code": 200,
		"msg":  "验证码发送成功",
	})
}

// Register
// @Tags         公共方法
// @Summary      用户注册
// @Param        email   formData   string  true  "email"
// @Param        code   formData   string  true  "code"
// @Param        name   formData   string  true  "name"
// @Param        password   formData   string  true  "password"
// @Param        phone   formData   string  false  "phone"
// @Success      200  string json "{"code":"200","msg":,"",data:""}"
// @Router       /register [post]
func Register(context *gin.Context) {
	mail := context.PostForm("mail")
	if mail == "" {
		context.JSON(200, gin.H{
			"code": -1,
			"msg":  "邮箱为空",
		})
		return
	}
	code := context.PostForm("code")
	if code == "" {
		context.JSON(200, gin.H{
			"code": -1,
			"msg":  "验证码为空",
		})
		return
	}
	name := context.PostForm("nama")
	if name == "" {
		context.JSON(200, gin.H{
			"code": -1,
			"msg":  "用户名为空",
		})
		return
	}
	password := context.PostForm("password")
	if password == "" {
		context.JSON(200, gin.H{
			"code": -1,
			"msg":  "密码为空",
		})
		return
	}
	phone := context.PostForm("phone")
	code, err := models.RDB.Get(context, mail).Result()
	if err != nil {
		context.JSON(200, gin.H{
			"code": -1,
			"msg":  err.Error() + "获取验证码错误",
		})
		return
	}
}
