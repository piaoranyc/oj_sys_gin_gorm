package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"oj/define"
	"oj/help"
	"oj/models"
	"strconv"
	"time"
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
	token, err := help.GenerateJwt(data.Identity, data.Name, data.IsAdmin)
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

	code := help.GenerateCode()
	models.RDB.Set(context, email, code, time.Second*300)
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
// @Param        mail   formData   string  true  "mail"
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
	userCode := context.PostForm("code")
	if userCode == "" {
		context.JSON(200, gin.H{
			"code": -1,
			"msg":  "验证码为空",
		})
		return
	}
	name := context.PostForm("name")
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
	sysCode, err := models.RDB.Get(context, mail).Result()
	if err != nil {
		log.Println(err)
		context.JSON(200, gin.H{
			"code": -1,
			"msg":  err.Error() + "获取验证码错误",
		})
		return
	}
	if sysCode != userCode {
		context.JSON(200, gin.H{
			"code": -1,
			"msg":  "验证码错误",
		})
		return
	}
	var count int64
	err = models.DB.Where("mail = ?", mail).Model(new(models.UserBasic)).Count(&count).Error
	if err != nil {
		context.JSON(200, gin.H{
			"code": -1,
			"msg":  err.Error() + "获取数据失败",
		})
		return
	}
	if count > 0 {
		context.JSON(200, gin.H{
			"code": 200,
			"msg":  "该游戏已经被注册",
		})
		return
	}
	data := models.UserBasic{
		Identity: help.GetUUID(),
		Name:     name,
		Password: help.GetMd5(password),
		Phone:    phone,
		Mail:     mail,
	}
	err = models.DB.Create(&data).Error
	if err != nil {
		context.JSON(200, gin.H{
			"code": -1,
			"msg":  err.Error() + "创建用户错误",
		})
		return
	}
	//生成token
	jwt, err := help.GenerateJwt(data.Identity, data.Name, data.IsAdmin)
	if err != nil {
		context.JSON(200, gin.H{
			"code": -1,
			"msg":  "generate token fail" + err.Error(),
		})
		return
	}

	context.JSON(200, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"token": jwt,
		},
	})
}

// GetRankList
// @Tags         公共方法
// @Summary      排行榜
// @Param        keyword   query   string  false  "keyword"
// @Param        page   query      int  false  "请输入当前页,默认第一页"
// @Param        size   query      int  false  "size"
// @Param        category_identity   query    string  false  "分类的唯一标识"
// @Success      200  string json "{"code":"200","msg":,"",data:""}"
// @Router       /rank-list [get]
func GetRankList(context *gin.Context) {
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

	list := make([]*models.UserBasic, 0)
	err = models.DB.Model(new(models.UserBasic)).Count(&count).Order("finish_problem_num DESC, submit_num ASC").
		Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		context.JSON(200, gin.H{
			"code": -1,
			"msg":  err.Error() + "get rank list fail",
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"msg": 200,
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})
}
